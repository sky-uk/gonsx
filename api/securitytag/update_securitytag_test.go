package securitytag

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"encoding/xml"
)

var updateSecurityTagAPI *UpdateSecurityTagAPI

func updateSetup() {
	SecurityTagID := "testtag-1"
	SecurityTagName :="testTag"
	SecurityTagDescription := "A description"
	updateSecurityTagAPI = NewUpdate(SecurityTagID,SecurityTagName, SecurityTagDescription)

}

func TestSecurityTagUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateSecurityTagAPI.Method())
}

func TestSecurityTagUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/securitytags/tag/testtag-1", updateSecurityTagAPI.Endpoint() )
}

func TestSecurityTagUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<securityTag><objectId>testtag-1</objectId><name>testTag</name><description>A description</description><type><typeName>SecurityTag</typeName></type></securityTag>"
	xmlBytes, err := xml.Marshal(updateSecurityTagAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

