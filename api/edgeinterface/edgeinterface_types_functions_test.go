package edgeinterface

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func setup() (edgeInterfacesList *EdgeInterfaces) {
	edgeInterfacesList = &EdgeInterfaces{}
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

func TestFilterByName(t *testing.T) {
	edgeInterfaces := setup()

	firstFiltered := edgeInterfaces.FilterByName("firstInterface")
	assert.Equal(t, "1", firstFiltered.Index)

	secondFiltered := edgeInterfaces.FilterByName("secondInterface")
	assert.Equal(t, "2", secondFiltered.Index)
}

func TestStringImplementation(t *testing.T) {
	edgeInterface := setup()
	expectedOutput := "[{{ } firstInterface  1500 internal true virtualwire-1 {[]} 1} {{ } secondInterface  1500 internal true virtualwire-1 {[]} 2}]"
	assert.Equal(t, expectedOutput, fmt.Sprint(edgeInterface))
}