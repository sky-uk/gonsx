package securitygroup

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createSecurityGroupAPI *CreateSecurityGroupAPI

func createSetup() {
	createSecurityGroupAPI = NewCreate("globalroot-0", "OVP_sg_test")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createSecurityGroupAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/securitygroup/bulk/globalroot-0", createSecurityGroupAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<securitygroup><name>OVP_sg_test</name></securitygroup>"

	xmlBytes, err := xml.Marshal(createSecurityGroupAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
