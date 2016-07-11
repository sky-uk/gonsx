package edgeinterface

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllEdgeInterfacesApi *GetAllEdgeInterfacesApi

func setupGetAll() {
	getAllEdgeInterfacesApi = NewGetAll("edge-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllEdgeInterfacesApi.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces", getAllEdgeInterfacesApi.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<interfaces><interface><name>InterfaceName</name></interface><interface><name>InterfaceName2</name></interface></interfaces>")

	xmlerr := xml.Unmarshal(xmlContent, getAllEdgeInterfacesApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllEdgeInterfacesApi.GetResponse().Interfaces, 2)
	assert.Equal(t, "InterfaceName", getAllEdgeInterfacesApi.GetResponse().Interfaces[0].Name)
}
