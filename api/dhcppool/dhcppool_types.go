package dhcppool


type DHCPPool struct {

	IPRange string `xml:"ipRange"`
	DefaultGateway string `xml:"defaultGateway"`
	SubnetMask string `xml:subnetMask`
	DomainName string `xml:"domainName"`
	PrimaryNameServer string `xml:"primaryNameServer"`
	SecondaryNameServer string `xml:"secondaryNameServer"`
	LeaseTime uint8 `xml:"leaseTime"`
	AutoConfigDNS bool `xml:"autoConfigureDNS"`
	NextServer string `xml:"nextServer"`

}

type DHCPOptions struct {
	Option121 option121 `xml:"option121"`
	Option66 string `xml:"option66"`
	option67 string `xml:"option67"`

}


type option121 struct {
	StaticRoute staticRoute `xml:"option121>staticRoute"`
}

type staticRoute struct {
	DestinationSubnet string `xml:"destinationSubnet"`
	Router string `xml:"router"`
}
