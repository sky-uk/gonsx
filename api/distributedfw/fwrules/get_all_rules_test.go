package fwrules

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllRules *GetAllRulesAPI

//var allRules *Section

func setupGetAllRules() {
	getAllRules = NewGetAll("LAYER3", "1110")
}

func TestGetAllEndPoint(t *testing.T) {
	setupGetAllRules()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer3sections/1110", getAllRules.Endpoint())
}

func TestGetAllMethod(t *testing.T) {
	setupGetAllRules()
	assert.Equal(t, http.MethodGet, getAllRules.Method())
}

/*func TestGetAllResponse(t *testing.T) {
	setupGetAllRules()
	getAllRules.SetResponseObject("test response")
	assert.Equal(t, "test response", getAllRules.GetResponse())

}*/
