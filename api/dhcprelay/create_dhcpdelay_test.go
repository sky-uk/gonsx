package dhcprelay

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

)


var createDhcpRelayAPI *CreateDhcpRelayAPI

func setupCreateRelay() {
	var createDhcp DhcpRelay
	createDhcpRelayAPI = NewCreate("edge-5",createDhcp)
}

func TestCreateRelayMethod(t *testing.T) {
	setupCreateRelay()
	assert.Equal(t, http.MethodPut, createDhcpRelayAPI.Method())

}