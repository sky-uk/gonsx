package fwrules

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateFWRulesAPI default struct
type CreateFWRulesAPI struct {
	*api.BaseAPI
}

// NewGetAll - Returns all the rules in the specified context
func NewCreate(newRule Rule) *CreateFWRulesAPI {
	this := new(CreateFWRulesAPI)
	var endpoint string
	switch newRule.RuleType {
	case "LAYER3":
		endpoint = "/api/4.0/firewall/globalroot-0/config/layer3sections/"+string(newRule.SectionID)+"/rules/"

	case "LAYER2":
		endpoint = "/api/4.0/firewall/globalroot-0/config/layer2sections/"+string(newRule.SectionID)+"/rules/"
	}

	this.BaseAPI = api.NewBaseAPI(http.MethodPost, endpoint, newRule, new(Rule))
	return this
}

// GetResponse - Returns ResponseObject from GetAllFirewallRulesAPI of Rule type.
func (createAPI CreateFWRulesAPI) GetResponse() *Rule {
	return createAPI.ResponseObject().(*Rule)
}
