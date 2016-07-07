package lswitch

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type GetLogicalSwitchApi struct {
	*api.BaseApi
}

func NewGet(id string) *GetLogicalSwitchApi {
	this := new(GetLogicalSwitchApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/switches/" + id, nil, new(VdsContext))
	return this
}

func (this GetLogicalSwitchApi) GetResponse() *VdsContext {
	return this.ResponseObject().(*VdsContext)
}
