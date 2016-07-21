package securitytag

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllSecurityTagsAPI *GetAllSecurityTagsAPI

func setupGetAll() {
	getAllSecurityTagsAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllSecurityTagsAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/services/securitytags/tag", getAllSecurityTagsAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<securityTags><securityTag><type><typeName>SecurityTag</typeName></type><objectId>securitytag-1</objectId><name>Test</name></securityTag></securityTags>")

	xmlerr := xml.Unmarshal(xmlContent, getAllSecurityTagsAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllSecurityTagsAPI.GetResponse().SecurityTags, 1)
	assert.Equal(t, "securitytag-1", getAllSecurityTagsAPI.GetResponse().SecurityTags[0].ObjectID)
	assert.Equal(t, "SecurityTag", getAllSecurityTagsAPI.GetResponse().SecurityTags[0].TypeName)
}
