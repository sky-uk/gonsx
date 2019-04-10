package nat

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createNatRuleAPI *CreateNatRuleAPI

func createSetup() {
	var rule Rule
	rule.Action = "dnat"

	createNatRuleAPI = NewCreateRule("edge-11", rule)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createNatRuleAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/4.0/edges/edge-11/nat/config/rules", createNatRuleAPI.Endpoint())
}
