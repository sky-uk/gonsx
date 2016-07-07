package lswitch

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"net/http"
)

var getAllLogicalSwitchesApi *GetAllLogicalSwitchesApi

func setupGetAll() {
	getAllLogicalSwitchesApi = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllLogicalSwitchesApi.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/vdn/switches", getAllLogicalSwitchesApi.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<vdsContexts><vdsContext><switch><objectId>dvs-1</objectId></switch></vdsContext></vdsContexts>")

	xmlerr := xml.Unmarshal(xmlContent, getAllLogicalSwitchesApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllLogicalSwitchesApi.GetResponse().VdsContextList, 1)
	assert.Equal(t, "dvs-1", getAllLogicalSwitchesApi.GetResponse().VdsContextList[0].Switch.ObjectId)
}
