package api

type BaseApi struct {
	method	 	string
	endpoint 	string
	requestObject 	interface{}
	responseObject 	interface{}
}

func NewBaseApi(method string, endpoint string, requestObject interface{}, responseObject interface{}) *BaseApi {
	return &BaseApi{method, endpoint, requestObject, responseObject}
}

func (this BaseApi) Method() string {
	return this.method
}

func (this BaseApi) Endpoint() string {
	return this.endpoint
}

func (this BaseApi) RequestObject() interface{} {
	return this.requestObject
}

func (this BaseApi) ResponseObject() interface{} {
	return this.responseObject
}
