package edgeinterface

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteEdgeInterfaceApi *DeleteEdgeInterfaceApi

func setupDelete() {
	deleteEdgeInterfaceApi = NewDelete("1", "edge-1")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteEdgeInterfaceApi.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces/?index=1", deleteEdgeInterfaceApi.Endpoint())
}
