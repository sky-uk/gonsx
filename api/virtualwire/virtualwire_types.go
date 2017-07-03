package virtualwire

import "encoding/xml"

// VirtualWires - top level xml element
type VirtualWires struct {
	DataPage DataPage `xml:"dataPage"`
}

// DataPage within VirtualWires
type DataPage struct {
	VirtualWires []VirtualWire `xml:"virtualWire"`
}

// VirtualWire is a single virtual wire object within virtualWire list.
type VirtualWire struct {
	XMLName          xml.Name `xml:"virtualWire"`
	Name             string   `xml:"name,omitempty"`
	ObjectID         string   `xml:"objectId,omitempty"`
	ControlPlaneMode string   `xml:"controlPlaneMode,omitempty"`
	Description      string   `xml:"description,omitempty"`
	TenantID         string   `xml:"tenantId,omitempty"`
}

// CreateSpec is used in create call on VirtualWire api.
type CreateSpec struct {
	XMLName          xml.Name `xml:"virtualWireCreateSpec"`
	Name             string   `xml:"name,omitempty"`
	ControlPlaneMode string   `xml:"controlPlaneMode,omitempty"`
	Description      string   `xml:"description,omitempty"`
	TenantID         string   `xml:"tenantId,omitempty"`
}
