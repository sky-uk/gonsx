package securitypolicy

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createSecurityPolicyAPI *CreateSecurityPolicyAPI

func createSetup() {
	securityGroupIDs := []string{"securitygroup-0001", "securitygroup-0002"}
	createSecurityPolicyAPI = NewCreate("OVP_test_security_policy", "5501", "this is a long description.", securityGroupIDs)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createSecurityPolicyAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/2.0/services/policy/securitypolicy", createSecurityPolicyAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<securityPolicy><name>OVP_test_security_policy</name><description>this is a long description.</description><precedence>5501</precedence><actionsByCategory><category></category></actionsByCategory><securityGroupBinding><objectId>securitygroup-0001</objectId></securityGroupBinding><securityGroupBinding><objectId>securitygroup-0002</objectId></securityGroupBinding></securityPolicy>"

	xmlBytes, err := xml.Marshal(createSecurityPolicyAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestCreateUnMarshalling(t *testing.T) {
	// TODO
}
