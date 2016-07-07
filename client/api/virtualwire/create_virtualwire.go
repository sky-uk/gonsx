package virtualwire

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type CreateLogicalSwitchApi struct {
	*api.BaseApi
}

func NewCreate(name, desc, tenantID, scopeId string) *CreateLogicalSwitchApi {
	this := new(CreateLogicalSwitchApi)
	requestPayload := new(VirtualWireCreateSpec)
	requestPayload.Name = name
	requestPayload.TenantID = tenantID
	requestPayload.Description = desc
	// TODO: need to make it argument
	requestPayload.ControlPlaneMode = "UNICAST_MODE"

	this.BaseApi = api.NewBaseApi(http.MethodPost, "/api/2.0/vdn/scopes/" + scopeId +"/virtualwires", requestPayload, new(string))
	return this
}

func (this CreateLogicalSwitchApi) GetResponse() string {
	return this.ResponseObject().(string)
}
