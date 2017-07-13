package sections

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

)

var  testSection Section
var   createFWSectionAPI *CreateFWSectionsAPI

func setupCreateSection() {
	testSection.Name = "My Test Section"
	testSection.ID = "1"
	testSection.Type = "LAYER3"
	createFWSectionAPI = NewCreate(testSection)
}

func setupCreateL2Section() {
	testSection.Name = "My Test Section"
	testSection.ID = "1"
	testSection.Type = "LAYER2"
	createFWSectionAPI = NewCreate(testSection)
}



func TestCreateMethod(t *testing.T) {
	setupCreateSection()
	assert.Equal(t, http.MethodPost, createFWSectionAPI.Method())
}

func TestL3Endpoint(t *testing.T){
	setupCreateSection()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer3sections", createFWSectionAPI.Endpoint())
}

func TestL2Endpoint(t *testing.T){
	setupCreateL2Section()
	assert.Equal(t, "/api/4.0/firewall/globalroot-0/config/layer2sections", createFWSectionAPI.Endpoint())
}



func TestMethod(t *testing.T){
	setupCreateSection()
	assert.Equal(t, http.MethodPost, createFWSectionAPI.Method())
}


func TestResponse(t *testing.T) {
	setupCreateSection()
	createFWSectionAPI.SetResponseObject(&testSection)
	assert.Equal(t, testSection, createFWSectionAPI.GetResponse())
}