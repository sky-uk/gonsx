package client

import (
	"fmt"
	"log"
	"encoding/xml"
	"net/http"
	"crypto/tls"
	"io/ioutil"

	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
)

func NewNSXClient(url string, user string, password string, ignoreSSL bool) *NSXClient {
	nsxClient := new(NSXClient)
	nsxClient.URL = url
	nsxClient.User = user
	nsxClient.Password = password
	nsxClient.IgnoreSSL = ignoreSSL
	return nsxClient
}

type NSXClient struct {
	URL		string
	User 		string
	Password	string
	IgnoreSSL	bool
}

func (nsxClient *NSXClient) get(api *api.NSXApi) {
	requestURL := fmt.Sprintf("%s%s", nsxClient.URL, api.Ctx.Endpoint)
	log.Println(requestURL)
	req, err := http.NewRequest(api.Ctx.Method, requestURL, nil)
	req.SetBasicAuth(nsxClient.User, nsxClient.Password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: nsxClient.IgnoreSSL},
	}
	httpClient := &http.Client{Transport: tr}
	res, err := httpClient.Do(req)

	if err != nil{
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(res.Body)
	log.Print(string(bodyText))

	xmlerr := xml.Unmarshal(bodyText, &api.Ctx.GetObject())
	if xmlerr != nil { panic(xmlerr) }
}

// TODO: implement create

// TODO: implement update