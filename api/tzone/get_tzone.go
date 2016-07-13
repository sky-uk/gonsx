package tzone

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

type GetTransportZoneApi struct {
	*api.BaseApi
}

func NewGet(id string) *GetTransportZoneApi {
	this := new(GetTransportZoneApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/scopes/"+id, nil, new(NetworkScope))
	return this
}

func (this GetTransportZoneApi) GetResponse() *NetworkScope {
	return this.ResponseObject().(*NetworkScope)
}
