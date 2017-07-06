package dhcprelay

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// UpdateDhcpRelayAPI struct
type UpdateDHCPRelayAPI struct {
	*api.BaseAPI
}

// NewUpdate creates a new object of UpdateDhcpRelayAPI
func NewUpdate(edgeID string, dhcpRelay DhcpRelay) *UpdateDHCPRelayAPI {
	this := new(UpdateDHCPRelayAPI)
	requestPayload := dhcpRelay
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/4.0/edges/"+edgeID+"/dhcp/config/relay", requestPayload, new(string))
	return this
}

// NewCreate pseudo definition of NewCreate, returns an object of UpdateDhcpRelayAPI
func NewCreate(edgeID string, dhcpRelay DhcpRelay) *UpdateDHCPRelayAPI {
	return NewUpdate(edgeID, dhcpRelay)
}

// GetResponse returns the ResponseObject from UpdateDhcpRelayAPI
func (updateAPI UpdateDHCPRelayAPI) GetResponse() string {
	return updateAPI.ResponseObject().(string)
}
