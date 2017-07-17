package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getVirtualWireAPI *GetVirtualWireAPI

func setupGet(id string) {
	getVirtualWireAPI = NewGet(id)
}

func TestGetMethod(t *testing.T) {
	setupGet("virtualwire-1")
	assert.Equal(t, http.MethodGet, getVirtualWireAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("virtualwire-1")
	assert.Equal(t, "/api/2.0/vdn/virtualwires/virtualwire-1", getVirtualWireAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("virtualwire-1")
	xmlContent := []byte("<virtualWire><objectId>virtualwire-1</objectId></virtualWire>")

	xmlerr := xml.Unmarshal(xmlContent, getVirtualWireAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "virtualwire-1", getVirtualWireAPI.GetResponse().ObjectID)
}

func TestGetComplexUnMarshalling(t *testing.T) {
	setupGet("virtualwire-1")
	xmlContent := []byte("<virtualWire><objectId>virtualwire-1</objectId><name>name</name><vdnId>vdnId</vdnId><vdsContextWithBacking><switch><objectId>dvs-1</objectId></switch></vdsContextWithBacking></virtualWire>")

	xmlerr := xml.Unmarshal(xmlContent, getVirtualWireAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "virtualwire-1", getVirtualWireAPI.GetResponse().ObjectID)
	assert.Equal(t, "vdnId", getVirtualWireAPI.GetResponse().VdnID)
	assert.Equal(t, "name", getVirtualWireAPI.GetResponse().Name)
	assert.Equal(t, "dvs-1", getVirtualWireAPI.GetResponse().VdsContext[0].Switch.ObjectID)
}

func TestGetComplexUnMarshallingWithMultipleVdsContextBlocks(t *testing.T) {
	setupGet("virtualwire-1")
	xmlContent := []byte("<virtualWire><objectId>virtualwire-1</objectId><vdsContextWithBacking><switch><objectId>dvs-1</objectId></switch></vdsContextWithBacking><vdsContextWithBacking><switch><objectId>dvs-2</objectId></switch></vdsContextWithBacking></virtualWire>")

	xmlerr := xml.Unmarshal(xmlContent, getVirtualWireAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "virtualwire-1", getVirtualWireAPI.GetResponse().ObjectID)
	assert.Equal(t, "dvs-1", getVirtualWireAPI.GetResponse().VdsContext[0].Switch.ObjectID)
	assert.Equal(t, "dvs-2", getVirtualWireAPI.GetResponse().VdsContext[1].Switch.ObjectID)
}
