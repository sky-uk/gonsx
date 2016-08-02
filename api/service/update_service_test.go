package service

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateServiceAPI *UpdateServiceAPI

func updateSetup() {
	testServiceObj := new(ApplicationService)
	testServiceObj.Name = "application-0001"
	testServiceObj.Description = "dummy description"
	element := Element{ApplicationProtocol: "TCP", Value: "8080"}
	testServiceObj.Element = []Element{element}

	updateServiceAPI = NewUpdate("application-0001", testServiceObj)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateServiceAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/application/application-0001", updateServiceAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<application><name>application-0001</name><description>dummy description</description><element><applicationProtocol>TCP</applicationProtocol><value>8080</value></element></application>"

	xmlBytes, err := xml.Marshal(updateServiceAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestUpdateUnMarshalling(t *testing.T) {
	// TODO
}
