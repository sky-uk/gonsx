package edgeinterface

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createEdgeInterfaceAPI *CreateEdgeInterfaceAPI

func createSetup() {
	createEdgeInterfaceAPI = NewCreate("edge-1", "interface_name", "virtualWire-1",
		"10.190.160.1", "255.255.255.0", "internal", 1500)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createEdgeInterfaceAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/4.0/edges/edge-1/interfaces/?action=patch", createEdgeInterfaceAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<interfaces><interface><name>interface_name</name><mtu>1500</mtu><type>internal</type><isConnected>true</isConnected><connectedToId>virtualWire-1</connectedToId><addressGroups><addressGroup><primaryAddress>10.190.160.1</primaryAddress><subnetMask>255.255.255.0</subnetMask></addressGroup></addressGroups></interface></interfaces>"

	xmlBytes, err := xml.Marshal(createEdgeInterfaceAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
