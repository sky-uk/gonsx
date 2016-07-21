package securitytag

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var assignSecurityTagAPI *AssignSecurityTagAPI

func setupAssign() {
	assignSecurityTagAPI = NewAssign("securitytag-1", "vm-1")
}

func TestAssignMethod(t *testing.T) {
	setupAssign()
	assert.Equal(t, http.MethodPut, assignSecurityTagAPI.Method())
}

func TestAssignEndpoint(t *testing.T) {
	setupAssign()
	assert.Equal(t, "/api/2.0/services/securitytags/tag/securitytag-1/vm/vm-1", assignSecurityTagAPI.Endpoint())
}
