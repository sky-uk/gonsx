package ipset

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateIpSetAPI api object
type CreateIpSetAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateIpSetAPI.
func NewCreate(scopeID string, ipSet *IpSet) *CreateIpSetAPI {
	this := new(CreateIpSetAPI)
	requestPayload := ipSet

	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/api/2.0/services/ipset/"+scopeID, requestPayload, new(string))
	return this
}

// GetResponse returns a ResponseObject of CreateIpSetAPI.
func (ca CreateIpSetAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
