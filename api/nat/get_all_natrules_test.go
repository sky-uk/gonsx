package nat

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllNatRulesAPI *GetAllNatRulesAPI

func setupGetAll() {
	getAllNatRulesAPI = NewGetAll("edge-11")
}

func TestGatAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllNatRulesAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-11/nat/config", getAllNatRulesAPI.Endpoint())
}
