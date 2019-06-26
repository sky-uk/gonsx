package api

import (
	"net/http"
)

// NSXApi object.
type NSXApi interface {
	Method() string
	Endpoint() string
	RequestObject() interface{}
	ResponseObject() interface{}
	RequestHeaders() http.Header
	ResponseHeaders() http.Header
	StatusCode() int
	RawResponse() []byte
	Error() error

	SetResponseObject(interface{})
	SetStatusCode(int)
	SetRawResponse([]byte)
	SetResponseHeader(http.Header)
}
