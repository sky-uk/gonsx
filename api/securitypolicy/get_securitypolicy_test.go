package securitypolicy

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSecurityPolicyAPI *GetSecurityPolicyAPI

func setupGet(securityPolicyID string) {
	getSecurityPolicyAPI = NewGet(securityPolicyID)
}

func TestGetMethod(t *testing.T) {
	setupGet("securitypolicy-001")
	assert.Equal(t, http.MethodGet, getSecurityPolicyAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("securitypolicy-001")
	assert.Equal(t, "/api/2.0/services/policy/securitypolicy/securitypolicy-001", getSecurityPolicyAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("securitypolicy-001")
	xmlContent := []byte("<securityPolicy><objectId>securitypolicy-001</objectId><name>OVP_test_security_policy</name><description>this is a long description.</description><precedence>5501</precedence><actionsByCategory><category></category></actionsByCategory><securityGroupBinding><objectId>securitygroup-0001</objectId></securityGroupBinding><securityGroupBinding><objectId>securitygroup-0002</objectId></securityGroupBinding></securityPolicy>")

	xmlerr := xml.Unmarshal(xmlContent, getSecurityPolicyAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "securitypolicy-001", getSecurityPolicyAPI.GetResponse().ObjectID)
	assert.Equal(t, "OVP_test_security_policy", getSecurityPolicyAPI.GetResponse().Name)
}
