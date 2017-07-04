package virtualwire

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateVirtualWireAPI *UpdateVirtualWireAPI
var virtualWireUpdate VirtualWire

func updateSetup() {
	virtualWireUpdate.Name = "name1"
	virtualWireUpdate.Description = "desc1"
	virtualWireUpdate.ControlPlaneMode = "UNICAST_MODE"
	virtualWireUpdate.ObjectID = "virtualwire-1"
	updateVirtualWireAPI = NewUpdate(virtualWireUpdate)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateVirtualWireAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/vdn/virtualwires/"+virtualWireUpdate.ObjectID, updateVirtualWireAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	expectedXML := "<virtualWire><name>name1</name><objectId>virtualwire-1</objectId><controlPlaneMode>UNICAST_MODE</controlPlaneMode><description>desc1</description></virtualWire>"

	xmlBytes, err := xml.Marshal(virtualWireUpdate)

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))

}
