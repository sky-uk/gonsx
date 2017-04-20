package securitygroup

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createSecurityGroupAPI *CreateSecurityGroupAPI

func createSetup() {
	newDynamicCriteria1 := NewDynamicCriteria("OR", "VM.SECURITY_TAG", "ovp_test_app4", "contains")
	dynamicCriteriaList := []DynamicCriteria{*newDynamicCriteria1}

	newDynamicSet1 := NewDynamicSet("OR", dynamicCriteriaList)
	dynamicSetList := []DynamicSet{*newDynamicSet1}

	newDynamicMemberDefinition1 := NewDynamicMemberDefinition(dynamicSetList)

	createSecurityGroupAPI = NewCreate("globalroot-0", "OVP_sg_test", newDynamicMemberDefinition1)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createSecurityGroupAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/securitygroup/bulk/globalroot-0", createSecurityGroupAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<securitygroup><name>OVP_sg_test</name><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.SECURITY_TAG</key><criteria>contains</criteria><value>ovp_test_app4</value></dynamicCriteria></dynamicSet></dynamicMemberDefinition></securitygroup>"

	xmlBytes, err := xml.Marshal(createSecurityGroupAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
