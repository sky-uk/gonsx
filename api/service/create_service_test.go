package service

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createServiceApi *CreateServiceApi

func createSetup() {
	createServiceApi = NewCreate("globalroot-0", "test_8080", "Test TCP", "TCP", "8080")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createServiceApi.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/application/globalroot-0", createServiceApi.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXml := "<application><name>test_8080</name><type><TypeName></TypeName></type><description>Test TCP</description><element><applicationProtocol>TCP</applicationProtocol><value>8080</value></element></application>"

	xmlBytes, err := xml.Marshal(createServiceApi.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXml, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
