package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllVirtualWiresAPI *GetAllVirtualWiresAPI

func setupGetAll() {
	getAllVirtualWiresAPI = NewGetAll("vdnscope-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllVirtualWiresAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/vdn/scopes/vdnscope-1/virtualwires", getAllVirtualWiresAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<virtualWires><dataPage><virtualWire><objectId>virtualwire-1</objectId><name>test_name</name></virtualWire></dataPage></virtualWires>")

	xmlerr := xml.Unmarshal(xmlContent, getAllVirtualWiresAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllVirtualWiresAPI.GetResponse().DataPage.VirtualWires, 1)
	assert.Equal(t, "virtualwire-1", getAllVirtualWiresAPI.GetResponse().DataPage.VirtualWires[0].ObjectID)
}
