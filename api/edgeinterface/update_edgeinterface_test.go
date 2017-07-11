package edgeinterface

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateEdgeInterfaceAPI *UpdateEdgeInterfaceAPI

func setupUpdate() {
	edge := EdgeInterface{Name: "foo", Mtu: 1500}
	updateEdgeInterfaceAPI = NewUpdate("edge-1", 1, edge)
}

func TestUpdateMethod(t *testing.T) {
	setupUpdate()
	assert.Equal(t, http.MethodPut, updateEdgeInterfaceAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	setupUpdate()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces/1", updateEdgeInterfaceAPI.Endpoint())
}
