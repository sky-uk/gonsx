package ipset

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
    "fmt"
)

// UpdateIpSetAPI object
type UpdateIpSetAPI struct {
	*api.BaseAPI
}

// NewUpdate creates a new object of UpdateIpSetAPI
func NewUpdate(ipsetID string, ipsetPayload *IpSet) *UpdateIpSetAPI {
	this := new(UpdateIpSetAPI)
	endpointURL := "/api/2.0/services/ipset/" + ipsetID
    fmt.Println(endpointURL)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, endpointURL, ipsetPayload, new(IpSet))
	return this
}

// GetResponse returns the ResponseObject from UpdateServiceAPI
func (updateAPI UpdateIpSetAPI) GetResponse() *IpSet {
	return updateAPI.ResponseObject().(*IpSet)
}
