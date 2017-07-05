package ipset

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteIpSetAPI *DeleteIpSetAPI

func setupDelete() {
	deleteIpSetAPI = NewDelete("ipset-0001")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteIpSetAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/2.0/services/ipset/ipset-0001", deleteIpSetAPI.Endpoint())
}
