package ipset

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getIpSetAPI *GetIpSetAPI

func setupGet(id string) {
	getIpSetAPI = NewGet(id)
}

func TestGetMethod(t *testing.T) {
	setupGet("ipset-27")
	assert.Equal(t, http.MethodGet, getIpSetAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet("ipset-27")
	assert.Equal(t, "/api/2.0/services/ipset/ipset-27", getIpSetAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet("ipset-27")
	xmlContent := []byte(`<ipset>
			<objectId>ipset-27</objectId>
			<objectTypeName>IPSet</objectTypeName>
			<vsmUuid>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</vsmUuid>
			<nodeId>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</nodeId>
			<revision>2</revision>
			<type>
					<typeName>IPSet</typeName>
			</type>
			<name>test_ipset_name</name>
			<description>test_ipset_description</description>
			<scope>
					<id>globalroot-0</id>
					<objectTypeName>GlobalRoot</objectTypeName>
					<name>Global</name>
			</scope>
			<clientHandle></clientHandle>
			<extendedAttributes/>
			<isUniversal>false</isUniversal>
			<universalRevision>0</universalRevision>
			<inheritanceAllowed>false</inheritanceAllowed>
			<value>10.50.0.0/8</value>
	</ipset>`)

	xmlErr := xml.Unmarshal(xmlContent, getIpSetAPI.ResponseObject())

	assert.Nil(t, xmlErr)
	assert.Equal(t, "test_ipset_name", getIpSetAPI.GetResponse().Name)
	assert.Equal(t, "test_ipset_description", getIpSetAPI.GetResponse().Description)
	assert.Equal(t, "10.50.0.0/8", getIpSetAPI.GetResponse().Value)
}
