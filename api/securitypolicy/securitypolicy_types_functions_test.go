package securitypolicy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func constructSecurityPolicy() *SecurityPolicy{
	securityGroupIDs := []string{"securitygroup-001", "securitygroup-002"}

	securityPolicy := &SecurityPolicy{}
	securityPolicy.Name = "OVP_test_security_policy"
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
}

func TestRemoveSecurityGroupBinding(t *testing.T) {
	securityPolicy := constructSecurityPolicy()

	assert.Equal(t, "securitygroup-001", securityPolicy.SecurityGroupBinding[0].ObjectID)
	securityPolicy.RemoveSecurityGroupBinding("securitygroup-001")
	assert.Equal(t, "securitygroup-002", securityPolicy.SecurityGroupBinding[0].ObjectID)
}

