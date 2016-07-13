package main

import (
	"fmt"
	"os"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/virtualwire"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("syntax error\nUsages: %s [NSX Manager Address] [Username] [Password] [INput file]\n\n", os.Args[0])
		os.Exit(1)
	}
	nsxManager := os.Args[1]
	nsxUser := os.Args[2]
	nsxPassword := os.Args[3]

	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient("https://"+nsxManager, nsxUser, nsxPassword, true, true)

	//
	// Get All VirtualWires (Logical Switches).
	//
	api := virtualwire.NewGetAll("vdnscope-19")
	nsxclient.Do(api)
	// Get ID of our VirtualWire with name "test"
	virtual_wire_id := (api.GetResponse().FilterByName("test").ObjectID)
	// check if we got virtual wire id, or create a new one.
	if virtual_wire_id == "" {
		create_api := virtualwire.NewCreate("test", "test desc", "tenant id", "vdnscope-19")
		nsxclient.Do(create_api)
		fmt.Println("Status code:", create_api.StatusCode())
		virtual_wire_id = create_api.GetResponse()
	}

	//
	// Deleting a virtual wire.
	//
	delete_api := virtualwire.NewDelete(virtual_wire_id)
	nsxclient.Do(delete_api)

	// check if it was a successful.
	if delete_api.StatusCode() == 204 {
		fmt.Println("Virtual Wire deleted.")
	} else {
		fmt.Println("Failed to delete virtual wire.")
		fmt.Println("Status code: ", delete_api.StatusCode())
	}


}
