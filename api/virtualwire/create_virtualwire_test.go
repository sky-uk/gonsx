package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createVirtualWireAPI *CreateVirtualWireAPI
var createSpec CreateSpec

func createSetup() {
	createSpec.Name = "test"
	createSpec.ControlPlaneMode = "UNICAST_MODE"
	createSpec.Description = "test desc"
	createSpec.TenantID = "test"
	createVirtualWireAPI = NewCreate(createSpec, "vdnscope-1")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createVirtualWireAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/vdn/scopes/vdnscope-1/virtualwires", createVirtualWireAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	expectedXML := "<virtualWireCreateSpec><name>test</name><controlPlaneMode>UNICAST_MODE</controlPlaneMode><description>test desc</description><tenantId>test</tenantId></virtualWireCreateSpec>"

	xmlBytes, err := xml.Marshal(createSpec)

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
