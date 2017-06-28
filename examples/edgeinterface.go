package main

import (
	"bufio"
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/edgeinterface"
	"github.com/sky-uk/gonsx/api/virtualwire"
	"os"
)

// RunEdgeinterfaceExample - Runs the edge interface example.
func RunEdgeinterfaceExample(nsxManager, nsxUser, nsxPassword string, debug bool) {

	reader := bufio.NewReader(os.Stdin)

	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All VirtualWires.
	//
	fmt.Println("Getting all virtual wires in vdnscope-1...")
	api := virtualwire.NewGetAll("vdnscope-1")
	nsxclient.Do(api)

	// Get ID of our virtualwire with name "test"
	virtualWireID := (api.GetResponse().FilterByName("test").ObjectID)

	// check if we got virtual wire id, otherwise let's create one and get the ID.
	if virtualWireID == "" {
		fmt.Println("VirtualWire name test doesn't exist, going to create it...")
		createAPI := virtualwire.NewCreate("test", "test desc", "tenant id", "vdnscope-1")
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

	fmt.Println("Status code:", edgeInterfaceAPI.StatusCode())
	// Check the status code and process accordingly.
	if edgeInterfaceAPI.StatusCode() == 200 {
		fmt.Println("Interface created, Response:\n", edgeInterfaceAPI.GetResponse())
	} else {
		fmt.Println("Failed to create interface.")
	}

	//
	// Deleting Edge Interface.
	//
	//  first get all edge interfaces.
	getAllEdgeInterfacesAPI := edgeinterface.NewGetAll("edge-7")
	nsxclient.Do(getAllEdgeInterfacesAPI)

	if getAllEdgeInterfacesAPI.StatusCode() == 200 {
		// Find the one we are interested in.
		interfaceObj := (getAllEdgeInterfacesAPI.GetResponse().FilterByName("app_virtualwire_one"))

        if interfaceObj != nil {
            fmt.Println("Found edge interface, index: ", interfaceObj.Index)

		    // create delete call object.
		    fmt.Printf("Going to delete edgeinterface with index %s, proceed?", interfaceObj.Index)
		    reader.ReadString('\n')

		    edgeDeleteAPI := edgeinterface.NewDelete(interfaceObj.Index, "edge-7")
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

}
