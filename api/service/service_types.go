package service

import "encoding/xml"

type ApplicationsList struct {
	XMLName      xml.Name             `xml:"list"`
	Applications []ApplicationService `xml:"application"`
}

type ApplicationService struct {
	XMLName     xml.Name         `xml:"application"`
	Name        string           `xml:"name"`
	ObjectId    string           `xml:"objectId,omitempty"`
	Type        string           `xml:"type>TypeName"`
	Revision    int              `xml:"revision,omitempty"`
	Description string           `xml:"description"`
	Element     []ServiceElement `xml:"element"`
}

type ServiceElement struct {
	XMLName             xml.Name `xml:"element"`
	ApplicationProtocol string   `xml:"applicationProtocol"`
	Value               string   `xml:"value"`
}
