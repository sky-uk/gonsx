package edgeinterface

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteEdgeInterfaceAPI *DeleteEdgeInterfaceAPI

func setupDelete() {
	deleteEdgeInterfaceAPI = NewDelete("edge-1", 1)
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteEdgeInterfaceAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces/1", deleteEdgeInterfaceAPI.Endpoint())
}
