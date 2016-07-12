package tzone

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

type GetAllTransportZonesApi struct {
	*api.BaseApi
}

func NewGetAll() *GetAllTransportZonesApi {
	this := new(GetAllTransportZonesApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/scopes", nil, new(NetworkScopeList))
	return this
}

func (this GetAllTransportZonesApi) GetResponse() *NetworkScopeList {
	return this.ResponseObject().(*NetworkScopeList)
}
