package dhcppool

// DHCPPool struct
type DHCPPool struct {
	IPRange             string `xml:"ipRange"`
	DefaultGateway      string `xml:"defaultGateway"`
	SubnetMask          string `xml:"subnetMask"`
	DomainName          string `xml:"domainName"`
	PrimaryNameServer   string `xml:"primaryNameServer"`
	SecondaryNameServer string `xml:"secondaryNameServer"`
	LeaseTime           uint8  `xml:"leaseTime"`
	AutoConfigDNS       bool   `xml:"autoConfigureDNS"`
	NextServer          string `xml:"nextServer"`
}

// DHCPOptions struct
type DHCPOptions struct {
	Option121 option121 `xml:"option121"`
	Option66  string    `xml:"option66"`
	Option67  string    `xml:"option67"`
}

// option121 struct
type option121 struct {
	StaticRoute staticRoute `xml:"option121>staticRoute"`
}

// staticRoute struct
type staticRoute struct {
	DestinationSubnet string `xml:"destinationSubnet"`
	Router            string `xml:"router"`
}
