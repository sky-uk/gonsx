package securitypolicy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func constructSecurityPolicy(objectID, name string) *SecurityPolicy {
	securityGroupIDs := []string{"securitygroup-001", "securitygroup-002"}

	securityPolicy := &SecurityPolicy{}
	securityPolicy.Name = name
	securityPolicy.ObjectID = objectID
	securityPolicy.Precedence = "50001"
	securityPolicy.Description = "this is a long description."
	securityPolicy.SecurityGroupBinding = []SecurityGroup{}

	var securityGroupBindingList = []SecurityGroup{}
	for _, secGroupID := range securityGroupIDs {
		securityGroupBinding := SecurityGroup{ObjectID: secGroupID}
		securityGroupBindingList = append(securityGroupBindingList, securityGroupBinding)
	}
	securityPolicy.SecurityGroupBinding = securityGroupBindingList

	var secondarySecurityGroupList = []SecurityGroup{}
	secondarySecurityGroup := SecurityGroup{ObjectID: "securitygroup-197"}
	secondarySecurityGroupList = append(secondarySecurityGroupList, secondarySecurityGroup)

	// Next build the rule using the secondarySecurityGroup list.
	newRule := Action{
		Class:                  "firewallSecurityAction",
		Name:                   "DummyRule",
		Action:                 "allow",
		Category:               "firewall",
		Direction:              "outbound",
		VsmUUID:                "4221A849-079E-D13E-6B36-068D4F1222A9",
		SecondarySecurityGroup: secondarySecurityGroupList,
	}

	// Build actionsByCategory list.
	actionsByCategory := ActionsByCategory{Category: "firewall"}
	actionsByCategory.Actions = []Action{newRule}

	// Update the security policy with actionsByCategory
	securityPolicy.ActionsByCategory = actionsByCategory

	return securityPolicy
}

func TestAddSecurityGroupBinding(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")

	assert.Equal(t, "securitygroup-001", securityPolicy.SecurityGroupBinding[0].ObjectID)
	securityPolicy.AddSecurityGroupBinding("securitygroup-new")
	assert.Equal(t, "securitygroup-new", securityPolicy.SecurityGroupBinding[2].ObjectID)

	// test re-adding doesn't raise any errors.
	securityPolicy.AddSecurityGroupBinding("securitygroup-new")
	assert.Equal(t, "securitygroup-new", securityPolicy.SecurityGroupBinding[2].ObjectID)
}

func TestRemoveSecurityGroupBinding(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")

	assert.Equal(t, "securitygroup-001", securityPolicy.SecurityGroupBinding[0].ObjectID)
	securityPolicy.RemoveSecurityGroupBinding("securitygroup-001")
	assert.Equal(t, "securitygroup-002", securityPolicy.SecurityGroupBinding[0].ObjectID)

	// test removing again doesn't raise any errors.
	securityPolicy.RemoveSecurityGroupBinding("securitygroup-001")
	assert.Equal(t, "securitygroup-002", securityPolicy.SecurityGroupBinding[0].ObjectID)
}

func TestRemoveFirewallActionByName(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")

	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 1)
	securityPolicy.RemoveFirewallActionByName("DummyRule")
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 0)
}

func TestRemoveFirewallActionByUUID(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	ruleUUIDToRemove := "4221A849-079E-D13E-6B36-068D4F1222A9"

	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 1)
	securityPolicy.RemoveFirewallActionByUUID("wrong uuid")
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 1)
	securityPolicy.RemoveFirewallActionByUUID(ruleUUIDToRemove)
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 0)
}

func TestAddFirewallAction(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")

	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 1)
	addErr := securityPolicy.AddOutboundFirewallAction("new_action", "allow", "outbound", []string{"securitygroup-001"}, []string{"application-1"})
	assert.Nil(t, addErr)
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 2)

	securityPolicy.RemoveFirewallActionByName("DummyRule")
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 1)
	securityPolicy.RemoveFirewallActionByName("new_action")
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 0)

	// test failures for wrong action.
	wrongActionErr := securityPolicy.AddOutboundFirewallAction("new_action_2", "disallow_wrong", "inbound", []string{"securitygroup-001"}, []string{"application-1"})
	assert.NotNil(t, wrongActionErr)
	assert.Equal(t, "Action can be only 'allow' or 'block'", fmt.Sprint(wrongActionErr))

	// test failures for wrong action.
	wrongDirectionErr := securityPolicy.AddOutboundFirewallAction("new_action_2", "block", "inbound_wrong", []string{"securitygroup-001"}, []string{"application-1"})
	assert.NotNil(t, wrongDirectionErr)
	assert.Equal(t, "Direction can be only 'inbound' or 'outbound'", fmt.Sprint(wrongDirectionErr))

	// Now test adding new action on empty ActionsByCategory
	securityPolicy.AddOutboundFirewallAction("new_action_2", "block", "inbound", []string{"securitygroup-001"}, []string{"application-1"})
	assert.Len(t, securityPolicy.ActionsByCategory.Actions, 1)
	assert.True(t, securityPolicy.ActionsByCategory.Actions[0].IsEnabled)

}

func TestCheckFirewallRuleByUUID(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	uuidToCheck := "4221A849-079E-D13E-6B36-068D4F1222A9"

	assert.True(t, securityPolicy.CheckFirewallRuleByUUID(uuidToCheck))
	assert.False(t, securityPolicy.CheckFirewallRuleByUUID("not valid uuid"))

}

func TestGetFirewallRuleByUUID(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	uuidToCheck := "4221A849-079E-D13E-6B36-068D4F1222A9"

	actionObject := securityPolicy.GetFirewallRuleByUUID(uuidToCheck)
	assert.Equal(t, "DummyRule", actionObject.Name)
}

func TestGetFirewallRuleByName(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	expectedUUID := "4221A849-079E-D13E-6B36-068D4F1222A9"

	actionObject := securityPolicy.GetFirewallRuleByName("DummyRule")
	assert.Equal(t, expectedUUID, actionObject.VsmUUID)

	actionObject = securityPolicy.GetFirewallRuleByName("Wrong name")
	assert.Equal(t, "", actionObject.VsmUUID)
}

func TestFilterByName(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	securityPoliciesList := &SecurityPolicies{SecurityPolicies: []SecurityPolicy{*securityPolicy}}

	filteredSecurityPolicy := securityPoliciesList.FilterByName("OVP_test_security_policy")
	assert.Equal(t, "securitypolicy-0001", filteredSecurityPolicy.ObjectID)
	assert.Equal(t, "OVP_test_security_policy", filteredSecurityPolicy.Name)

}

func TestRemoveSecurityPolicyByName(t *testing.T) {
	firstSecurityPolicy := constructSecurityPolicy("securitypolicy-0001", "first_security_policy")
	secondSecurityPolicy := constructSecurityPolicy("securitypolicy-0002", "second_security_policy")

	securityPoliciesList := &SecurityPolicies{
		SecurityPolicies: []SecurityPolicy{
			*firstSecurityPolicy,
			*secondSecurityPolicy,
		},
	}

	assert.Equal(t, "first_security_policy", securityPoliciesList.SecurityPolicies[0].Name)
	assert.Len(t, securityPoliciesList.SecurityPolicies, 2)
	updatedPoliciesList := securityPoliciesList.RemoveSecurityPolicyByName("first_security_policy")
	assert.Equal(t, "second_security_policy", updatedPoliciesList.SecurityPolicies[0].Name)
	assert.Len(t, updatedPoliciesList.SecurityPolicies, 1)
}

func TestStringOutput(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	securityPoliciesList := &SecurityPolicies{SecurityPolicies: []SecurityPolicy{*securityPolicy}}

	stringOutputOfList := fmt.Sprint(securityPoliciesList)
	stringOutputOfOne := fmt.Sprint(securityPolicy)

	assert.Equal(t, "SecurityPolicies object, contains security policy objects.", stringOutputOfList)
	assert.Equal(t, "SecurityPolicy with objectId: securitypolicy-0001", stringOutputOfOne)
}

func TestMarshalToXML(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	convertedXML := securityPolicy.MarshalToXML()

	expectedXML := "<securityPolicy><objectId>securitypolicy-0001</objectId><name>OVP_test_security_policy</name><description>this is a long description.</description><precedence>50001</precedence><actionsByCategory><category>firewall</category><action class=\"firewallSecurityAction\"><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><name>DummyRule</name><action>allow</action><category>firewall</category><direction>outbound</direction><secondarySecurityGroup><objectId>securitygroup-197</objectId></secondarySecurityGroup></action></actionsByCategory><securityGroupBinding><objectId>securitygroup-001</objectId></securityGroupBinding><securityGroupBinding><objectId>securitygroup-002</objectId></securityGroupBinding></securityPolicy>"
	assert.Equal(t, expectedXML, convertedXML)

}

func TestMarshalToXMLAddOutBound(t *testing.T) {
	securityPolicy := constructSecurityPolicy("securitypolicy-0001", "OVP_test_security_policy")
	securityPolicy.AddOutboundFirewallAction("new_action", "allow", "outbound", []string{"securitygroup-001"}, []string{"application-1"})
	convertedXML := securityPolicy.MarshalToXML()

	expectedXML := "<securityPolicy><objectId>securitypolicy-0001</objectId><name>OVP_test_security_policy</name><description>this is a long description.</description><precedence>50001</precedence><actionsByCategory><category>firewall</category><action class=\"firewallSecurityAction\"><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><name>DummyRule</name><action>allow</action><category>firewall</category><direction>outbound</direction><secondarySecurityGroup><objectId>securitygroup-197</objectId></secondarySecurityGroup></action><action class=\"firewallSecurityAction\"><name>new_action</name><action>allow</action><category>firewall</category><direction>outbound</direction><isEnabled>true</isEnabled><secondarySecurityGroup><objectId>securitygroup-001</objectId></secondarySecurityGroup><applications><application><objectId>application-1</objectId></application></applications></action></actionsByCategory><securityGroupBinding><objectId>securitygroup-001</objectId></securityGroupBinding><securityGroupBinding><objectId>securitygroup-002</objectId></securityGroupBinding></securityPolicy>"
	assert.Equal(t, expectedXML, convertedXML)

}
