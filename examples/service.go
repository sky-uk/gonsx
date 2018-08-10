package main

import (
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/service"
)

// RunServiceExample  Implementes Service example.
func RunServiceExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All Services.
	//
	// Create api object.
	getAllAPI := service.NewGetAll("globalroot-0")

	// make api call.
	err := nsxclient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllApplications := getAllAPI.GetResponse().Applications
		for _, service := range AllApplications {
			fmt.Printf("objectId: %-20s name: %-20s\n", service.ObjectID, service.Name)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	//
	// Get Single Service
	//
	// Get All ( we re-utilize the GetAll object from above here )
	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		service := getAllAPI.GetResponse().FilterByName("OVP_test1")
		if service.ObjectID != "" {
			fmt.Println(service)
		} else {
			fmt.Println("Not found!")
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	//
	// Create single service.
	//
	createAPI := service.NewCreate("globalroot-0", "test", "desc", "TCP", "8080")
	err = nsxclient.Do(createAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if createAPI.StatusCode() == 201 {
		applicationID := createAPI.ResponseObject()
		fmt.Println("Service created successfully.")
		fmt.Println("objectId:", applicationID)
	} else {
		fmt.Println("Failed to created the service!")
		fmt.Println(createAPI.ResponseObject())
	}

	// UPDATE
	//
	// Updating a single service.
	// Get list of all applications. Search through looking for application match.
	// Update the attribute/s of the service.
	getAllAPI = service.NewGetAll("globalroot-0")

	// make api call.
	err = nsxclient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() != 200 {
		fmt.Printf("Status code: %v, Response: %v\n", getAllAPI.StatusCode(), getAllAPI.ResponseObject())
	}

	// Get All ( we re-utilize the GetAll object from above here )
	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		service := getAllAPI.GetResponse().FilterByName("test")
		if service.ObjectID != "" {
			fmt.Println("Found service: ", service.ObjectID, service.Name)
		} else {
			fmt.Println("Not found!")
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	// Change the name of the service from test to test_https and change the port to TCP/443.
	serviceToModify := getAllAPI.GetResponse().FilterByName("test")
	serviceToModify.Name = "test_https"
	modifyElement := service.Element{ApplicationProtocol: "TCP", Value: "443"}
	serviceToModify.Element = []service.Element{modifyElement}
	updateAPI := service.NewUpdate(serviceToModify.ObjectID, serviceToModify)

	err = nsxclient.Do(updateAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if updateAPI.StatusCode() == 200 {
		newObject := updateAPI.GetResponse()
		fmt.Println("Service updated successfully.")
		fmt.Println("objectId:", newObject.ObjectID)
	} else {
		fmt.Println("Failed to update the service!")
		fmt.Println(updateAPI.ResponseObject())
	}

	//
	// Deleting a single service.
	//

	// Let's refresh the getAllAPI call, so that it has the last created data.
	err = nsxclient.Do(getAllAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	applicationIDToDelete := getAllAPI.GetResponse().FilterByName("test_https")
	deleteAPI := service.NewDelete(applicationIDToDelete.ObjectID)
	err = nsxclient.Do(deleteAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if deleteAPI.StatusCode() == 200 {
		fmt.Println("Service deleted successfully.")
	} else {
		fmt.Println("Failed to delete the service!")
		fmt.Println("Status code:", deleteAPI.StatusCode())
		fmt.Println("Response:", deleteAPI.ResponseObject())
	}

}
