package securitytag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() (securityTagsList *SecurityTags) {
	securityTagsList = &SecurityTags{}
	firstSecurityTag := SecurityTag{
		Name:        "securityTag1",
		ObjectID:    "securitytag-1",
		TypeName:    "SecurityTag",
		Description: "test",
		Revision:    1,
	}
	secondSecurityTag := SecurityTag{
		Name:        "securityTag2",
		ObjectID:    "securitytag-2",
		TypeName:    "SecurityTag",
		Description: "test",
		Revision:    2,
	}
	securityTagsList.SecurityTags = []SecurityTag{firstSecurityTag, secondSecurityTag}
	return securityTagsList
}

func setupAttached() (basicInfoList *BasicInfoList) {
	basicInfoList = &BasicInfoList{}
	firstBasicInfo := BasicInfo{
		Name:     "vm1",
		ObjectID: "vm-1",
	}
	secondBasicInfo := BasicInfo{
		Name:     "vm2",
		ObjectID: "vm-2",
	}
	basicInfoList.BasicInfoList = []BasicInfo{firstBasicInfo, secondBasicInfo}
	return basicInfoList
}

func setupAttachedToVM() (securityTagAttachmentList *AttachmentList) {
	securityTagAttachmentList = new(AttachmentList)
	firstAttachedSecurityTag := Attachment{ObjectID: "securitytag-127"}
	secondAttachedSecurityTag := Attachment{ObjectID: "securitytag-128"}
	thirdAttachedSecurityTag := Attachment{ObjectID: "securitytag-129"}
	securityTagAttachmentList.SecurityTagAttachments = []Attachment{firstAttachedSecurityTag, secondAttachedSecurityTag, thirdAttachedSecurityTag}
	return securityTagAttachmentList
}

func TestFilterByName(t *testing.T) {
	securityTags := setup()

	firstFiltered := securityTags.FilterByName("securityTag1")
	assert.Equal(t, "securitytag-1", firstFiltered.ObjectID)

	secondFiltered := securityTags.FilterByName("securityTag2")
	assert.Equal(t, "securitytag-2", secondFiltered.ObjectID)
}

func TestCheckByName(t *testing.T) {
	securityTags := setup()

	firstCheck := securityTags.CheckByName("securityTag1")
	assert.Equal(t, true, firstCheck)

	secondCheck := securityTags.CheckByName("securityTagNo")
	assert.Equal(t, false, secondCheck)
}

func TestFilterByNameAttached(t *testing.T) {
	basicInfo := setupAttached()

	firstFiltered := basicInfo.FilterByIDAttached("vm-1")
	assert.Equal(t, "vm1", firstFiltered.Name)

	secondFiltered := basicInfo.FilterByIDAttached("vm-2")
	assert.Equal(t, "vm2", secondFiltered.Name)
}

func TestStringImplementationSecurityTags(t *testing.T) {
	securityTags := setup()
	assert.Equal(t, "Security tags contains a list of securitytags", securityTags.String())
}

func TestStingImplementationSecurityTag(t *testing.T) {
	firstSecurityTag := SecurityTag{
		Name:        "securityTag1",
		ObjectID:    "securitytag-1",
		TypeName:    "SecurityTag",
		Description: "test",
		Revision:    1,
	}
	assert.Equal(t, "Security tag name securityTag1 and id securitytag-1", firstSecurityTag.String())
}

func TestCheckByObjectID(t *testing.T) {
	securityTagAttachmentList := setupAttachedToVM()
	firstCheck := securityTagAttachmentList.CheckByObjectID("securitytag-127")
	assert.Equal(t, true, firstCheck)
	secondCheck := securityTagAttachmentList.CheckByObjectID("doesNotExist")
	assert.Equal(t, false, secondCheck)
}

func TestVerifyAttachments(t *testing.T) {
	securityTagAttachmentList := setupAttachedToVM()

	listToVerifyOne := new(SecurityTags)
	firstAttachedSecurityTag := SecurityTag{ObjectID: "securitytag-127"}
	secondAttachedSecurityTag := SecurityTag{ObjectID: "securitytag-128"}
	thirdAttachedSecurityTag := SecurityTag{ObjectID: "securitytag-129"}
	listToVerifyOne.SecurityTags = []SecurityTag{firstAttachedSecurityTag, secondAttachedSecurityTag, thirdAttachedSecurityTag}

	assert.Equal(t, []string(nil), securityTagAttachmentList.VerifyAttachments(listToVerifyOne))

	listToVerifyTwo := new(SecurityTags)
	fourthAttachedSecurityTag := SecurityTag{ObjectID: "securitytag-135"}
	fifthAttachedSecurityTag := SecurityTag{ObjectID: "securitytag-150"}
	listToVerifyTwo.SecurityTags = []SecurityTag{fourthAttachedSecurityTag, fifthAttachedSecurityTag}
	assert.Equal(t, []string{"securitytag-135", "securitytag-150"}, securityTagAttachmentList.VerifyAttachments(listToVerifyTwo))
}
