package tzone

type NetworkScopeList struct {
	NetworkScopeList	[]NetworkScope	`xml:"vdnScope"`
}

type NetworkScope struct {
	ObjectId	string	`xml:"objectId"`
	Name		string	`xml:"name"`
}
