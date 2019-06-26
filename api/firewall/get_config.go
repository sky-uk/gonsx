package firewall

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type GetFirewallConfig struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewGetFirewallConfig() *GetFirewallConfig {
	this := new(GetFirewallConfig)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/firewall/globalroot-0/config", nil, new(FirewallConfiguration))

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca GetFirewallConfig) GetResponse() string {
	return ca.ResponseObject().(string)
}
