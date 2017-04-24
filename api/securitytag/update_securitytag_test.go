package securitytag

import (
	//"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateSecurityTagAPI *UpdateSecurityTagAPI

func updateSetup() {
	SecurityTagID := "testtag-1"
	SecurityTagName :="testTag"
	SecurityTagDescription := "A description"
	updateSecurityTagAPI = NewUpdate(SecurityTagID,SecurityTagName, SecurityTagDescription)

}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateSecurityTagAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/2.0/services/securitytags/tag", updateSecurityTagAPI.Endpoint() )
}

