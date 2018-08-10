package ipset

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// GetIPSetAPI base object.
type GetIPSetAPI struct {
	*api.BaseAPI
}

// NewGet returns a new object of GetIpSetAPI.
func NewGet(scopeID string) *GetIPSetAPI {
	this := new(GetIPSetAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/ipset/"+scopeID, nil, new(IPSet))
	return this
}

// GetResponse returns ResponseObject of GetIpSetAPI.
func (ga GetIPSetAPI) GetResponse() *IPSet {
	return ga.ResponseObject().(*IPSet)
}
