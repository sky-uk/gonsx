package service

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateServiceAPI *UpdateServiceAPI

func updateSetup() {
	updateServiceAPI = NewUpdate("globalroot-0", "test_8080", "Test TCP", "TCP", "8080")
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPost, updateServiceAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/application/globalroot-0", updateServiceAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<application><name>test_8080</name><description>Test TCP</description><element><applicationProtocol>TCP</applicationProtocol><value>8080</value></element></application>"

	xmlBytes, err := xml.Marshal(updateServiceAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestUpdateUnMarshalling(t *testing.T) {
	// TODO
}
