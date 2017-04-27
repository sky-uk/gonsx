package securitytag

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllSecurityTagsAttachedToVmAPI - struct
type GetAllSecurityTagsAttachedToVmAPI struct {
	*api.BaseAPI
}

// NewGetAll - Generates a new GetAllSecurityTagsAttachedToVmAPI object.
func NewGetAllAttachedToVM(vmID string) *GetAllSecurityTagsAttachedToVmAPI {
	this := new(GetAllSecurityTagsAttachedToVmAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/securitytags/vm/"+vmID, nil, new(SecurityTags))
	return this
}

// GetResponse returns the ResponseObject from GetAllSecurityTagsAttachedToVmAPI
func (getAPI GetAllSecurityTagsAttachedToVmAPI) GetResponse() *SecurityTags {
	return getAPI.ResponseObject().(*SecurityTags)
}

