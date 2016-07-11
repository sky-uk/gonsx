package edgeinterface


import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type GetAllEdgeInterfacesApi struct {
	*api.BaseApi
}

func NewGetAll(edgeId string) *GetAllEdgeInterfacesApi {
	this := new(GetAllEdgeInterfacesApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/4.0/edges/" + edgeId + "/interfaces", nil, new(EdgeInterfaces))
	return this
}

func (this GetAllEdgeInterfacesApi) GetResponse() *EdgeInterfaces {
	return this.ResponseObject().(*EdgeInterfaces)
}
