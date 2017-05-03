package securitytag

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateAttachedSecurityTagsAPI *UpdateAttachedSecurityTagsAPI

func setUpUpdateAllAttachedSecurityTags() {
	vmID := "vm-426"
	securityTagAttachmentOne := Attachment{ObjectID: "securitytag-127"}
	securityTagAttachmentTwo := Attachment{ObjectID: "securitytag-128"}
	securityTagAttachmentThree := Attachment{ObjectID: "securitytag-129"}
	requestPayload := new(AttachmentList)
	requestPayload.AddSecurityTagToAttachmentList(securityTagAttachmentOne)
	requestPayload.AddSecurityTagToAttachmentList(securityTagAttachmentTwo)
	requestPayload.AddSecurityTagToAttachmentList(securityTagAttachmentThree)
	updateAttachedSecurityTagsAPI = NewUpdateAttachedTags(vmID, requestPayload)
}

func TestNewUpdateAttachedSecurityTagsMethod(t *testing.T) {
	setUpUpdateAllAttachedSecurityTags()
	assert.Equal(t, http.MethodPost, updateAttachedSecurityTagsAPI.Method())
}

func TestNewUpdateAttachedSecurityTagsEndpoint(t *testing.T) {
	setUpUpdateAllAttachedSecurityTags()
	assert.Equal(t, "/api/2.0/services/securitytags/vm/vm-426?action=ASSIGN_TAGS", updateAttachedSecurityTagsAPI.Endpoint())
}

func TestNewUpdateAttachedSecurityTagsMarshalling(t *testing.T) {
	setUpUpdateAllAttachedSecurityTags()
	expectedXML := "<securityTags><securityTag><objectId>securitytag-127</objectId></securityTag><securityTag><objectId>securitytag-128</objectId></securityTag><securityTag><objectId>securitytag-129</objectId></securityTag></securityTags>"
	xmlBytes, err := xml.Marshal(updateAttachedSecurityTagsAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}

func TestNewUpdateAttachedSecurityTagsResponse(t *testing.T) {
	setUpUpdateAllAttachedSecurityTags()
	updateAttachedSecurityTagsAPI.SetResponseObject("string output returned")
	assert.Equal(t, "string output returned", updateAttachedSecurityTagsAPI.GetResponse())
}
