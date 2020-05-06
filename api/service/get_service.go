package service

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetServiceAPI base object.
type GetServiceAPI struct {
	*api.BaseAPI
}

// NewGet returns a new object of GetServiceAPI.
func NewGet(serviceID string) *GetServiceAPI {
	this := new(GetServiceAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/application/"+serviceID, nil, new(ApplicationService))
	return this
}

// GetResponse returns ResponseObject of GetServiceAPI.
func (ga GetServiceAPI) GetResponse() *ApplicationService {
	return ga.ResponseObject().(*ApplicationService)
}
