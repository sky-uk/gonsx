package securitygroup

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() (securityGrouplist *List) {
	dynamicCriteria := DynamicCriteria{
		Operator: "OR",
		Key:      "VM.NAME",
		Value:    "test_vm",
		Criteria: "contains",
	}
	dynamicCriteriaList := []DynamicCriteria{dynamicCriteria}

	dynamicSet := DynamicSet{
		Operator:        "OR",
		DynamicCriteria: dynamicCriteriaList,
	}
	dynamicSetList := []DynamicSet{dynamicSet}

	securityGrouplist = &List{}
	first := SecurityGroup{
		Name:     "TEST_SG_1",
		ObjectID: "securitygroup-001",
		DynamicMemberDefinition: &DynamicMemberDefinition{
			DynamicSet: dynamicSetList,
		},
	}
	second := SecurityGroup{
		Name:     "TEST_SG_2",
		ObjectID: "securitygroup-002",
		DynamicMemberDefinition: &DynamicMemberDefinition{
			DynamicSet: dynamicSetList,
		},
	}
	securityGrouplist.SecurityGroups = []SecurityGroup{first, second}
	return securityGrouplist
}

func TestFilterByName(t *testing.T) {
	securityGroupList := setup()

	firstFiltered := securityGroupList.FilterByName("TEST_SG_1")
	assert.Equal(t, "securitygroup-001", firstFiltered.ObjectID)

	secondFiltered := securityGroupList.FilterByName("TEST_SG_2")
	assert.Equal(t, "securitygroup-002", secondFiltered.ObjectID)
}

func TestStringImplementation(t *testing.T) {
	securityGroupList := setup()

	stringOutputOfList := fmt.Sprintln(securityGroupList)
	assert.Equal(t, "SecurityGroupList object, contains SecurityGroup objects.\n", stringOutputOfList)

	stringOutputOfObject := fmt.Sprintln(securityGroupList.SecurityGroups[0])
	assert.Equal(t, "objectId: securitygroup-001    name: TEST_SG_1           .\n", stringOutputOfObject)

}

func TestAddDynamicMemberDefinitionSet(t *testing.T) {
	securityGroupList := setup()
	securityGroup := securityGroupList.SecurityGroups[0]

	assert.Len(t, securityGroup.DynamicMemberDefinition.DynamicSet, 1)

	newDynamicCriteria := DynamicCriteria{
		Operator: "ADD",
		Key:      "VM.NAME",
		Value:    "test_vm_name2",
		Criteria: "contains",
	}
	newDynamicCriteriaList := []DynamicCriteria{newDynamicCriteria}
	securityGroup.AddDynamicMemberDefinitionSet("OR", newDynamicCriteriaList)
	assert.Len(t, securityGroup.DynamicMemberDefinition.DynamicSet, 2)
}
