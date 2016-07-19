package dhcprelay

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func setup() *DhcpRelay {
	dhcpRelay := &DhcpRelay{}
	firstRelayAgent := RelayAgent{
		VnicIndex: "1",
		GiAddress: "10.10.10.1",
	}
	secondRelayAgent := RelayAgent{
		VnicIndex: "2",
		GiAddress: "10.10.10.2",
	}

	dhcpRelay.RelayAgents = []RelayAgent{firstRelayAgent, secondRelayAgent}
	dhcpRelay.RelayServer.IPAddress = "10.10.10.10"
	return dhcpRelay
}

func TestFilterByIpAddress(t *testing.T) {
	dhcpRelay := setup()

	firstFiltered := dhcpRelay.FilterByIPAddress("10.10.10.1")
	assert.Equal(t, "1", firstFiltered.VnicIndex)

	secondFiltered := dhcpRelay.FilterByIPAddress("10.10.10.2")
	assert.Equal(t, "2", secondFiltered.VnicIndex)
}

func TestFilterByVnicIndex(t *testing.T) {
	dhcpRelay := setup()

	firstFiltered := dhcpRelay.FilterByVnicIndex("1")
	assert.Equal(t, "10.10.10.1", firstFiltered.GiAddress)

	secondFiltered := dhcpRelay.FilterByVnicIndex("2")
	assert.Equal(t, "10.10.10.2", secondFiltered.GiAddress)
}

func TestCheckByVnicIndex(t *testing.T) {
	dhcpRelay := setup()
	assert.Equal(t, true, dhcpRelay.CheckByVnicIndex("1"))
	assert.Equal(t, false, dhcpRelay.CheckByVnicIndex("10"))
}

func TestRemoveByVnicIndex(t *testing.T) {
	dhcpRelay := setup()
	assert.Equal(t, true, dhcpRelay.CheckByVnicIndex("1"))

	newDhcpRelay := dhcpRelay.RemoveByVnicIndex("1")
	assert.Equal(t, false, newDhcpRelay.CheckByVnicIndex("1"))
}


func TestStringImplementation(t *testing.T) {
	dhcpRelay := setup()

	relay_server_string := fmt.Sprintln(dhcpRelay.RelayServer)
	assert.Equal(t, "DhcpRelayServer ipAddress:10.10.10.10.\n", relay_server_string)

	relay_agents_string := fmt.Sprintln(dhcpRelay.RelayAgents[0])
	assert.Equal(t, "DhcpRelayAgent VnicIndex:1, GiAddress:10.10.10.1.\n", relay_agents_string)

}
