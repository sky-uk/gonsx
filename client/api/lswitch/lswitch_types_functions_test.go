package lswitch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFilterByName(t *testing.T) {
	vdsContextList := new(VdsContextList)
	vdsContext1 := new(VdsContext)
	vdsContext1.Switch.ObjectId = "id-1"
	vdsContext1.Switch.Name = "name-1"
	vdsContext2 := new(VdsContext)
	vdsContext2.Switch.ObjectId = "id-2"
	vdsContext2.Switch.Name = "name-2"
	vdsContextList.VdsContextList = []VdsContext {*vdsContext1, *vdsContext2}

	vdsContextFiltered := vdsContextList.FilterByName("name-2")

	assert.Equal(t, "id-2", vdsContextFiltered.Switch.ObjectId)
}