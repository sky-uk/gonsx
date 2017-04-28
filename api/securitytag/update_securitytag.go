package securitytag

import (
	"fmt"
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

//UpdateSecurityTagAPI - struct
type UpdateSecurityTagAPI struct {
	*api.BaseAPI
}

//NewUpdate - Generates a new UpdateSecurityTagAPI object.
func NewUpdate(securityTagID string, securityTagPayload *SecurityTag) *UpdateSecurityTagAPI {
	this := new(UpdateSecurityTagAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/2.0/services/securitytags/tag/"+securityTagID, securityTagPayload, new(SecurityTag))
	return this
}

// GetResponse returns the ResponseObject from UpdateSecurityTagAPI
func (updateAPI UpdateSecurityTagAPI) GetResponse() string {
	return updateAPI.ResponseObject().(string)
}
