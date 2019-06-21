package firewall

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityPolicyAPI api object
type DeleteFirewallRuleAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreatePolicyAPI.
func NewDeleteRule(sectionId int, etag string, ruleId int) *DeleteFirewallRuleAPI {
	this := new(DeleteFirewallRuleAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, fmt.Sprintf(
		"/api/4.0/firewall/globalroot-0/config/layer3sections/%d/rules/%d", sectionId, ruleId,
	), nil, nil)
	this.SetRequestHeader("If-Match", etag)

	return this
}

// GetResponse returns a ResponseObject of CreateServiceAPI.
func (ca DeleteFirewallRuleAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
