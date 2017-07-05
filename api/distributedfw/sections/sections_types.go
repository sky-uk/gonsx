package sections

import "github.com/sky-uk/gonsx/api/distributedfw/fwrules"

// Section - Contains the rules
type Section struct {
	ID    string         `xml:"id,attr"`
	Name  string         `xml:"name,attr"`
	Type  string         `xml:"type,attr"`
	Rules []fwrules.Rule `xml:"rule"`
}
