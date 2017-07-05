package fwrules

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createFWRuleAPI *CreateFWRulesAPI
var newRule Rule

func setupCreateRule() {
	newRule.SectionID = 1234
	newRule.RuleType = "LAYER3"
	createFWRuleAPI = NewCreate(newRule)
}

func setupCreateL2Rule() {
	newRule.SectionID = 1234
	newRule.RuleType = "LAYER2"
	createFWRuleAPI = NewCreate(newRule)
}

func TestCreateMethod(t *testing.T) {
	setupCreateRule()
	assert.Equal(t, http.MethodPost, createFWRuleAPI.Method())
}

func TestCreateL3Endpoint(t *testing.T) {
	setupCreateRule()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer3sections/1234/rules/", createFWRuleAPI.Endpoint())
}

func TestCreateL2Endpoint(t *testing.T) {
	setupCreateL2Rule()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer2sections/1234/rules/", createFWRuleAPI.Endpoint())
}

func TestCreateResponse(t *testing.T) {
	setupCreateRule()
	createFWRuleAPI.SetResponseObject(&newRule)
	assert.Equal(t, &newRule, createFWRuleAPI.GetResponse())
}

func TestRuleLayer3Type(t *testing.T) {
	setupCreateRule()
	assert.Equal(t, "LAYER3", newRule.RuleType)
}

func TestRuleLayer2Type(t *testing.T) {
	setupCreateRule()
	newRule.RuleType = "LAYER2"
	assert.Equal(t, "LAYER2", newRule.RuleType)
}

func TestSectionID(t *testing.T) {
	setupCreateRule()
	assert.Equal(t, 1234, newRule.SectionID)
}
