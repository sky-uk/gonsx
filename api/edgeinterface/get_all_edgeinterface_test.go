package edgeinterface

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllEdgeInterfacesAPI *GetAllEdgeInterfacesAPI

func setupGetAll() {
	getAllEdgeInterfacesAPI = NewGetAll("edge-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllEdgeInterfacesAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces", getAllEdgeInterfacesAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<interfaces><interface><name>InterfaceName</name></interface><interface><name>InterfaceName2</name></interface></interfaces>")

	xmlerr := xml.Unmarshal(xmlContent, getAllEdgeInterfacesAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllEdgeInterfacesAPI.GetResponse().Interfaces, 2)
	assert.Equal(t, "InterfaceName", getAllEdgeInterfacesAPI.GetResponse().Interfaces[0].Name)
}
