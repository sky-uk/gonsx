package dhcprelay

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateDhcpRelayApi *UpdateDhcpRelayApi

func createSetup() {
	firstRelayAgent := new(RelayAgent)
	firstRelayAgent.VnicIndex = "16"
	firstRelayAgent.GiAddress = "10.152.165.1"

	secondRelayAgent := new(RelayAgent)
	secondRelayAgent.VnicIndex = "17"
	secondRelayAgent.GiAddress = "10.152.164.1"

	relayAgentsList := []RelayAgent{*firstRelayAgent, *secondRelayAgent}
	updateDhcpRelayApi = NewCreate("10.152.160.10", "edge-50", relayAgentsList)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, updateDhcpRelayApi.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/4.0/edges/edge-50/dhcp/config/relay", updateDhcpRelayApi.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXml := "<relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>16</vnicIndex><giAddress>10.152.165.1</giAddress></relayAgent><relayAgent><vnicIndex>17</vnicIndex><giAddress>10.152.164.1</giAddress></relayAgent></relayAgents></relay>"

	xmlBytes, err := xml.Marshal(updateDhcpRelayApi.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXml, string(xmlBytes))

}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
