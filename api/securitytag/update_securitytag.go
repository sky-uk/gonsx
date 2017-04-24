package securitytag

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

//UpdateSecurityTagAPI - struct
type UpdateSecurityTagAPI struct {
	*api.BaseAPI
}

// NewUpdate - Generates a new UpdateSecurityTagAPI object.
func NewUpdate(name, desc string) *UpdateSecurityTagAPI {
	this := new(UpdateSecurityTagAPI)
	requestPayload := new(SecurityTag)
	requestPayload.Name = name
	requestPayload.Description = desc
	// TODO: need to make it argument
	requestPayload.TypeName = "SecurityTag"
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/2.0/services/securitytags/tag", requestPayload, new(string))
	return this
}

// GetResponse returns the ResponseObject from UpdateSecurityTagAPI
func (updateAPI UpdateSecurityTagAPI) GetResponse() string {
	return updateAPI.ResponseObject().(string)
}
