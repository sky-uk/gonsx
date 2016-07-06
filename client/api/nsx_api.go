package api

type NSXApi interface {
	GetMethod() 		string
	GetEndpoint() 		string
	GetResponseObject() 	interface{}
}
