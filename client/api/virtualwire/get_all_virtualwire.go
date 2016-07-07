package virtualwire

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type GetAllVirtualWiresApi struct {
	*api.BaseApi
}

func NewGetAll(scopeId string) *GetAllVirtualWiresApi {
	this := new(GetAllVirtualWiresApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/scopes/" + scopeId + "/virtualwires", nil, new(VirtualWires))
	return this
}

func (this GetAllVirtualWiresApi) GetResponse() *VirtualWires {
	return this.ResponseObject().(*VirtualWires)
}