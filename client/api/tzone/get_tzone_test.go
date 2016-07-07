package tzone

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"net/http"
)

var getTransportZoneApi *GetTransportZoneApi

func setupGet(id string) {
	getTransportZoneApi = NewGet(id)
}

func TestGetMethod(t *testing.T) {
	setupGet("id1")
	assert.Equal(t, http.MethodGet, getTransportZoneApi.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("id1")
	assert.Equal(t, "/api/2.0/vdn/scopes/id1", getTransportZoneApi.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("id1")
	xmlContent := []byte("<vdnScope><objectId>vdnscope-1</objectId></vdnScope>")

	xmlerr := xml.Unmarshal(xmlContent, getTransportZoneApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "vdnscope-1", getTransportZoneApi.GetResponse().ObjectId)
}
