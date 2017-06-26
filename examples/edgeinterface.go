package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/edgeinterface"
	"github.com/sky-uk/gonsx/api/virtualwire"
)

// RunEdgeinterfaceExample - Runs the edge interface example.
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
	virtualWireID := (api.GetResponse().FilterByName("test").ObjectID)

	// check if we got virtual wire id, otherwise let's create one and get the ID.
	if virtualWireID == "" {
		createAPI := virtualwire.NewCreate("test", "test desc", "tenant id", "vdnscope-19")
		nsxclient.Do(createAPI)
		fmt.Println("Status code:", createAPI.StatusCode())
		virtualWireID = createAPI.GetResponse()
	}

	//
	// Create Edge Interface for the virtual wire we created above
	//
	addressGroup := edgeinterface.AddressGroup{PrimaryAddress: "10.10.10.1", SubnetMask: "255.255.255.0"}
	addressGroupList := []edgeinterface.AddressGroup{addressGroup}

	edgeInterface := edgeinterface.EdgeInterface{
		Name:          "app_virtualwire_one",
		ConnectedToID: virtualWireID,
		Type:          "internal",
		Mtu:           1500,
		IsConnected:   true,
		AddressGroups: edgeinterface.AddressGroups{addressGroupList},
	}

	requestPayload := new(edgeinterface.EdgeInterfaces)
	requestPayload.Interfaces = []edgeinterface.EdgeInterface{edgeInterface}

	edgeInterfaceAPI := edgeinterface.NewCreate(requestPayload, "edge-50")

	nsxclient.Do(edgeInterfaceAPI)

	// Check the status code and process accordingly.
	if edgeInterfaceAPI.StatusCode() == 200 {
		fmt.Println("Status code:", edgeInterfaceAPI.StatusCode())
		fmt.Println("Response:", edgeInterfaceAPI.GetResponse())
	} else {
		fmt.Println("Status code:", edgeInterfaceAPI.StatusCode())
		fmt.Println("Failed to create interface.")
	}

	//
	// Deleting Edge Interface.
	//
	//  first get all edge interfaces.
	getAllEdgeInterfacesAPI := edgeinterface.NewGetAll("edge-50")
	nsxclient.Do(getAllEdgeInterfacesAPI)

	if getAllEdgeInterfacesAPI.StatusCode() == 200 {
		// Find the one we are interested in.
		interfaceObj := (getAllEdgeInterfacesAPI.GetResponse().FilterByName("app_virtualwire_one"))

		fmt.Print(interfaceObj)

		// create delete call object.
		edgeDeleteAPI := edgeinterface.NewDelete(interfaceObj.Index, "edge-50")
		nsxclient.Do(edgeDeleteAPI)

		// check if it was a successful.
		if edgeDeleteAPI.StatusCode() == 204 {
			fmt.Println("Deleted edge interface")
		} else {
			fmt.Println("Failed to delete edge inteface.")
			fmt.Println("Status code: ", edgeDeleteAPI.StatusCode())
		}
	}

}
