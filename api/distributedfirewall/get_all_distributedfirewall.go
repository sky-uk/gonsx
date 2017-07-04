package distributedfirewall

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllFirewallRulesAPI default struct
type GetAllFirewallRulesAPI struct {
	*api.BaseAPI
}

// NewGetAll - Returns all the rules in the specified context
func NewGetAll(contextID string) *GetAllFirewallRulesAPI{
	this := new(GetAllFirewallRulesAPI)
	endpoint := "/api/4.0/firewall/"+contextID+"/config"
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, endpoint, nil, new(FirewallConfiguration))
	return this


}

// GetResponse - Returns ResponseObject from GetAllFirewallRulesAPI of Rule type.
func (getAllAPI GetAllFirewallRulesAPI) GetResponse() *FirewallConfiguration {
	return getAllAPI.ResponseObject().(*FirewallConfiguration)
}