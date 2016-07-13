package service

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

type GetAllServiceApi struct {
	*api.BaseApi
}

func NewGetAll(scopeId string) *GetAllServiceApi {
	this := new(GetAllServiceApi)
	this.BaseApi = api.NewBaseApi(http.MethodGet, "/api/2.0/services/application/"+scopeId, nil, new(ApplicationsList))
	return this
}

func (this GetAllServiceApi) GetResponse() *ApplicationsList {
	return this.ResponseObject().(*ApplicationsList)
}
