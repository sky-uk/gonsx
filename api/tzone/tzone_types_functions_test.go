package tzone

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFilterByName(t *testing.T) {
	networkScopeList := new(NetworkScopeList)
	networkScope1 := new(NetworkScope)
	networkScope1.ObjectId = "id-1"
	networkScope1.Name = "name-1"
	networkScope2 := new(NetworkScope)
	networkScope2.ObjectId = "id-2"
	networkScope2.Name = "name-2"

	networkScopeList.NetworkScopeList = []NetworkScope { *networkScope1, *networkScope2}

	networkScopeFiltered := networkScopeList.FilterByName("name-2")
	assert.Equal(t, "id-2", networkScopeFiltered.ObjectId)
}