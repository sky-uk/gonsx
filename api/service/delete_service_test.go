package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteServiceAPI *DeleteServiceAPI

func setupDelete() {
	deleteServiceAPI = NewDelete("application-200")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteServiceAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/2.0/services/application/application-200", deleteServiceAPI.Endpoint())
}
