package fwrules

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteFWRuleAPI *DeleteFWRuleAPI
var deletedRule Rule

func setupDeleteRule() {
	newRule.SectionID = 1234
	newRule.RuleType = "LAYER3"
	deleteFWRuleAPI = NewDelete(deletedRule)
}

func TestDeleteMethod(t *testing.T) {
	setupDeleteRule()

}
