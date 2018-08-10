package ipset

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// GetAllIPSetAPI base object.
type GetAllIPSetAPI struct {
	*api.BaseAPI
}

// NewGetAll returns a new object of GetAllIpSetAPI.
func NewGetAll(scopeID string) *GetAllIPSetAPI {
	this := new(GetAllIPSetAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/ipset/scope/"+scopeID, nil, new(List))
	return this
}

// GetResponse returns ResponseObject of GetAllIpSetAPI.
func (ga GetAllIPSetAPI) GetResponse() *List {
	return ga.ResponseObject().(*List)
}
