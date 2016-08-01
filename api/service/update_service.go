package service

import (
        "github.com/sky-uk/gonsx/api"
        "net/http"
)

// UpdateServiceAPI ...
type UpdateServiceAPI struct {
        *api.BaseAPI
}

// NewUpdate creates a new object of UpdateServiceAPI
func NewUpdate(scopeID, name, desc, proto, ports string) *UpdateServiceAPI {
        this := new(UpdateServiceAPI)
        securityPolicyPayload := new(ApplicationService)
        securityPolicyPayload.Name = name
	securityPolicyPayload.Description = desc

	element := Element{ApplicationProtocol: proto, Value: ports}
        securityPolicyPayload.Element = []Element{element}

        endpointURL := "/api/2.0/services/application/" + scopeID
        this.BaseAPI = api.NewBaseAPI(http.MethodPost, endpointURL, securityPolicyPayload, new(string))
        return this
}

// GetResponse returns the ResponseObject from UpdateServiceAPI
func (updateAPI UpdateServiceAPI) GetResponse() string {
        return updateAPI.ResponseObject().(string)
}
