package nsxclient

import (
	"encoding/xml"
	"net/http"
	"crypto/tls"
	"bytes"
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

type Scope struct {
	XMLName		xml.Name	`xml:"scope"`
	Id		string 		`xml:"id"`
	ObjectTypeName	string 		`xml:"objectTypeName"`
	Name		string 		`xml:"name"`
}

type Element struct {
	XMLName			xml.Name	`xml:"element"`
	ApplicationProtocol	string		`xml:"applicationProtocol"`
	Value			string		`xml:"value"`
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