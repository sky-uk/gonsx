package virtualwire

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type GetVirtualWireApi struct {
	*api.BaseApi
}

func NewGet(scopeId string) *GetVirtualWireApi {
	this := new(GetVirtualWireApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/scopes/"+ scopeId + "/virtualwires", nil, new(VirtualWires))
	return this
}

func (this GetVirtualWireApi) GetResponse() *VirtualWires {
	return this.ResponseObject().(*VirtualWires)
}
