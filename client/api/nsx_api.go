package api

type NSXApi interface {
	Method() 		string
	Endpoint() 		string
	RequestObject() 	interface{}
	ResponseObject() 	interface{}
	StatusCode()		int
	RawResponse()		[]byte
	Err() 			error

	SetResponseObject(interface{})
	SetStatusCode(int)
	SetRawResponse([]byte)
}
