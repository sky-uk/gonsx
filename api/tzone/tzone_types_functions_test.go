package tzone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterByName(t *testing.T) {
	networkScopeList := new(NetworkScopeList)
	networkScope1 := new(NetworkScope)
	networkScope1.ObjectID = "id-1"
	networkScope1.Name = "name-1"
	networkScope2 := new(NetworkScope)
	networkScope2.ObjectID = "id-2"
	networkScope2.Name = "name-2"

	networkScopeList.NetworkScopeList = []NetworkScope{*networkScope1, *networkScope2}

	networkScopeFiltered := networkScopeList.FilterByName("name-2")
	assert.Equal(t, "id-2", networkScopeFiltered.ObjectID)
}
