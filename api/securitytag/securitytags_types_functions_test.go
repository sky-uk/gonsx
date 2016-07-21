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
	}
	secondSecurityTag := SecurityTag{
		Name:        "securityTag2",
		ObjectID:    "securitytag-2",
		TypeName:    "SecurityTag",
		Description: "test",
	}
	securityTagsList.SecurityTags = []SecurityTag{firstSecurityTag, secondSecurityTag}
	return securityTagsList
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

func TestStringImplementation(t *testing.T) {
	securityTags := setup()
	assert.Equal(t, "[{{ } securitytag-1 securityTag1 test SecurityTag} {{ } securitytag-2 securityTag2 test SecurityTag}]", securityTags.String())
}
