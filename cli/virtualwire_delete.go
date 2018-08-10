package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
	"os"
)

var deleteVirtualWireID string

func deleteVirtualWire(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if deleteVirtualWireID == "" {
		fmt.Println("Error Virtualwire ID is required. Usage: -id virtualwire-XX")
		os.Exit(1)
	}

	deleteVirtualWireAPI := virtualwire.NewDelete(deleteVirtualWireID)
	err := client.Do(deleteVirtualWireAPI)
	if err != nil {
		fmt.Printf("\nError deleting virtual wire %s. Error %v\n", deleteVirtualWireAPI, err.Error())
		os.Exit(2)
	}
	if deleteVirtualWireAPI.StatusCode() == http.StatusOK {
		fmt.Printf("Successfully deleted virtualwire %s\n", deleteVirtualWireID)
	} else {
		fmt.Printf("\nError HTTP response code != 200 when deleting %s. Response: %v", deleteVirtualWireID, deleteVirtualWireAPI.RawResponse())
		os.Exit(3)
	}
}

func init() {
	deleteVirtualWireFlags := flag.NewFlagSet("virtualwire-delete", flag.ExitOnError)
	deleteVirtualWireFlags.StringVar(&deleteVirtualWireID, "id", "", "usage: -id virtualwire-XX")
	RegisterCliCommand("virtualwire-delete", deleteVirtualWireFlags, deleteVirtualWire)
}
