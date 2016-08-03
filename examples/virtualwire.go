package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/virtualwire"
)

// RunVirtualWireExample  Implementes VirtualWire example.
func RunVirtualWireExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All VirtualWires (Logical Switches).
	//
	api := virtualwire.NewGetAll("vdnscope-19")
	nsxclient.Do(api)
	// Get ID of our VirtualWire with name "test"
	virtualWireID := (api.GetResponse().FilterByName("test").ObjectID)
	// check if we got virtual wire id, or create a new one.
	if virtualWireID == "" {
		createVirtualWireAPI := virtualwire.NewCreate("test", "test desc", "tenant id", "vdnscope-19")
		nsxclient.Do(createVirtualWireAPI)
		fmt.Println("Status code:", createVirtualWireAPI.StatusCode())
		virtualWireID = createVirtualWireAPI.GetResponse()
	}

	//
	// Updating a virtualwire
	//

	updateAPI := virtualwire.NewUpdate("test2", "test2 desc", virtualWireID)
	nsxclient.Do(updateAPI)

	if updateAPI.StatusCode() == 200 {
		fmt.Println("Updated virtual wire")
	} else {
		fmt.Println("Failed to update virtualwire")
		fmt.Println("Status code:", updateAPI.StatusCode())
	}

	//
	// Deleting a virtual wire.
	//
	deleteAPI := virtualwire.NewDelete(virtualWireID)
	nsxclient.Do(deleteAPI)

	// check if it was a successful.
	if deleteAPI.StatusCode() == 204 {
		fmt.Println("Virtual Wire deleted.")
	} else {
		fmt.Println("Failed to delete virtual wire.")
		fmt.Println("Status code: ", deleteAPI.StatusCode())
	}

}
