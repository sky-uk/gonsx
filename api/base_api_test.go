package api

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseApi(t *testing.T) {
	method := "myMethod"
	endpoint := "/myEndpoint"
	requestObject := string("")
	responseObject := string("")
	statusCode := 200
	rawResponse := []byte("some server response in []byte")
	err := errors.New("an error")

	api := NewBaseAPI(method, endpoint, requestObject, responseObject)

	api.SetStatusCode(statusCode)
	api.SetRawResponse(rawResponse)
	api.SetResponseObject(responseObject)
	api.SetError(err)

	assert.Equal(t, method, api.Method())
	assert.Equal(t, endpoint, api.Endpoint())
	assert.Equal(t, requestObject, api.RequestObject())
	assert.Equal(t, responseObject, api.ResponseObject())

	assert.Equal(t, statusCode, api.StatusCode())
	assert.Equal(t, rawResponse, api.RawResponse())
	assert.Equal(t, err, api.Error())
}
