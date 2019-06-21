package firewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type CreateFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewCreateRule(sectionId int, etag string, rule *Rule) *CreateFirewallRuleAPI {
	this := new(CreateFirewallRuleAPI)
	rule.SectionId = sectionId
	if rule.PacketType == "" {
		rule.PacketType = "any"
	}
	if rule.Action == "" {
		rule.Action = Allow
	}
	if rule.AppliedToList == nil || rule.AppliedToList.Elements == nil || len(rule.AppliedToList.Elements) == 0 {
		rule.AppliedToList = &AppliedToList{
			Elements: []Element{
				{
					Type:  DISTRIBUTED_FIREWALL,
					Value: string(DISTRIBUTED_FIREWALL),
				},
			},
		}
	}

	this.BaseAPI = api.NewBaseAPI(http.MethodPost,
		fmt.Sprintf("/api/4.0/firewall/globalroot-0/config/layer3sections/%d/rules", sectionId),
		rule, new(Rule))
	this.SetRequestHeader("If-Match", etag)

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca CreateFirewallRuleAPI) GetResponse() *Rule {
	return ca.ResponseObject().(*Rule)
}
