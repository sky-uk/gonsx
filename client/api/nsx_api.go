package api

type NSXApi interface {
	Method() 		string
	Endpoint() 		string
	RequestObject() 	interface{}
	ResponseObject() 	interface{}
	SetResponseObject(interface{})
}
