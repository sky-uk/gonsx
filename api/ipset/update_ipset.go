package ipset

import (
	"fmt"
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// UpdateIPSetAPI object
type UpdateIPSetAPI struct {
	*api.BaseAPI
}

// NewUpdate creates a new object of UpdateIpSetAPI
func NewUpdate(ipsetID string, ipsetPayload *IPSet) *UpdateIPSetAPI {
	this := new(UpdateIPSetAPI)
	endpointURL := "/api/2.0/services/ipset/" + ipsetID
	fmt.Println(endpointURL)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, endpointURL, ipsetPayload, new(IPSet))
	return this
}

// GetResponse returns the ResponseObject from UpdateServiceAPI
func (updateAPI UpdateIPSetAPI) GetResponse() *IPSet {
	return updateAPI.ResponseObject().(*IPSet)
}
