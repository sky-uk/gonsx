package edgeinterface

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type DeleteEdgeInterfaceApi struct {
	*api.BaseApi
}

func NewDelete(interfaceIndex, edgeId string) *DeleteEdgeInterfaceApi {
	this := new(DeleteEdgeInterfaceApi)
	this.BaseApi = api.NewBaseApi(http.MethodDelete, "/api/4.0/edges/"+ edgeId + "/interfaces/?index=" + interfaceIndex, nil, nil)
	return this
}