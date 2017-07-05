package ipset

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateIpSetAPI *UpdateIpSetAPI

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

	ipset := IpSet{
		Name:     "TEST_SG_1",
		ObjectID: "ipset-001",
		DynamicMemberDefinition: &DynamicMemberDefinition{
			DynamicSet: dynamicSetList,
		},
	}
	updateIpSetAPI = NewUpdate("ipset-001", &ipset)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateIpSetAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/ipset/bulk/ipset-001", updateIpSetAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<ipset><objectId>ipset-001</objectId><name>TEST_SG_1</name><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.NAME</key><criteria>contains</criteria><value>test_vm</value></dynamicCriteria></dynamicSet></dynamicMemberDefinition></ipset>"

	xmlBytes, err := xml.Marshal(updateIpSetAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}
