package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/edgeinterface"
	"github.com/sky-uk/gonsx/api/virtualwire"
)

func RunEdgeinterfaceExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All VirtualWires.
	//
	api := virtualwire.NewGetAll("vdnscope-19")
	nsxclient.Do(api)

	// Get ID of our virtualwire with name "test"
	virtual_wire_id := (api.GetResponse().FilterByName("test").ObjectID)

	// check if we got virtual wire id, otherwise let's create one and get the ID.
	if virtual_wire_id == "" {
		create_api := virtualwire.NewCreate("test", "test desc", "tenant id", "vdnscope-19")
		nsxclient.Do(create_api)
		fmt.Println("Status code:", create_api.StatusCode())
		virtual_wire_id = create_api.GetResponse()
	}

	//
	// Create Edge Interface for the virtual wire we created above
	//
	edge_interface_api := edgeinterface.NewCreate("edge-50", "app_virtualwire_one", virtual_wire_id, "10.10.10.1", "255.255.255.0", "internal", 1500)

	nsxclient.Do(edge_interface_api)

	// Check the status code and process accordingly.
	if edge_interface_api.StatusCode() == 200 {
		fmt.Println("Status code:", edge_interface_api.StatusCode())
		fmt.Println("Response:", edge_interface_api.GetResponse())
	} else {
		fmt.Println("Status code:", edge_interface_api.StatusCode())
		fmt.Println("Failed to create interface.")
	}

	//
	// Deleting Edge Interface.
	//
	//  first get all edge interfaces.
	get_all_edge_interfaces := edgeinterface.NewGetAll("edge-50")
	nsxclient.Do(get_all_edge_interfaces)

	if get_all_edge_interfaces.StatusCode() == 200 {
		// Find the one we are interested in.
		interface_obj := (get_all_edge_interfaces.GetResponse().FilterByName("app_virtualwire_one"))

		fmt.Print(interface_obj)

		// create delete call object.
		edge_delete_api := edgeinterface.NewDelete(interface_obj.Index, "edge-50")
		nsxclient.Do(edge_delete_api)

		// check if it was a successful.
		if edge_delete_api.StatusCode() == 204 {
			fmt.Println("Deleted edge interface")
		} else {
			fmt.Println("Failed to delete edge inteface.")
			fmt.Println("Status code: ", edge_delete_api.StatusCode())
		}
	}

}
