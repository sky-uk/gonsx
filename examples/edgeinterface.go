package main

import (
	"bufio"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/edgeinterface"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
	"os"
)

// RunEdgeinterfaceExample - Runs the edge interface example.
func RunEdgeinterfaceExample(nsxManager, nsxUser, nsxPassword string, debug bool) {

	reader := bufio.NewReader(os.Stdin)

	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	scopeID := "vdnscope-1"
	//
	// Get All VirtualWires.
	//
	fmt.Println("Getting all virtual wires in " + scopeID + "....")
	api := virtualwire.NewGetAll(scopeID)
	nsxclient.Do(api)

	// Get ID of our virtualwire with name "test"
	virtualWireID := (api.GetResponse().FilterByName("test").ObjectID)

	// check if we got virtual wire id, otherwise let's create one and get the ID.
	if virtualWireID == "" {
		fmt.Println("VirtualWire name test doesn't exist, going to create it...")
		virtualWireCreateSpec := virtualwire.CreateSpec{Name: "gonsx-example", Description: "gonsx example edge interface", TenantID: "tenant_id", ControlPlaneMode: "UNICAST_MODE"}
		createAPI := virtualwire.NewCreate(virtualWireCreateSpec, scopeID)
		nsxclient.Do(createAPI)
		fmt.Println("Status code:", createAPI.StatusCode())
		virtualWireID = createAPI.GetResponse()
	}

	fmt.Println("Virtual Wire name test ID: ", virtualWireID)

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

	fmt.Println("Going to create edgeinterface with id edge-7, proceed?")
	reader.ReadString('\n')

	edgeInterfaceAPI := edgeinterface.NewCreate(requestPayload, "edge-7")

	nsxclient.Do(edgeInterfaceAPI)

	edgeCreated := new(edgeinterface.EdgeInterfaces)

	// Check the status code and process accordingly.
	if edgeInterfaceAPI.StatusCode() == http.StatusOK {
		edgeCreated = edgeInterfaceAPI.GetResponse()
		fmt.Println("Interface created, Response:\n", edgeCreated)
	} else {
		fmt.Println("Failed to create interface.")
		os.Exit(1)
	}

	//
	// Deleting Edge Interface.
	//
	//  first get all edge interfaces.
	fmt.Printf("Going to delete interface of index: %d", edgeCreated.Interfaces[0].Index)
	getEdgeAPI := edgeinterface.NewGet("edge-7", edgeCreated.Interfaces[0].Index)
	nsxclient.Do(getEdgeAPI)

	if getEdgeAPI.StatusCode() == 200 {
		// Find the one we are interested in.
		interfaceObj := getEdgeAPI.GetResponse()

		fmt.Println("Found edge interface, index: ", interfaceObj.Index)

		// create delete call object.
		fmt.Printf("Going to delete edgeinterface with index %d, proceed?", interfaceObj.Index)
		reader.ReadString('\n')

		edgeDeleteAPI := edgeinterface.NewDelete("edge-7", interfaceObj.Index)
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
