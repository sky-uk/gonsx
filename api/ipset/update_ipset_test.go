package ipset

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateIPSetAPI *UpdateIPSetAPI

func updateSetup(id string) {

	ipset := IPSet{
		Name:  "TEST_SG_1",
		Value: "blah",
	}
	updateIPSetAPI = NewUpdate(id, &ipset)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup("ipset-001")
	assert.Equal(t, http.MethodPut, updateIPSetAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup("ipset-001")
	assert.Equal(t, "/api/2.0/services/ipset/ipset-001", updateIPSetAPI.Endpoint())
}
