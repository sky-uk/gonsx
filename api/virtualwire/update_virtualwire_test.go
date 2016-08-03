package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateVirtualWireAPI *UpdateVirtualWireAPI

func updateSetup() {
	updateVirtualWireAPI = NewUpdate("name2", "desc2", "virtualwire-1")
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateVirtualWireAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/vdn/virtualwires/virtualwire-1", updateVirtualWireAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	object := new(VirtualWire)
	object.Name = "name1"
	object.Description = "desc1"
	object.ControlPlaneMode = "UNICAST_MODE"
	expectedXML := "<virtualWire><name>name1</name><controlPlaneMode>UNICAST_MODE</controlPlaneMode><description>desc1</description></virtualWire>"

	xmlBytes, err := xml.Marshal(object)

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))

}
