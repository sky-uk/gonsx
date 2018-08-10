package main

import (
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
)

// RunVirtualWireExample  Implementes VirtualWire example.
func RunVirtualWireExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	scopeID := "vdnscope-1"

	//
	// Get All VirtualWires (Logical Switches).
	//
	api := virtualwire.NewGetAll(scopeID)
	nsxclient.Do(api)
	// Get ID of our VirtualWire with name "test"
	virtualWireID := (api.GetResponse().FilterByName("gonsx-example").ObjectID)
	// check if we got virtual wire id, or create a new one.
	if virtualWireID == "" {
		virtualWireCreateSpec := virtualwire.CreateSpec{Name: "gonsx-example", Description: "gonsx example logical switch", TenantID: "tenant_id", ControlPlaneMode: "UNICAST_MODE"}
		createVirtualWireAPI := virtualwire.NewCreate(virtualWireCreateSpec, scopeID)
		nsxclient.Do(createVirtualWireAPI)
		fmt.Println("Status code:", createVirtualWireAPI.StatusCode())
		virtualWireID = createVirtualWireAPI.GetResponse()
	}

	//
	// Updating a virtualwire
	//
	virtualWireUpdate := virtualwire.VirtualWire{Name: "gonsx-example-updated", ObjectID: virtualWireID, ControlPlaneMode: "UNICAST_MODE", Description: "gonsx example logical switch updated", TenantID: "tenant_id"}
	updateAPI := virtualwire.NewUpdate(virtualWireUpdate)
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
	if deleteAPI.StatusCode() == http.StatusOK {
		fmt.Println("Virtual Wire deleted.")
	} else {
		fmt.Println("Failed to delete virtual wire.")
		fmt.Println("Status code: ", deleteAPI.StatusCode())
	}

}
