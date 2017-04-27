package securitytag

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// UpdateAttachedSecurityTagAPI - struct
type UpdateAttachedSecurityTagsAPI struct {
	*api.BaseAPI
}

// NewDelete - Generates a new UpdateAttachedSecurityTagAPI object.
func NewUpdateAttachedTags(vmID string,securityTagPayload *SecurityTagAttachmentList) *UpdateAttachedSecurityTagsAPI {

	this := new(UpdateAttachedSecurityTagsAPI)


	//endpointUrl := "GET /2.0/services/securitytags/vm/"+vmID
	//currentSecurityTags := api.NewBaseAPI(http.MethodGet, endpointUrl, nil, nil)
	//TODO
	//Compare current attached tags with ones in payload
	//Remove any attached tags which are not in payload


	endpointURL := "/api/2.0/services/securitytags/vm/" + vmID + "?action=ASSIGN_TAGS"
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, endpointURL,securityTagPayload,nil)

	return this
}

// GetResponse returns the ResponseObject from CreateSecurityTagAPI
func (updateAPI UpdateAttachedSecurityTagsAPI) GetResponse() string {
	return updateAPI.ResponseObject().(string)
}