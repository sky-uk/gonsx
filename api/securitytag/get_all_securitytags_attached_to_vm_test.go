package securitytag

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllSecurityTagsAttachedToVMAPI *GetAllSecurityTagsAttachedToVMAPI

func setUpGetAllAttachedToVM() {
	getAllSecurityTagsAttachedToVMAPI = NewGetAllAttachedToVM("vm-246")
}

func TestGetAllAttachedToVmMethod(t *testing.T) {
	setUpGetAllAttachedToVM()
	assert.Equal(t, http.MethodGet, getAllSecurityTagsAttachedToVMAPI.Method())
}

func TestNewGetAllAttachedToVMEndpoint(t *testing.T) {
	setUpGetAllAttachedToVM()
	assert.Equal(t, "/api/2.0/services/securitytags/vm/vm-246", getAllSecurityTagsAttachedToVMAPI.Endpoint())
}

func TestGetAllSecurityTagsAttachedToVmUnMarshalling(t *testing.T) {
	setUpGetAllAttachedToVM()
	xmlContent := []byte("<securityTags><securityTag><objectId>securitytag-127</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</vsmUuid><nodeId>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</nodeId><revision>1</revision><type><typeName>SecurityTag</typeName></type><name>cdtest_tag_1</name><description>Craigs test tag 1</description><clientHandle/><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>1</vmCount></securityTag><securityTag><objectId>securitytag-128</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</vsmUuid><nodeId>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</nodeId><revision>4</revision><type><typeName>SecurityTag</typeName></type><name>cdtest_tag_2</name><description>Craigs test tag 2</description><clientHandle/><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>1</vmCount></securityTag><securityTag><objectId>securitytag-129</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</vsmUuid><nodeId>42022C8A-4FEE-8BCB-1066-FE34E3ED0D84</nodeId><revision>2</revision><type><typeName>SecurityTag</typeName></type><name>cdtest_tag_3</name><description>Craigs test tag 3</description><clientHandle/><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>1</vmCount></securityTag></securityTags>")
	xmlErr := xml.Unmarshal(xmlContent, getAllSecurityTagsAttachedToVMAPI.ResponseObject())

	assert.Nil(t, xmlErr)
	assert.Len(t, getAllSecurityTagsAttachedToVMAPI.GetResponse().SecurityTags, 3)
	assert.Equal(t, "securitytag-127", getAllSecurityTagsAttachedToVMAPI.GetResponse().SecurityTags[0].ObjectID)

}
