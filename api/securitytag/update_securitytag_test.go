package securitytag

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateSecurityTagAPI *UpdateSecurityTagAPI

func updateSetup() {
  
	updatePayload := SecurityTag{
		ObjectID: "testtag-1",
		Name: "testTag",
		Description: "A description",
		TypeName: "SecurityTag",

	}
	updateSecurityTagAPI = NewUpdate("testtag-1", &updatePayload)

}

func TestSecurityTagUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateSecurityTagAPI.Method())
}

func TestSecurityTagUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/securitytags/tag/testtag-1", updateSecurityTagAPI.Endpoint())
}

func TestSecurityTagUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<securityTag><objectId>testtag-1</objectId><name>testTag</name><description>A description</description><type><typeName>SecurityTag</typeName></type></securityTag>"
	xmlBytes, err := xml.Marshal(updateSecurityTagAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}
