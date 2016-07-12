package dhcprelay

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllDhcpRelaysApi *GetAllDhcpRelaysApi

func setupGetAll() {
	getAllDhcpRelaysApi = NewGetAll("edge-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllDhcpRelaysApi.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/dhcp/config/relay", getAllDhcpRelaysApi.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>17</vnicIndex><giAddress>10.152.164.1</giAddress></relayAgent><relayAgent><vnicIndex>16</vnicIndex><giAddress>10.152.165.1</giAddress></relayAgent></relayAgents></relay>")

	xmlerr := xml.Unmarshal(xmlContent, getAllDhcpRelaysApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllDhcpRelaysApi.GetResponse().RelayAgents, 2)
	assert.Equal(t, "10.152.160.10", getAllDhcpRelaysApi.GetResponse().RelayServer.IpAddress)
	assert.Equal(t, "17", getAllDhcpRelaysApi.GetResponse().RelayAgents[0].VnicIndex)
	assert.Equal(t, "10.152.164.1", getAllDhcpRelaysApi.GetResponse().RelayAgents[0].GiAddress)
}
