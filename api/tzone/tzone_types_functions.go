package tzone

import "fmt"

func (s NetworkScopeList) String() string {
	return fmt.Sprintf("%s", s.NetworkScopeList)
}

func (s NetworkScope) String() string {
	return fmt.Sprintf("id: %s, name: %s", s.ObjectId, s.Name)
}

func (v NetworkScopeList) FilterByName(name string) *NetworkScope {
	var networkScopeFound NetworkScope
	for _, networkScope := range v.NetworkScopeList {
		if(networkScope.Name == name) {
			networkScopeFound = networkScope
			break
		}
	}
	return &networkScopeFound
}
