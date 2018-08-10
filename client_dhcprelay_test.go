package gonsx

import (
	"github.com/tadaweb/gonsx/api/dhcprelay"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClientGetAllDhcpRelays(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>17</vnicIndex><giAddress>10.152.164.1</giAddress></relayAgent><relayAgent><vnicIndex>16</vnicIndex><giAddress>10.152.165.1</giAddress></relayAgent></relayAgents></relay>`
	setup(http.StatusOK, xmlContent)
	defer server.Close()

	api := dhcprelay.NewGetAll("edge-50")
	nsxClient.Do(api)

	assert.Equal(t, []string{"10.152.160.10"}, api.GetResponse().RelayServer.IPAddress)
	assert.Len(t, api.GetResponse().RelayAgents, 2)
	assert.Equal(t, "17", api.GetResponse().RelayAgents[0].VnicIndex)
	assert.Equal(t, "10.152.164.1", api.GetResponse().RelayAgents[0].GiAddress)
	assert.Equal(t, "16", api.GetResponse().RelayAgents[1].VnicIndex)
	assert.Equal(t, "10.152.165.1", api.GetResponse().RelayAgents[1].GiAddress)

}
