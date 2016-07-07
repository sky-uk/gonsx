package lswitch

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"net/http"
)

var getLogicalSwitchApi *GetLogicalSwitchApi

func setupGet(id string) {
	getLogicalSwitchApi = NewGet(id)
}

func TestGetMethod(t *testing.T) {
	setupGet("id1")
	assert.Equal(t, http.MethodGet, getLogicalSwitchApi.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("id1")
	assert.Equal(t, "/api/2.0/vdn/switches/id1", getLogicalSwitchApi.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("id1")
	xmlContent := []byte("<vdsContext><switch><objectId>dvs-1</objectId></switch></vdsContext>")

	xmlerr := xml.Unmarshal(xmlContent, getLogicalSwitchApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "dvs-1", getLogicalSwitchApi.GetResponse().Switch.ObjectId)
}
