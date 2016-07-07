package nsxclient

import (
	"encoding/xml"
	"net/http"
	"crypto/tls"
	"bytes"
	"fmt"
	"io/ioutil"
)


type Application struct {
	XMLName		xml.Name	`xml:"application"`
	ObjectID	string		`xml:"objectId"`
	Revision	int		`xml:"revision"`
	Type		Type		`xml:"type"`
	Name		string		`xml:"name"`
	Description	string		`xml:"description"`
	Element		[]Element	`xml:"element"`
}

type Type struct {
	XMLName		xml.Name	`xml:"type"`
	TypeName	string		`xml:"typeName"`
}

type Element struct {
	XMLName			xml.Name	`xml:"element"`
	ApplicationProtocol	string		`xml:"applicationProtocol"`
	Value			string		`xml:"value"`
}

type Query struct {
	VdnScope []Scope `xml:"vdnScope"`
}


type Scope struct {
	Name string `xml:"name"`
	ObjectID string `xml:"objectId"`
}


// Wrapper around httpClient, provides the context for other functions.
type NsxClient struct {
	nsxManagerHost 	string
	nsxUser		string
	nsxPassword	string
	httpClient 	*http.Client
}

func NewNsxClient(nsxmanager, nsxusername, nsxpassword string) (*NsxClient, error) {
	tr := &http.Transport{ TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	c := NsxClient{
		httpClient: &http.Client{Transport: tr},
		nsxManagerHost:   nsxmanager,
		nsxUser: nsxusername,
		nsxPassword: nsxpassword,
	}
	return &c, nil
}

func (c *NsxClient) GetScope(scopeName string) string{
	req, err := http.NewRequest("GET", "https://" + c.nsxManagerHost + "/api/2.0/vdn/scopes", nil)
	req.SetBasicAuth(c.nsxUser, c.nsxPassword)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	r, err := ioutil.ReadAll(resp.Body)
	var q Query
	err = xml.Unmarshal(r, &q)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	fmt.Println(len(q.VdnScope))
	if len(q.VdnScope) > 1 {
		for _, scopeLoop := range q.VdnScope {
			//fmt.Print(test)
			if (scopeLoop.Name == scopeName) {
				return scopeLoop.ObjectID
			}
		}
	} else if len(q.VdnScope) == 1 {
		return q.VdnScope[0].ObjectID
	}
	return "No such scope"

}

func (c *NsxClient) _doPost(service_xml *Application) (*http.Response, error) {
	var buf bytes.Buffer
	xmlbuf := xml.NewEncoder(&buf)
	err := xmlbuf.Encode(service_xml)
	if err != nil {
		return nil, err
	}
	url := "https://" + c.nsxManagerHost + "/api/2.0/services/application/globalroot-0"
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.nsxUser, c.nsxPassword)
	req.Header.Set("Content-Type", "application/xml")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *NsxClient) CreateService(name, desc, proto, port string) (*http.Response, error) {
	var xml Application
	var e Element
	xml.Name = name
	xml.Description = desc
	e.ApplicationProtocol = proto
	e.Value = port
	xml.Element = append(xml.Element, e)
	return c._doPost(&xml)
}