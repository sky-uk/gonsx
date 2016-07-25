package securitypolicy

import (
	"encoding/xml"
	"fmt"
)

func (sp SecurityPolicy) String() string {
	return fmt.Sprintf("SecurityPolicy with objectId: %s", sp.ObjectID)
}

// MarshalToXML converts the object into XML
func (sp SecurityPolicy) MarshalToXML() string {
	xmlBytes, _ := xml.Marshal(sp)
	return string(xmlBytes)
}

// AddSecurityGroupBinding - Adds security group to list of SecurityGroupBinding if it doesn't exists.
func (sp *SecurityPolicy) AddSecurityGroupBinding(objectID string) {
	for _, secGroup := range sp.SecurityGroupBinding {
		if secGroup.ObjectID == objectID {
			return
		}
	}
	// if we reached here that means we couldn't find one, and let's add the sec group.
	sp.SecurityGroupBinding = append(sp.SecurityGroupBinding, SecurityGroup{ObjectID: objectID})
	return
}

// RemoveSecurityGroupBinding - Adds security group to list of SecurityGroupBinding if it doesn't exists.
func (sp *SecurityPolicy) RemoveSecurityGroupBinding(objectID string) {
	for idx, secGroup := range sp.SecurityGroupBinding {
		if secGroup.ObjectID == objectID {
			sp.SecurityGroupBinding = append(sp.SecurityGroupBinding[:idx], sp.SecurityGroupBinding[idx+1:]...)
			return
		}
	}
	return
}

func (spList SecurityPolicies) String() string {
	return fmt.Sprint("SecurityPolicies object, contains security policy objects.")
}

// FilterByName returns a single security policy object if it matches the name in SecurityPolicies list.
func (spList SecurityPolicies) FilterByName(name string) *SecurityPolicy {
	var securityPolicyFound SecurityPolicy
	for _, securityPolicy := range spList.SecurityPolicies {
		if securityPolicy.Name == name {
			securityPolicyFound = securityPolicy
			break
		}
	}
	return &securityPolicyFound
}
