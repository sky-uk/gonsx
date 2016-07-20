package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func setup() (applicationsList *ApplicationsList) {
	applicationsList = &ApplicationsList{}
	firstApplicationService := ApplicationService{
		Name:     "Test_80",
		ObjectID: "application-001",
	}
	secondApplicationService := ApplicationService{
		Name:     "Test_8080",
		ObjectID: "application-002",
	}
	applicationsList.Applications = []ApplicationService{firstApplicationService, secondApplicationService}
	return applicationsList
}

func TestFilterByName(t *testing.T) {
	applicationsList := setup()

	firstFiltered := applicationsList.FilterByName("Test_80")
	assert.Equal(t, "application-001", firstFiltered.ObjectID)

	secondFiltered := applicationsList.FilterByName("Test_8080")
	assert.Equal(t, "application-002", secondFiltered.ObjectID)
}

func TestCheckByName(t *testing.T) {
	applicationsList := setup()

	firstCheck := applicationsList.CheckByName("Test_80")
	assert.Equal(t, true, firstCheck)

	secondCheck := applicationsList.CheckByName("Not_Existing")
	assert.Equal(t, false, secondCheck)
}

func TestStringImplementation(t *testing.T) {
	applicationsList := setup()

	stringOutputOfList := fmt.Sprintln(applicationsList)
	assert.Equal(t, "ApplicationsList object, contains service objects.\n", stringOutputOfList)

	stringOutputOfObject := fmt.Sprintln(applicationsList.Applications[0])
	assert.Equal(t, "objectId: application-001      name: Test_80             .\n", stringOutputOfObject)

}