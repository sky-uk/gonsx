package api

type NSXApi struct {
	Ctx	*baseContext
}

type BaseRequest struct {
	Method		string
	Endpoint	string
}


type baseContext struct {
	Method 		string
	Endpoint	string
}

func (* baseContext) GetObject() interface {}

func NewBaseRequest(methodArg string, endpoint string) BaseRequest {
	request := new(BaseRequest)
	request.Method = methodArg
	request.Endpoint = endpoint
	return *request
}

type BaseResponse interface {
	StatusCode()	 int
	ResponseObject() interface{}
}