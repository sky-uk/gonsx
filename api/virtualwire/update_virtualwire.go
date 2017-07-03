package virtualwire

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// UpdateVirtualWireAPI base api object.
type UpdateVirtualWireAPI struct {
	*api.BaseAPI
}

// NewUpdate returns a new object of UpdateVirtualWireAPI.
func NewUpdate(virtualWire VirtualWire) *UpdateVirtualWireAPI {
	this := new(UpdateVirtualWireAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/2.0/vdn/virtualwires/"+virtualWire.ObjectID, virtualWire, nil)
	return this
}

// GetUpdateResponse returns ResponseObject of UpdateVirtualWireAPI.
func (updateVirtualWire UpdateVirtualWireAPI) GetUpdateResponse() string {
	return updateVirtualWire.ResponseObject().(string)
}
