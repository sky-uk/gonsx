package virtualwire

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type GetVirtualWireApi struct {
	*api.BaseApi
}

func NewGet(id string) *GetVirtualWireApi {
	this := new(GetVirtualWireApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/virtualwires/" + id, nil, new(VirtualWire))
	return this
}

func (this GetVirtualWireApi) GetResponse() *VirtualWire {
	return this.ResponseObject().(*VirtualWire)
}
