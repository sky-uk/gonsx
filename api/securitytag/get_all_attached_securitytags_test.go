package securitytag

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllAttachedSecurityTagsAPI *GetAllAttachedSecurityTagsAPI

func setupGetAllAttached() {
	getAllAttachedSecurityTagsAPI = NewGetAllAttached("securitytag-1")
}

func TestGetAllAttachedMethod(t *testing.T) {
	setupGetAllAttached()
	assert.Equal(t, http.MethodGet, getAllAttachedSecurityTagsAPI.Method())
}

func TestGetAllAttachedEndpoint(t *testing.T) {
	setupGetAllAttached()
	assert.Equal(t, "/api/2.0/services/securitytags/tag/securitytag-1/vm", getAllAttachedSecurityTagsAPI.Endpoint())
}

func TestGetAllAttachedUnMarshalling(t *testing.T) {
	setupGetAllAttached()
	xmlContent := []byte("<basicinfolist><basicinfo><objectId>vm-1</objectId><name>test-vm</name></basicinfo></basicinfolist>")

	xmlerr := xml.Unmarshal(xmlContent, getAllAttachedSecurityTagsAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllAttachedSecurityTagsAPI.GetResponse().BasicInfoList, 1)
	assert.Equal(t, "vm-1", getAllAttachedSecurityTagsAPI.GetResponse().BasicInfoList[0].ObjectID)
	assert.Equal(t, "test-vm", getAllAttachedSecurityTagsAPI.GetResponse().BasicInfoList[0].Name)
}
