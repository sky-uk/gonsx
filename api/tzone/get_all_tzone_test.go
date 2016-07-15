package tzone

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllTransportZonesAPI *GetAllTransportZonesAPI

func setupGetAll() {
	getAllTransportZonesAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllTransportZonesAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/vdn/scopes", getAllTransportZonesAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<vdnScopes><vdnScope><objectId>vdnscope-1</objectId></vdnScope></vdnScopes>")

	xmlerr := xml.Unmarshal(xmlContent, getAllTransportZonesAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllTransportZonesAPI.GetResponse().NetworkScopeList, 1)
	assert.Equal(t, "vdnscope-1", getAllTransportZonesAPI.GetResponse().NetworkScopeList[0].ObjectID)
}
