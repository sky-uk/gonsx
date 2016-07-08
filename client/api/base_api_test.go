package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBaseApi(t *testing.T) {
	method := "myMethod"
	endpoint := "/myEndpoint"
	requestObject := string("")
	responseObject := string("")
	api := NewBaseApi(method, endpoint, requestObject, responseObject)

	assert.Equal(t, method, api.Method())
	assert.Equal(t, endpoint, api.Endpoint())
	assert.Equal(t, requestObject, api.RequestObject())
	assert.Equal(t, responseObject, api.ResponseObject())
}
