package ipset

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getIpSetAPI *GetIpSetAPI

func setupGet() {
	getIpSetAPI = NewGet("globalroot-0")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getIpSetAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGet()
	assert.Equal(t, "/api/2.0/services/ipset/globalroot-0", getIpSetAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGet()
	xmlContent := []byte("<list><ipset><objectId>ipset-178</objectId><objectTypeName>IpSet</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>5</revision><type><typeName>IpSet</typeName></type><name>OVP_sg2</name><description></description><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><inheritanceAllowed>false</inheritanceAllowed><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.SECURITY_TAG</key><criteria>=</criteria><value>app4</value><isValid>true</isValid></dynamicCriteria></dynamicSet><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.NAME</key><criteria>contains</criteria><value>test</value><isValid>true</isValid></dynamicCriteria><dynamicCriteria><operator>OR</operator><key>VM.NAME</key><criteria>ends_with</criteria><value>test</value><isValid>true</isValid></dynamicCriteria></dynamicSet></dynamicMemberDefinition></ipset><ipset><objectId>ipset-1</objectId><objectTypeName>IpSet</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>IpSet</typeName></type><name>Activity Monitoring Data Collection</name><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><inheritanceAllowed>false</inheritanceAllowed></ipset><ipset><objectId>ipset-177</objectId><objectTypeName>IpSet</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>2</revision><type><typeName>IpSet</typeName></type><name>OVP_sg1</name><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><inheritanceAllowed>false</inheritanceAllowed><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.SECURITY_TAG</key><criteria>contains</criteria><value>app3</value><isValid>true</isValid></dynamicCriteria></dynamicSet></dynamicMemberDefinition></ipset></list>")

	xmlErr := xml.Unmarshal(xmlContent, getIpSetAPI.ResponseObject())

	assert.Nil(t, xmlErr)
	assert.Len(t, getIpSetAPI.GetResponse().IpSets, 3)
	assert.Equal(t, "OVP_sg2", getIpSetAPI.GetResponse().IpSets[0].Name)
	// bit deep level checks to test that we un-marshalling the whole object.
	assert.Len(t, getIpSetAPI.GetResponse().IpSets[0].DynamicMemberDefinition.DynamicSet, 2)
	assert.Equal(t, "OR", getIpSetAPI.GetResponse().IpSets[0].DynamicMemberDefinition.DynamicSet[0].Operator)
	assert.Equal(t, "app4", getIpSetAPI.GetResponse().IpSets[0].DynamicMemberDefinition.DynamicSet[0].DynamicCriteria[0].Value)
}
