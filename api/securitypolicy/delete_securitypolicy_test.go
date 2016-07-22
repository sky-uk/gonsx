package securitypolicy

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteSecurityPolicyAPI *DeleteSecurityPolicyAPI

func setupDelete() {
	deleteSecurityPolicyAPI = NewDelete("securitypolicy-01", false)
}

func setupForcedDelete() {
	deleteSecurityPolicyAPI = NewDelete("securitypolicy-01", true)
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteSecurityPolicyAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/2.0/services/policy/securitypolicy/securitypolicy-01?force=false", deleteSecurityPolicyAPI.Endpoint())
	setupForcedDelete()
	assert.Equal(t, "/api/2.0/services/policy/securitypolicy/securitypolicy-01?force=true", deleteSecurityPolicyAPI.Endpoint())

}
