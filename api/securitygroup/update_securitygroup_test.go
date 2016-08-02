package securitygroup

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateSecurityGroupAPI *UpdateSecurityGroupAPI

func updateSetup() {
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

	securityGroup := SecurityGroup{
		Name:     "TEST_SG_1",
		ObjectID: "securitygroup-001",
		DynamicMemberDefinition: &DynamicMemberDefinition{
			DynamicSet: dynamicSetList,
		},
	}
	updateSecurityGroupAPI = NewUpdate("securitygroup-001", &securityGroup)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateSecurityGroupAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/securitygroup/bulk/securitygroup-001", updateSecurityGroupAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<securitygroup><objectId>securitygroup-001</objectId><name>TEST_SG_1</name><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.NAME</key><criteria>contains</criteria><value>test_vm</value></dynamicCriteria></dynamicSet></dynamicMemberDefinition></securitygroup>"

	xmlBytes, err := xml.Marshal(updateSecurityGroupAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}
