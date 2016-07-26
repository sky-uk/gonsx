package securitytag

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllSecurityTagsAPI - struct
type GetAllAttachedSecurityTagsAPI struct {
	*api.BaseAPI
}

// NewGetAll - Generates a new GetAllSecurityTagsAPI object.
func NewGetAllAttached(tagID string) *GetAllAttachedSecurityTagsAPI {
	this := new(GetAllAttachedSecurityTagsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/securitytags/tag/" + tagID + "/vm", nil, new(BasicInfoList))
	return this
}

// GetResponse returns the ResponseObject from CreateSecurityTagAPI
func (getAPI GetAllAttachedSecurityTagsAPI) GetResponse() *BasicInfoList {
	return getAPI.ResponseObject().(*BasicInfoList)
}
