package service

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createServiceAPI *CreateServiceAPI

func createSetup() {
	createServiceAPI = NewCreate("globalroot-0", "test_8080", "Test TCP", "TCP", "8080")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createServiceAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/application/globalroot-0", createServiceAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<application><name>test_8080</name><description>Test TCP</description><element><applicationProtocol>TCP</applicationProtocol><value>8080</value></element></application>"

	xmlBytes, err := xml.Marshal(createServiceAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
