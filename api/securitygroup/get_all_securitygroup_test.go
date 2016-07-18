package securitygroup

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllSecurityGroupAPI *GetAllSecurityGroupAPI

func setupGetAll() {
	getAllSecurityGroupAPI = NewGetAll("globalroot-0")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllSecurityGroupAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/services/securitygroup/scope/globalroot-0", getAllSecurityGroupAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<list><securitygroup><objectId>securitygroup-178</objectId><objectTypeName>SecurityGroup</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>5</revision><type><typeName>SecurityGroup</typeName></type><name>OVP_sg2</name><description></description><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><inheritanceAllowed>false</inheritanceAllowed><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.SECURITY_TAG</key><criteria>=</criteria><value>app4</value><isValid>true</isValid></dynamicCriteria></dynamicSet><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.NAME</key><criteria>contains</criteria><value>test</value><isValid>true</isValid></dynamicCriteria><dynamicCriteria><operator>OR</operator><key>VM.NAME</key><criteria>ends_with</criteria><value>test</value><isValid>true</isValid></dynamicCriteria></dynamicSet></dynamicMemberDefinition></securitygroup><securitygroup><objectId>securitygroup-1</objectId><objectTypeName>SecurityGroup</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>SecurityGroup</typeName></type><name>Activity Monitoring Data Collection</name><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><inheritanceAllowed>false</inheritanceAllowed></securitygroup><securitygroup><objectId>securitygroup-177</objectId><objectTypeName>SecurityGroup</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>2</revision><type><typeName>SecurityGroup</typeName></type><name>OVP_sg1</name><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><inheritanceAllowed>false</inheritanceAllowed><dynamicMemberDefinition><dynamicSet><operator>OR</operator><dynamicCriteria><operator>OR</operator><key>VM.SECURITY_TAG</key><criteria>contains</criteria><value>app3</value><isValid>true</isValid></dynamicCriteria></dynamicSet></dynamicMemberDefinition></securitygroup></list>")

	xmlErr := xml.Unmarshal(xmlContent, getAllSecurityGroupAPI.ResponseObject())

	assert.Nil(t, xmlErr)
	assert.Len(t, getAllSecurityGroupAPI.GetResponse().SecurityGroups, 3)
	assert.Equal(t, "OVP_sg2", getAllSecurityGroupAPI.GetResponse().SecurityGroups[0].Name)
	// bit deep level checks to test that we un-marshalling the whole object.
	assert.Len(t, getAllSecurityGroupAPI.GetResponse().SecurityGroups[0].DynamicMemberDefinition.DynamicSet, 2)
	assert.Equal(t, "OR", getAllSecurityGroupAPI.GetResponse().SecurityGroups[0].DynamicMemberDefinition.DynamicSet[0].Operator)
	assert.Equal(t, "app4", getAllSecurityGroupAPI.GetResponse().SecurityGroups[0].DynamicMemberDefinition.DynamicSet[0].DynamicCriteria[0].Value)
}
