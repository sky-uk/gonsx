package ipset

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createIpSetAPI *CreateIpSetAPI

func createSetup() {

  ipSets := new(IpSets)
  varIpSet := IpSet{ Value : "10.50.0.0/8", Name : "CIDR_ENV", Description : "CIDR_ENV_Description"}

  ipSets.IpSets = append(ipSets.IpSets, varIpSet)

  createIpSetAPI = NewCreate("globalroot-0", ipSets)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createIpSetAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/ipset/globalroot-0", createIpSetAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<ipset><name>CIDR_ENV</name><description>CIDR_ENV_Description</description><value>10.50.0.0/8</value></ipset>"

	xmlBytes, err := xml.Marshal(createIpSetAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
