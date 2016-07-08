package virtualwire

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"encoding/xml"
)

var createVirtualWireApi *CreateVirtualWireApi

func createSetup() {
	createVirtualWireApi = NewCreate("test", "test desc", "test", "vdnscope-1")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createVirtualWireApi.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/vdn/switches", createVirtualWireApi.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	object := new(VirtualWireCreateSpec)
	object.Name = "test"
	object.Description = "test desc"
	object.ControlPlaneMode = "UNICAST_MODE"
	object.TenantID = "test"
	expectedXml := "<virtualWireCreateSpec><name>test</name><controlPlaneMode>UNICAST_MODE</controlPlaneMode><description>test desc</description><tenantId>test</tenantId></virtualWireCreateSpec>"

	xmlBytes, err :=xml.Marshal(object)

	assert.Nil(t, err)
	assert.Equal(t, expectedXml, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}

