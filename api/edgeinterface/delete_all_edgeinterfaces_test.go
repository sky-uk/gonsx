package edgeinterface

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteAllEdgeInterfaceAPI *DeleteAllEdgeInterfaceAPI

func setupDeleteAll() {
	deleteAllEdgeInterfaceAPI = NewDeleteAll("edge-1")
}

func TestDeleteAllMethod(t *testing.T) {
	setupDeleteAll()
	assert.Equal(t, http.MethodDelete, deleteAllEdgeInterfaceAPI.Method())
}

func TestDeleteAllEndpoint(t *testing.T) {
	setupDeleteAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces", deleteAllEdgeInterfaceAPI.Endpoint())
}
