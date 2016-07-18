package securitytags

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"encoding/xml"
)

var createSecurityTagAPI *CreateSecurityTagAPI

func createSetup() {
	createSecurityTagAPI = NewCreate("test", "test desc")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createSecurityTagAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/securitytags/tag", createSecurityTagAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	object := new(SecurityTag)
	object.Name = "test"
	object.Description = "test desc"
	object.TypeName = "SecurityTag"
	expectedXML := "<securityTag><name>test</name><description>test desc</description><type><typeName>SecurityTag</typeName></type></securityTag>"

	xmlBytes, err :=xml.Marshal(object)

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}
