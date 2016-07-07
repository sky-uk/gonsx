package lswitch

import "fmt"

func (s VdsContextList) String() string {
	return fmt.Sprintf("%s", s.VdsContextList)
}

func (s VdsContext) String() string {
	return fmt.Sprintf("id: %s, name: %s", s.Switch.ObjectId, s.Switch.Name)
}

func (v VdsContextList) FilterByName(name string) *VdsContext {
	var vdsContextFound VdsContext
	for _, vdsContext := range v.VdsContextList {
		if(vdsContext.Switch.Name == name) {
			vdsContextFound = vdsContext
			break
		}
	}
	return &vdsContextFound
}
