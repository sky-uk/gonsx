package securitytags

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var detachSecurityTagAPI *DetachSecurityTagAPI

func setupDetach(){
	detachSecurityTagAPI = NewDetach("securitytag-1", "vm-1")
}

func TestNewDetach(t *testing.T) {
	setupDetach()
	assert.Equal(t, http.MethodDelete, detachSecurityTagAPI.Method())
}

func TestDetachEndpoint(t *testing.T){
	setupDetach()
	assert.Equal(t, "/api/2.0/services/securitytags/tag/securitytag-1/vm/vm-1", detachSecurityTagAPI.Endpoint())
}
