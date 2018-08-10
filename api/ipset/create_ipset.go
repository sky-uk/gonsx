package ipset

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// CreateIPSetAPI api object
type CreateIPSetAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateIpSetAPI.
func NewCreate(scopeID string, ipSet *IPSet) *CreateIPSetAPI {
	this := new(CreateIPSetAPI)
	requestPayload := ipSet

	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/api/2.0/services/ipset/"+scopeID, requestPayload, new(string))
	return this
}

// GetResponse returns a ResponseObject of CreateIpSetAPI.
func (ca CreateIPSetAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
