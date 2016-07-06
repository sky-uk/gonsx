package api

import "fmt"

type tzContext struct {
	baseContext
}

type TransportZoneApi struct { }

func NewGetTransportZone() *NSXApi {
	api := new(NSXApi)
	ctx := new(tzContext)
	ctx.Method = "GET"
	ctx.Endpoint = "/api/2.0./vdn/scopes"
	api.Ctx = ctx
	return api
}

func (* tzContext) GetObject() interface{} {
	return new(TransportZoneResponseObject)
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
