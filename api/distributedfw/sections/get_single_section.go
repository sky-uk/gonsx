package sections

import (
	"github.com/sky-uk/gonsx/api"
	//"net/http"
)

// GetSingleSectionAPI default struct
type GetSingleSectionAPI struct {
	*api.BaseAPI
}

// NewGetSingle - Returns all the rules in the specified context
func NewGetSingle(sectionID, sectionType string) *GetSingleSectionAPI {
	this := new(GetSingleSectionAPI)
	/*
		var endpoint string
		switch ruleType {
		case "LAYER3":
			endpoint = "/api/4.0/firewall/globalroot-0/config/layer3sections/" + ruleSection + "/rules/" + ruleID

		case "LAYER2":
			endpoint = "/api/4.0/firewall/globalroot-0/config/layer2sections/" + ruleSection + "/rules/" + ruleID
		}

		this.BaseAPI = api.NewBaseAPI(http.MethodGet, endpoint, nil, new(Section))
	*/
	return this

}

// GetResponse - Returns ResponseObject from GetAllFirewallRulesAPI of Rule type.
func (getSingleAPI GetSingleSectionAPI) GetResponse() *Section {
	return getSingleAPI.ResponseObject().(*Section)
}
