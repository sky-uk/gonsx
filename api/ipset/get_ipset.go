package ipset

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetIpSetAPI base object.
type GetIpSetAPI struct {
	*api.BaseAPI
}

// NewGet returns a new object of GetIpSetAPI.
func NewGet(scopeID string) *GetIpSetAPI {
	this := new(GetIpSetAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/ipset/"+scopeID, nil, new(IpSet))
	return this
}

// GetResponse returns ResponseObject of GetIpSetAPI.
func (ga GetIpSetAPI) GetResponse() *IpSet {
	return ga.ResponseObject().(*IpSet)
}
