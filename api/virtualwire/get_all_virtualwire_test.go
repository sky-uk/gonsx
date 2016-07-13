package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllVirtualWiresApi *GetAllVirtualWiresApi

func setupGetAll() {
	getAllVirtualWiresApi = NewGetAll("vdnscope-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllVirtualWiresApi.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/vdn/scopes/vdnscope-1/virtualwires", getAllVirtualWiresApi.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<virtualWires><dataPage><virtualWire><objectId>virtualwire-1</objectId><name>test_name</name></virtualWire></dataPage></virtualWires>")

	xmlerr := xml.Unmarshal(xmlContent, getAllVirtualWiresApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllVirtualWiresApi.GetResponse().DataPage.VirtualWires, 1)
	assert.Equal(t, "virtualwire-1", getAllVirtualWiresApi.GetResponse().DataPage.VirtualWires[0].ObjectID)
}
