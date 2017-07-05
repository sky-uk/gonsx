package ipset

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// DeleteIpSetAPI base object.
type DeleteIpSetAPI struct {
	*api.BaseAPI
}

// NewDelete returns a new object of DeleteIpSetAPI.
func NewDelete(ipSetID string) *DeleteIpSetAPI {
	this := new(DeleteIpSetAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/2.0/services/ipset/"+ipSetID, nil, nil)
	return this
}
