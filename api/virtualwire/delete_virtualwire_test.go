package virtualwire

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteVirtualWireApi *DeleteVirtualWiresApi

func setupDelete() {
	deleteVirtualWireApi = NewDelete("virtualwire-1")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteVirtualWireApi.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/2.0/vdn/virtualwires/virtualwire-1", deleteVirtualWireApi.Endpoint())
}
