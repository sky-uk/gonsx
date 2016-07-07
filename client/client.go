package client

import (
	"fmt"
	"log"
	"encoding/xml"
	"net/http"
	"crypto/tls"
	"io/ioutil"

	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"bytes"
	"io"
)

func NewNSXClient(url string, user string, password string, ignoreSSL bool, debug bool) *NSXClient {
	nsxClient := new(NSXClient)
	nsxClient.URL = url
	nsxClient.User = user
	nsxClient.Password = password
	nsxClient.IgnoreSSL = ignoreSSL
	nsxClient.debug = debug
	return nsxClient
}

type NSXClient struct {
	URL		string
	User 		string
	Password	string
	IgnoreSSL	bool
	debug 		bool
}

func (nsxClient *NSXClient) Do(api api.NSXApi) {
	requestURL := fmt.Sprintf("%s%s", nsxClient.URL, api.Endpoint())

	var requestPayload io.Reader
	if(api.RequestObject() != nil) {
		requestXmlBytes, marshallingErr := xml.Marshal(api.RequestObject())
		log.Println(string(requestXmlBytes))
		if marshallingErr != nil {
			log.Fatal(marshallingErr)
		}
		requestPayload = bytes.NewReader(requestXmlBytes)
	}
	req, err := http.NewRequest(api.Method(), requestURL, requestPayload)

	req.SetBasicAuth(nsxClient.User, nsxClient.Password)
	req.Header.Set("Content-Type", "application/xml")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: nsxClient.IgnoreSSL},
	}
	httpClient := &http.Client{Transport: tr}
	res, err := httpClient.Do(req)
	defer res.Body.Close()

	if err != nil{
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(res.Body)

	if(nsxClient.debug) {
		log.Println(string(bodyText))
	}

	if api.ResponseObject() != nil && res.Header.Get("Content-Type") == "application/xml" {
		fmt.Println("parsing xml")
		xmlerr := xml.Unmarshal(bodyText, api.ResponseObject())
		if xmlerr != nil {
			panic(xmlerr)
		}
	} else {
		api.SetResponseObject(string(bodyText))
	}
}
