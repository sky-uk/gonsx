package lswitch

import (
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
	"net/http"
)

type CreateLogicalSwitchApi struct {
	*api.BaseApi
}

func NewCreate() *CreateLogicalSwitchApi {
	this := new(CreateLogicalSwitchApi)
	requestPayload := new(VdsContext)
	requestPayload.Switch.ObjectId = "dvs-61"
	requestPayload.Switch.Name = "dvs-luis-test"
	requestPayload.Switch.Type.TypeName = "VmwareDistributedVirtualSwitch"
	requestPayload.Switch.ObjectTypeName = "VmwareDistributedVirtualSwitch"
	requestPayload.Switch.Revision = 0
	requestPayload.Teaming = "LOADBALANCE_SRCID"
	requestPayload.MTU = 1600

	this.BaseApi = api.NewBaseApi(http.MethodPost, "/api/2.0/vdn/switches", requestPayload, new(VdsContext))
	return this
}

func (this CreateLogicalSwitchApi) GetResponse() *VdsContext {
	return this.ResponseObject().(*VdsContext)
}
