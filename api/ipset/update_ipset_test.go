package ipset

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateIpSetAPI *UpdateIpSetAPI

func updateSetup(id string) {

	ipset := IpSet{
		Name:     "TEST_SG_1",
		Value: "blah",
	}
	updateIpSetAPI = NewUpdate(id, &ipset)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup("ipset-001")
	assert.Equal(t, http.MethodPut, updateIpSetAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup("ipset-001")
	assert.Equal(t, "/api/2.0/services/ipset/ipset-001", updateIpSetAPI.Endpoint())
}
