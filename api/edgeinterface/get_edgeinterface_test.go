package edgeinterface

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getEdgeInterfaceAPI *GetEdgeInterfaceAPI

func setupGet() {
	getEdgeInterfaceAPI = NewGet("edge-1", 10)
}

func TestGetMethod(t *testing.T) {
	setupGet()
	assert.Equal(t, http.MethodGet, getEdgeInterfaceAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces/10", getEdgeInterfaceAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet()
	xmlContent := []byte("<interface><name>InterfaceName</name></interface>")

	xmlerr := xml.Unmarshal(xmlContent, getEdgeInterfaceAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "InterfaceName", getEdgeInterfaceAPI.GetResponse().Name)
}
