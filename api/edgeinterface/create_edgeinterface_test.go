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

func createObject() *EdgeInterfaces {
	edgeInterfacesList := new(EdgeInterfaces)
	firstInterface := EdgeInterface{
		Name:          "firstInterface",
		ConnectedToID: "virtualwire-1",
		Type:          "internal",
		Index:         "1",
		Mtu:           1500,
		IsConnected:   true,
	}
	secondInterface := EdgeInterface{
		Name:          "secondInterface",
		ConnectedToID: "virtualwire-1",
		Type:          "internal",
		Index:         "2",
		Mtu:           1500,
		IsConnected:   true,
	}
	edgeInterfacesList.Interfaces = []EdgeInterface{firstInterface, secondInterface}
	return edgeInterfacesList
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
