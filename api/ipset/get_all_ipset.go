package ipset

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllIpSetAPI base object.
type GetAllIpSetAPI struct {
	*api.BaseAPI
}

// NewGetAll returns a new object of GetAllIpSetAPI.
func NewGetAll(scopeID string) *GetAllIpSetAPI {
	this := new(GetAllIpSetAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/ipset/scope/"+scopeID, nil, new(List))
	return this
}

// GetResponse returns ResponseObject of GetAllIpSetAPI.
func (ga GetAllIpSetAPI) GetResponse() *List {
	return ga.ResponseObject().(*List)
}
