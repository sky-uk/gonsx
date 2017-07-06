package edge

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllEdgesAPI base object.
type GetAllEdgesAPI struct {
	*api.BaseAPI
}

// NewGetAll returns a new object of GetAllEdgesAPI.
func NewGetAll(pageSize, startIndex int) *GetAllEdgesAPI {
	url := fmt.Sprintf("/api/4.0/edges?pagesize=%d&startindex=%d", pageSize, startIndex)
	fmt.Println("URL:", url)
	this := new(GetAllEdgesAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new(PagedEdgeList))
	return this
}

// GetResponse returns ResponseObject of PagedEdgeList.
func (ga GetAllEdgesAPI) GetResponse() PagedEdgeList {
	return *ga.ResponseObject().(*PagedEdgeList)
}
