package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
)

var api *NSXApi

func setup() {
	api = NewGetTransportZone()
}

func TestMethod(t *testing.T) {
	setup()
	assert.Equal(t, "GET", api.NSXRequest.Method())
}

func TestEndpoint(t *testing.T) {
	setup()
	assert.Equal(t, "/api/2.0/vdn/scopes", api.NSXRequest.Endpoint())
}

func TestUnMarshalling(t *testing.T) {
	setup()
	xmlContent := []byte("<blah></blah>")
	xmlerr := xml.Unmarshal(xmlContent, &api.NSXResponse)
	assert.Nil(t, xmlerr)
	//assert.Equal(t, "blah", api.NSXResponse)
}

// TODO: add failure scenarios