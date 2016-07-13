package virtualwire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterByName(t *testing.T) {
	virtualWires := new(VirtualWires)
	dataPage := new(DataPage)
	virtualWire1 := new(VirtualWire)
	virtualWire1.ObjectID = "id-1"
	virtualWire1.Name = "name-1"
	virtualWire2 := new(VirtualWire)
	virtualWire2.ObjectID = "id-2"
	virtualWire2.Name = "name-2"
	dataPage.VirtualWires = []VirtualWire{*virtualWire1, *virtualWire2}
	virtualWires.DataPage = *dataPage

	virtualWireFiltered := virtualWires.FilterByName("name-2")

	assert.Equal(t, "id-2", virtualWireFiltered.ObjectID)
}
