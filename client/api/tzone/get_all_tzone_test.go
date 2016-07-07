package tzone

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"net/http"
)

var getAllTransportZonesApi *GetAllTransportZonesApi

func setupGetAll() {
	getAllTransportZonesApi = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllTransportZonesApi.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/vdn/scopes", getAllTransportZonesApi.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<vdnScopes><vdnScope><objectId>vdnscope-1</objectId></vdnScope></vdnScopes>")

	xmlerr := xml.Unmarshal(xmlContent, getAllTransportZonesApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllTransportZonesApi.GetResponse().NetworkScopeList, 1)
	assert.Equal(t, "vdnscope-1", getAllTransportZonesApi.GetResponse().NetworkScopeList[0].ObjectId)
}
