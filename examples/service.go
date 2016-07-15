package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/service"
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

	//
	// Deleting a virtual wire.
	//

	// Let's refresh the getAllAPI call, so that it has the last created data.
	err = nsxclient.Do(getAllAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	applicationIDToDelete := getAllAPI.GetResponse().FilterByName("test")
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
