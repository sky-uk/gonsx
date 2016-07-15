package tzone

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getTransportZoneAPI *GetTransportZoneAPI

func setupGet(id string) {
	getTransportZoneAPI = NewGet(id)
}

func TestGetMethod(t *testing.T) {
	setupGet("id1")
	assert.Equal(t, http.MethodGet, getTransportZoneAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("id1")
	assert.Equal(t, "/api/2.0/vdn/scopes/id1", getTransportZoneAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("id1")
	xmlContent := []byte("<vdnScope><objectId>vdnscope-1</objectId></vdnScope>")

	xmlerr := xml.Unmarshal(xmlContent, getTransportZoneAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "vdnscope-1", getTransportZoneAPI.GetResponse().ObjectID)
}
