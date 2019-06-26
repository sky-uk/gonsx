package firewall

import (
	"encoding/xml"
)

const (
	ALL_EDGES                   ElemType = "ALL_EDGES"
	Application                 ElemType = "Application"
	ApplicationGroup            ElemType = "ApplicationGroup"
	Datacenter                  ElemType = "Datacenter"
	DISTRIBUTED_FIREWALL        ElemType = "DISTRIBUTED_FIREWALL"
	DistributedVirtualPortgroup ElemType = "DistributedVirtualPortgroup"
	Edge                        ElemType = "Edge"
	GlobalRoot                  ElemType = "GlobalRoot"
	HostSystem                  ElemType = "HostSystem"
	IPSet                       ElemType = "IPSet"
	Ipv4Address                 ElemType = "Ipv4Address"
	Ipv6Address                 ElemType = "Ipv6Address"
	VirtualWire                 ElemType = "VirtualWire"
	MACSet                      ElemType = "MACSet"
	Network                     ElemType = "Network"
	ALL_PROFILE_BINDINGS        ElemType = "ALL_PROFILE_BINDINGS"
	ResourcePool                ElemType = "ResourcePool"
	SecurityGroup               ElemType = "SecurityGroup"
	Vnic                        ElemType = "Vnic"
)

const (
	In    Direction = "in"
	Out   Direction = "out"
	InOut Direction = "inout"
)

const (
	Allow  Action = "allow"
	Block  Action = "block"
	Reject Action = "reject"
)

type ElemType string

type Direction string

type Action string

type AppliedToList struct {
	Elements []Element `xml:"appliedTo"`
}

type Sources struct {
	Excluded bool      `xml:"excluded,attr"`
	Elements []Element `xml:"source,omitempty"`
}

type Destinations struct {
	Excluded bool      `xml:"excluded,attr"`
	Elements []Element `xml:"destination,omitempty"`
}

type Services struct {
	Elements []Element `xml:"service,omitempty"`
}

type Element struct {
	Name    string   `xml:"name,omitempty"`
	Value   string   `xml:"value"`
	Type    ElemType `xml:"type,omitempty"`
	IsValid bool     `xml:"isValid,omitempty"`
}

type Rule struct {
	XMLName       xml.Name       `xml:"rule"`
	ID            int            `xml:"id,attr,omitempty"`
	Disabled      bool           `xml:"disabled,attr"`
	Logged        bool           `xml:"logged,attr"`
	Name          string         `xml:"name"`
	Notes         string         `xml:"notes"`
	Action        Action         `xml:"action"`
	AppliedToList *AppliedToList `xml:"appliedToList,omitempty"`
	SectionId     int            `xml:"sectionId,omitempty"`
	Sources       *Sources       `xml:"sources,omitempty"`
	Destinations  *Destinations  `xml:"destinations,omitempty"`
	Services      *Services      `xml:"services,omitempty"`
	Direction     Direction      `xml:"direction,omitempty"`
	Precedence    string         `xml:"precedence,omitempty"`
	PacketType    string         `xml:"packetType,omitempty"`
}

type Section struct {
	GenerationNumber string `xml:"generationNumber,attr,omitempty"`
	ID               string `xml:"id,attr,omitempty"`
	Name             string `xml:"name,attr"`
	Stateless        string `xml:"stateless,attr,omitempty"`
	TcpStrict        string `xml:"tcpStrict,attr,omitempty"`
	Timestamp        string `xml:"timestamp,attr,omitempty"`
	Type             string `xml:"type,attr,omitempty"`
	UseSid           string `xml:"useSid,attr,omitempty"`
	Rule             []Rule `xml:"rule"`
}

type FirewallConfiguration struct {
	XMLName        xml.Name `xml:"firewallConfiguration"`
	Timestamp      string   `xml:"timestamp,attr,omitempty"`
	ContextId      string   `xml:"contextId,omitempty"`
	Layer3Sections *Section `xml:"layer3Sections"`
	Layer2Sections *Section `xml:"layer2Sections"`
}
