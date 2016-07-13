package edgeinterface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() (edgeInterfacesList *EdgeInterfaces) {
	edgeInterfacesList = &EdgeInterfaces{}
	firstInterface := EdgeInterface{
		Name:          "firstInterface",
		ConnectedToId: "virtualwire-1",
		Type:          "internal",
		Index:         "1",
		Mtu:           1500,
		IsConnected:   true,
	}
	secondInterface := EdgeInterface{
		Name:          "secondInterface",
		ConnectedToId: "virtualwire-1",
		Type:          "internal",
		Index:         "2",
		Mtu:           1500,
		IsConnected:   true,
	}
	edgeInterfacesList.Interfaces = []EdgeInterface{firstInterface, secondInterface}
	return edgeInterfacesList
}

func TestFilterByName(t *testing.T) {
	edgeInterfaces := setup()

	firstFiltered := edgeInterfaces.FilterByName("firstInterface")
	assert.Equal(t, "1", firstFiltered.Index)

	secondFiltered := edgeInterfaces.FilterByName("secondInterface")
	assert.Equal(t, "2", secondFiltered.Index)
}
