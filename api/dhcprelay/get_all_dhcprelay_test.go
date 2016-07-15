package dhcprelay

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllDhcpRelaysAPI *GetAllDhcpRelaysAPI

func setupGetAll() {
	getAllDhcpRelaysAPI = NewGetAll("edge-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllDhcpRelaysAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/dhcp/config/relay", getAllDhcpRelaysAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>17</vnicIndex><giAddress>10.152.164.1</giAddress></relayAgent><relayAgent><vnicIndex>16</vnicIndex><giAddress>10.152.165.1</giAddress></relayAgent></relayAgents></relay>")

	xmlerr := xml.Unmarshal(xmlContent, getAllDhcpRelaysAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllDhcpRelaysAPI.GetResponse().RelayAgents, 2)
	assert.Equal(t, "10.152.160.10", getAllDhcpRelaysAPI.GetResponse().RelayServer.IPAddress)
	assert.Equal(t, "17", getAllDhcpRelaysAPI.GetResponse().RelayAgents[0].VnicIndex)
	assert.Equal(t, "10.152.164.1", getAllDhcpRelaysAPI.GetResponse().RelayAgents[0].GiAddress)
}
