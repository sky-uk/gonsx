package ipset

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllIpSetAPI *GetAllIpSetAPI

func setupGetAll() {
	getAllIpSetAPI = NewGetAll("globalroot-0")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllIpSetAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/services/ipset/scope/globalroot-0", getAllIpSetAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte(`
<list>
    <ipset>
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
    </ipset>
</list>`)

	xmlErr := xml.Unmarshal(xmlContent, getAllIpSetAPI.ResponseObject())

	assert.Nil(t, xmlErr)
	assert.Equal(t, "test_ipset_name", getAllIpSetAPI.GetResponse().IpSets[0].Name)
	assert.Equal(t, "test_ipset_description", getAllIpSetAPI.GetResponse().IpSets[0].Description)
	assert.Equal(t, "10.50.0.0/8", getAllIpSetAPI.GetResponse().IpSets[0].Value)
}
