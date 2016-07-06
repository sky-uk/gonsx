package api

import (
	"fmt"
)

type GetTransportZoneApi struct {
	method 		string
	endpoint	string
	responseObject	interface{}
}

func NewGetTransportZone() *GetTransportZoneApi {
	this := new(GetTransportZoneApi)
	this.method = "GET"
	this.endpoint = "/api/2.0/vdn/scopes"
	this.responseObject = new(TransportZoneResponseObject)
	return this
}

func (this GetTransportZoneApi) GetMethod() string {
	return this.method
}

func (this GetTransportZoneApi) GetEndpoint() string {
	return this.endpoint
}

func (this GetTransportZoneApi) GetResponseObject() interface{} {
	return this.responseObject
}

func (this GetTransportZoneApi) GetResponse() *TransportZoneResponseObject {
	return this.responseObject.(*TransportZoneResponseObject)
}

type TransportZoneResponseObject struct {
	NetworkScopeList	[]NetworkScope	`xml:"vdnScope"`
}

type NetworkScope struct {
	ObjectId	string	`xml:"objectId"`
}

func (s NetworkScope) String() string {
	return fmt.Sprintf("%s", s.ObjectId)
}
