package securitygroup

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteSecurityGroupAPI *DeleteSecurityGroupAPI

func setupDelete() {
	deleteSecurityGroupAPI = NewDelete("securitygroup-0001")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteSecurityGroupAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/2.0/services/securitygroup/securitygroup-0001", deleteSecurityGroupAPI.Endpoint())
}
