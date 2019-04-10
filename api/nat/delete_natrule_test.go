package nat

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteNatRuleAPI *DeleteNatRuleAPI

func setupDelete() {
	deleteNatRuleAPI = NewDelete("edge-50", "1234")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteNatRuleAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/4.0/edges/edge-50/nat/config/rules/1234", deleteNatRuleAPI.Endpoint())
}
