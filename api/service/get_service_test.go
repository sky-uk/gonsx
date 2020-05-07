package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getServiceAPI *GetServiceAPI

func setupGet() {
	getServiceAPI = NewGet("application-5")
}

func TestGetMethod(t *testing.T) {
	setupGet()
	assert.Equal(t, http.MethodGet, getServiceAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet()
	assert.Equal(t, "/api/2.0/services/application/application-5", getServiceAPI.Endpoint())
}
