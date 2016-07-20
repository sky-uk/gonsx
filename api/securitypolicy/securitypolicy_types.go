package securitypolicy

import "encoding/xml"

// SecurityPolicies top level struct
type SecurityPolicies struct {
	SecurityPolicies []SecurityPolicy `xml:"securityPolicy"`
}

// SecurityPolicy object struct
type SecurityPolicy struct {
	XMLName            xml.Name          `xml:"securityPolicy"`
	ObjectID           string            `xml:"objectId,omitempty"`
	ObjectTypeName     string            `xml:"objectTypeName,omitempty"`
	VsmUUID            string            `xml:"vsmUuid,omitempty"`
	NodeID             string            `xml:"nodeId,omitempty"`
	Revision           int               `xml:"revision,omitempty"`
	TypeName           string            `xml:"type>typeName"`
	Name               string            `xml:"name,omitempty"`
	Description        string            `xml:"description"`
	IsUniversal        bool              `xml:"isUniversal"`
	InheritanceAllowed bool              `xml:"inheritanceAllowed"`
	ActionsByCategory  ActionsByCategory `xml:"actionsByCategory"`
}

// ActionsByCategory element of SecurityPolicy.
type ActionsByCategory struct {
	XMLName  xml.Name `xml:"actionsByCategory"`
	Category string   `xml:"category"`
	Actions  []Action `xml:"action"`
}

// Action element of ActionsByCategory list.
type Action struct {
	XMLName        xml.Name `xml:"action"`
	ObjectID       string   `xml:"objectId,omitempty"`
	ObjectTypeName string   `xml:"objectTypeName,omitempty"`
	VsmUUID        string   `xml:"vsmUuid,omitempty"`
	NodeID         string   `xml:"nodeId,omitempty"`
	Revision       int      `xml:"revision,omitempty"`
	TypeName       string   `xml:"type>typeName"`
	IsEnabled      bool     `xml:"isEnabled"`
}
