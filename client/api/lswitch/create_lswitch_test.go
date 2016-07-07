package lswitch

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"encoding/xml"
)

var createLogicalSwitchApi *CreateLogicalSwitchApi

func createSetup() {
	createLogicalSwitchApi = NewCreate()
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createLogicalSwitchApi.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/vdn/switches", createLogicalSwitchApi.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	object := new(VdsContext)
	object.Switch.ObjectId = "objectId"
	object.Switch.Name = "name"
	object.Switch.Type.TypeName = "typeName"
	object.Switch.ObjectTypeName = "objectTypeName"
	object.Switch.Revision = 99
	object.Teaming = "TEAMING"
	object.MTU = 1600
	expectedXml := "<vdsContext><switch><objectId>objectId</objectId><name>name</name><type><typeName>typeName</typeName></type><revision>99</revision><objectTypeName>objectTypeName</objectTypeName></switch><teaming>TEAMING</teaming><mtu>1600</mtu></vdsContext>"

	xmlBytes, err :=xml.Marshal(object)

	assert.Nil(t, err)
	assert.Equal(t, expectedXml, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}

