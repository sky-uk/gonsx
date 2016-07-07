package lswitch

import "encoding/xml"

type VdsContextList struct {
	VdsContextList []VdsContext	`xml:"vdsContext"`
}

type VdsContext struct {
	XMLName 	xml.Name	`xml:"vdsContext"`
	Switch		Switch		`xml:"switch"`
	Teaming    	string		`xml:"teaming"`
	MTU    		uint16		`xml:"mtu"`
}

type Switch struct {
	ObjectId	string		`xml:"objectId"`
	Name		string		`xml:"name"`
	Type    	Type		`xml:"type"`
	Revision    	uint16		`xml:"revision"`
	ObjectTypeName	string		`xml:"objectTypeName"`
}

type Type struct {
	TypeName	string		`xml:"typeName"`
}