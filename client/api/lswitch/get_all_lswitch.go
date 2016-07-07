package lswitch

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type GetAllLogicalSwitchesApi struct {
	*api.BaseApi
}

func NewGetAll() *GetAllLogicalSwitchesApi {
	this := new(GetAllLogicalSwitchesApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/vdn/switches", nil, new(VdsContextList))
	return this
}

func (this GetAllLogicalSwitchesApi) GetResponse() *VdsContextList {
	return this.ResponseObject().(*VdsContextList)
}