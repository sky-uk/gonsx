package distributedfw

type FirewallConfiguration struct {
	Layer3Sections Layer3Sections `xml:"layer3Sections"`
	Layer2Sections Layer2Sections `xml:"layer2Sections"`
}

type Layer2Sections struct {
	Sections []Section `xml:"section"`
}

type Layer3Sections struct {
	Sections []Section `xml:"section"`
}

type Section struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Rules []Rule `xml:"rule"`
}

type Rule struct {
	RuleID        string        `xml:"id,attr"`
	Name          string        `xml:"name"`
	Disabled      bool          `xml:"disabled,attr"`
	RuleType      string        `xml:"ruleType"`
	Logged        string        `xml:"logged,attr"`
	Source        string        `xml:"source"`
	Destination   string        `xml:"destination"`
	Action        string        `xml:"action"`
	EdgeID        string        `xml:"edgeId"`
	AppliedToList []AppliedTo   `xml:"appliedToList>appliedTo"`
	Sources       []Source      `xml:"sources>source"`
	Destinations  []Destination `xml:"destinations>destination"`
	Services      []Service     `xml:"services>service"`
	SectionID     uint8         `xml:"sectionID"`
	Direction     string        `xml:"direction"`
	PacketType    string        `xml:"packetType"`
}

type Service struct {
	Name            string `xml:"name"`
	Value           string `xml:"value"`
	DestinationPort uint16 `xml:"destinationPort"`
	Protocol        uint8  `xml:"protocol"`
	SubProtocol     uint8  `xml:"subProtocol"`
}

type AppliedTo struct {
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Type    string `xml:"type"`
	IsValid bool   `xml:"isValid"`
}

type Source struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Value   string `xml:"value"`
	IsValid bool   `xml:"isValid"`
}

type Destination struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Value   string `xml:"value"`
	IsValid bool   `xml:"isValid"`
}
