package dhcprelay

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createDhcpRelayAPI *CreateDhcpRelayAPI

func setupCreateRelay() {
	var createDhcp DhcpRelay
	createDhcpRelayAPI = NewCreate("edge-5", createDhcp)
}

func TestCreateRelayMethod(t *testing.T) {
	setupCreateRelay()
	assert.Equal(t, http.MethodPut, createDhcpRelayAPI.Method())

}

func TestCreateRelayEndpoint(t *testing.T){
	setupCreateRelay()
	assert.Equal(t,"/api/4.0/edges/edge-5/dhcp/config/relay", createDhcpRelayAPI.Endpoint())
}


func TestCreateRelayRespons(t *testing.T) {
	setupCreateRelay()
	createDhcpRelayAPI.SetResponseObject("example response")
	assert.Equal(t, "example response", createDhcpRelayAPI.GetResponse())
}