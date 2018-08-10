package ipset

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// DeleteIPSetAPI base object.
type DeleteIPSetAPI struct {
	*api.BaseAPI
}

// NewDelete returns a new object of DeleteIpSetAPI.
func NewDelete(ipSetID string) *DeleteIPSetAPI {
	this := new(DeleteIPSetAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/2.0/services/ipset/"+ipSetID, nil, nil)
	return this
}
