package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getVirtualWireApi *GetVirtualWireApi

func setupGet(id string) {
	getVirtualWireApi = NewGet(id)
}

func TestGetMethod(t *testing.T) {
	setupGet("virtualwire-1")
	assert.Equal(t, http.MethodGet, getVirtualWireApi.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("virtualwire-1")
	assert.Equal(t, "/api/2.0/vdn/virtualwires/virtualwire-1", getVirtualWireApi.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("virtualwire-1")
	xmlContent := []byte("<virtualWire><objectId>virtualwire-1</objectId></virtualWire>")

	xmlerr := xml.Unmarshal(xmlContent, getVirtualWireApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "virtualwire-1", getVirtualWireApi.GetResponse().ObjectID)
}
