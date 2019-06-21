package firewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type GetFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewGetRule(sectionId int, etag string, ruleId int) *GetFirewallRuleAPI {
	this := new(GetFirewallRuleAPI)

	this.BaseAPI = api.NewBaseAPI(http.MethodGet,
		fmt.Sprintf("/api/4.0/firewall/globalroot-0/config/layer3sections/%d/rules/%d", sectionId, ruleId),
		nil, new(Rule))
	this.SetRequestHeader("If-Match", etag)

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca GetFirewallRuleAPI) GetResponse() *Rule {
	return ca.ResponseObject().(*Rule)
}
