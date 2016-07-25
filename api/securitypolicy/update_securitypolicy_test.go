package securitypolicy

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateSecurityPolicyAPI *UpdateSecurityPolicyAPI

func updateSetup() {
	securityGroupIDs := []string{"securitygroup-0001", "securitygroup-0002"}

	requestPayload := new(SecurityPolicy)
	requestPayload.Name = "OVP_test_security_policy"
	requestPayload.Precedence = "5501"
	requestPayload.Description = "this is a long description."
	requestPayload.SecurityGroupBinding = []SecurityGroup{}

	var securityGroupBindingList = []SecurityGroup{}
	for _, secGroupID := range securityGroupIDs {
		securityGroupBinding := SecurityGroup{ObjectID: secGroupID}
		securityGroupBindingList = append(securityGroupBindingList, securityGroupBinding)
	}
	requestPayload.SecurityGroupBinding = securityGroupBindingList
	requestPayload.ActionsByCategory = ActionsByCategory{Actions: []Action{}}
	updateSecurityPolicyAPI = NewUpdate("securitypolicy-001", requestPayload)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateSecurityPolicyAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/policy/securitypolicy/securitypolicy-001", updateSecurityPolicyAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	updateSetup()
	expectedXML := "<securityPolicy><name>OVP_test_security_policy</name><description>this is a long description.</description><precedence>5501</precedence><actionsByCategory></actionsByCategory><securityGroupBinding><objectId>securitygroup-0001</objectId></securityGroupBinding><securityGroupBinding><objectId>securitygroup-0002</objectId></securityGroupBinding></securityPolicy>"

	xmlBytes, err := xml.Marshal(updateSecurityPolicyAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestGetResponse(t *testing.T) {
	updateSetup()
	updateSecurityPolicyAPI.SetResponseObject("string output returned")
	assert.Equal(t, "string output returned", updateSecurityPolicyAPI.GetResponse())
}
