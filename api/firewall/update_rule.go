package firewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type UpdateFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewUpdateRule(sectionId int, etag string, ruleId int, rule Rule) *UpdateFirewallRuleAPI {
	this := new(UpdateFirewallRuleAPI)
	rule.SectionId = sectionId
	rule.ID = ruleId
	if rule.PacketType == "" {
		rule.PacketType = "any"
	}
	if rule.Action == "" {
		rule.Action = Allow
	}

	this.BaseAPI = api.NewBaseAPI(http.MethodPut,
		fmt.Sprintf("/api/4.0/firewall/globalroot-0/config/layer3sections/%d/rules/%d", sectionId, ruleId),
		rule, new(Rule))
	this.SetRequestHeader("If-Match", etag)

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca UpdateFirewallRuleAPI) GetResponse() *Rule {
	return ca.ResponseObject().(*Rule)
}
