package fwrules

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteFWRuleAPI *DeleteFWRuleAPI
var deletedRule Rule

func setupDeleteL3Rule() {
	deletedRule.SectionID = 1234
	deletedRule.RuleType = "LAYER3"
	deleteFWRuleAPI = NewDelete(deletedRule)
}

func setupDeleteL2Rule() {
	deletedRule.SectionID = 1234
	deletedRule.RuleType = "LAYER2"
	deleteFWRuleAPI = NewDelete(deletedRule)
}

func TestDeleteMethod(t *testing.T) {
	setupDeleteL3Rule()
	assert.Equal(t, http.MethodDelete, deleteFWRuleAPI.Method())

}

func TestL3Endpoint(t *testing.T) {
	setupDeleteL3Rule()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer3sections", deleteFWRuleAPI.Endpoint())
}

func TestL2Endpoint(t *testing.T) {
	setupDeleteL2Rule()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer2sections", deleteFWRuleAPI.Endpoint())
}


func TestL3RuleType(t *testing.T) {
	setupDeleteL3Rule()
	assert.Equal(t, "LAYER3", deletedRule.RuleType)
}