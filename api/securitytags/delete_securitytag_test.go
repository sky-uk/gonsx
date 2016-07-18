package securitytags

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteSecurityTagAPI *DeleteSecurityTagAPI

func setupDelete() {
	deleteSecurityTagAPI = NewDelete("securitytag-1")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteSecurityTagAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/2.0/services/securitytags/tag/securitytag-1", deleteSecurityTagAPI.Endpoint())
}
