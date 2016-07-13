package service

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

type CreateServiceApi struct {
	*api.BaseApi
}

func NewCreate(scopeId, name, desc, proto, ports string) *CreateServiceApi {
	this := new(CreateServiceApi)
	requestPayload := new(ApplicationService)
	requestPayload.Name = name
	requestPayload.Description = desc

	element := ServiceElement{ApplicationProtocol: proto, Value: ports}
	requestPayload.Element = []ServiceElement{element}

	this.BaseApi = api.NewBaseApi(http.MethodPost, "/api/2.0/services/application/"+scopeId, requestPayload, new(ApplicationService))
	return this
}

func (this CreateServiceApi) GetResponse() *ApplicationService {
	return this.ResponseObject().(*ApplicationService)
}
