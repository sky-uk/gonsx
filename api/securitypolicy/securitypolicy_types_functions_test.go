package securitypolicy

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func constructSecurityPolicy() *SecurityPolicy {
	securityGroupIDs := []string{"securitygroup-001", "securitygroup-002"}

	securityPolicy := &SecurityPolicy{}
	securityPolicy.Name = "OVP_test_security_policy"
	securityPolicy.ObjectID = "securitypolicy-0001"
	securityPolicy.Precedence = "50001"
	securityPolicy.Description = "this is a long description."
	securityPolicy.SecurityGroupBinding = []SecurityGroup{}

	var securityGroupBindingList = []SecurityGroup{}
	for _, secGroupID := range securityGroupIDs {
		securityGroupBinding := SecurityGroup{ObjectID: secGroupID}
		securityGroupBindingList = append(securityGroupBindingList, securityGroupBinding)
	}
	securityPolicy.SecurityGroupBinding = securityGroupBindingList
	securityPolicy.ActionsByCategory = ActionsByCategory{Actions: []Action{}}

	return securityPolicy
}

func TestAddSecurityGroupBinding(t *testing.T) {
	securityPolicy := constructSecurityPolicy()

	assert.Equal(t, "securitygroup-001", securityPolicy.SecurityGroupBinding[0].ObjectID)
	securityPolicy.AddSecurityGroupBinding("securitygroup-new")
	assert.Equal(t, "securitygroup-new", securityPolicy.SecurityGroupBinding[2].ObjectID)

	// test re-adding doesn't raise any errors.
	securityPolicy.AddSecurityGroupBinding("securitygroup-new")
	assert.Equal(t, "securitygroup-new", securityPolicy.SecurityGroupBinding[2].ObjectID)
}


func TestRemoveSecurityGroupBinding(t *testing.T) {
	securityPolicy := constructSecurityPolicy()

	assert.Equal(t, "securitygroup-001", securityPolicy.SecurityGroupBinding[0].ObjectID)
	securityPolicy.RemoveSecurityGroupBinding("securitygroup-001")
	assert.Equal(t, "securitygroup-002", securityPolicy.SecurityGroupBinding[0].ObjectID)

	// test removing again doesn't raise any errors.
	securityPolicy.RemoveSecurityGroupBinding("securitygroup-001")
	assert.Equal(t, "securitygroup-002", securityPolicy.SecurityGroupBinding[0].ObjectID)
}

func TestFilterByName(t *testing.T) {
	securityPolicy := constructSecurityPolicy()
	securityPoliciesList := &SecurityPolicies{SecurityPolicies: []SecurityPolicy{*securityPolicy}}

	filteredSecurityPolicy := securityPoliciesList.FilterByName("OVP_test_security_policy")
	assert.Equal(t, "securitypolicy-0001", filteredSecurityPolicy.ObjectID)
	assert.Equal(t, "OVP_test_security_policy", filteredSecurityPolicy.Name)

}

func TestStringOutput(t *testing.T) {
	securityPolicy := constructSecurityPolicy()
	securityPoliciesList := &SecurityPolicies{SecurityPolicies: []SecurityPolicy{*securityPolicy}}

	stringOutputOfList := fmt.Sprint(securityPoliciesList)
	stringOutputOfOne := fmt.Sprint(securityPolicy)

	assert.Equal(t, "SecurityPolicies object, contains security policy objects.", stringOutputOfList)
	assert.Equal(t, "SecurityPolicy with objectId: securitypolicy-0001", stringOutputOfOne)
}

func TestMarshalToXML(t *testing.T) {
	securityPolicy := constructSecurityPolicy()
	convertedXML :=  securityPolicy.MarshalToXML()

	expectedXML := "<securityPolicy><objectId>securitypolicy-0001</objectId><name>OVP_test_security_policy</name><description>this is a long description.</description><precedence>50001</precedence><actionsByCategory></actionsByCategory><securityGroupBinding><objectId>securitygroup-001</objectId></securityGroupBinding><securityGroupBinding><objectId>securitygroup-002</objectId></securityGroupBinding></securityPolicy>"
	assert.Equal(t, expectedXML, convertedXML)

}
