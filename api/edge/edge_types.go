package edge

import "encoding/xml"

type Edge struct {
	XMLName  xml.Name `xml:"edge"`
	ObjectID string   `xml:"objectId,omitempty"`
	Version  string   `xml:"version,omitempty"`
	Type     string   `xml:"type,omitempty"`
	Name     string   `xml:"name"`
	FQDN     string   `xml:"fqdn"`
	Status   string   `xml:"status"`
	Tenant   string   `xml:"tenant"`
}

type PagedEdgeList struct {
	EdgePage EdgePage `xml:"edgePage,omitempty"`
}

type EdgePage struct {
	PagingInfo  PagingInfo    `xml:"pagingInfo,omitempty"`
	EdgeSummary []EdgeSummary `xml:"edgeSummary,omitempty"`
}

type PagingInfo struct {
	XMLName            xml.Name `xml:"pagingInfo"`
	PageSize           int      `xml:"pageSize,omitempty"`
	StartIndex         int      `xml:"startIndex"`
	TotalCount         int      `xml:"totalCount"`
	SortOrderAscending bool     `xml:"sortOrderAscending,omitempty"`
	SortBy             string   `xml:"sortBy,omitempty"`
}

type EdgeSummary struct {
	XMLName        xml.Name `xml:"edgeSummary"`
	ObjectID       string   `xml:"objectId,omitempty"`
	ObjectTypeName string   `xml:"objectTypeName,omitempty"`
	Revision       string   `xml:"revision,omitempty"`
	Type           string   `xml:"type,omitempty>typeName,omitempty"`
	Name           string   `xml:"name"`
}
