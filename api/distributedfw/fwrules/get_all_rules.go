package fwrules

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllRulesAPI default struct
type GetAllRulesAPI struct {
	*api.BaseAPI
}

// NewGetAll - Returns all the rules in the specified context
func NewGetAll(ruleType, section string) *GetAllRulesAPI {
	this := new(GetAllRulesAPI)
	var endpoint string
	if ruleType == "LAYER3" {
		endpoint = "/api/4.0/firewall/globalroot-0/config/layer3sections/" + section
	}

	this.BaseAPI = api.NewBaseAPI(http.MethodGet, endpoint, nil, new(Section))
	return this

}

// GetResponse - Returns ResponseObject from GetAllFirewallRulesAPI of Rule type.
func (getAllAPI GetAllRulesAPI) GetResponse() *Section {
	return getAllAPI.ResponseObject().(*Section)
}
