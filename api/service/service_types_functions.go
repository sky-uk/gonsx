package service

import (
	"fmt"
)

func (a ApplicationsList) String() string {
	return fmt.Sprintf("ApplicationsList object, contains service objects.")
}

func (a ApplicationService) String() string {
	return fmt.Sprintf("objectId: %-20s name: %-20s", a.ObjectID, a.Name)
}

// FilterByName returns a single service object if it matches the name in ApplicationsList
func (a ApplicationsList) FilterByName(name string) *ApplicationService {
	var serviceFound ApplicationService
	for _, service := range a.Applications {
		if service.Name == name {
			serviceFound = service
			break
		}
	}
	return &serviceFound
}
