package distributedfirewall

type Rule struct {
	RuleID        uint8       `xml:"ruleId"`
	Name          string      `xml:"name"`
	Disabled      bool        `xml:"disabled"`
	RuleType      string      `xml:"ruleType"`
	Source        string      `xml:"source"`
	Destination   string      `xml:"destination"`
	Action        string      `xml:"action"`
	EdgeID        string      `xml:"edgeId"`
	AppliedToList []AppliedTo `xml:"appliedToList"`
	Sources       []Source    `xml:"sources"`
	Services      []Service   `xml:"services"`
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
