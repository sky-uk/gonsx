package edge

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetEdgeAPI base object.
type GetEdgeAPI struct {
	*api.BaseAPI
}

// NewGet returns a new object of GetEdgeAPI.
func NewGet(edgeID string) *GetEdgeAPI {
	url := fmt.Sprintf("/api/4.0/edges/%s", edgeID)
	fmt.Println("URL:", url)
	this := new(GetEdgeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new(Edge))
	return this
}

// GetResponse returns ResponseObject of PagedEdgeList.
func (ga GetEdgeAPI) GetResponse() Edge {
	return *ga.ResponseObject().(*Edge)
}
