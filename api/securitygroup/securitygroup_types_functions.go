package securitygroup

import (
	"fmt"
)

func (sgl List) String() string {
	return fmt.Sprintf("SecurityGroupList object, contains SecurityGroup objects.")
}

func (sg SecurityGroup) String() string {
	return fmt.Sprintf("objectId: %-20s name: %-20s.", sg.ObjectID, sg.Name)
}

// FilterByName returns a single securitygroup object if it matches the name in SecurityGroup
func (sgl List) FilterByName(name string) *SecurityGroup {
	var securityGroupFound SecurityGroup
	for _, sg := range sgl.SecurityGroups {
		if sg.Name == name {
			securityGroupFound = sg
			break
		}
	}
	return &securityGroupFound
}

// AddDynamicSet adds new DynamicSet to DynamicMemberDefinition of SecurityGroup object.
func (dynMemberDef *DynamicMemberDefinition) AddDynamicSet(operator string, dynamicCriteriaList []DynamicCriteria) {
	newDynamicSet := DynamicSet{Operator: operator, DynamicCriteria: dynamicCriteriaList}
	dynMemberDef.DynamicSet = append(dynMemberDef.DynamicSet, newDynamicSet)
}
