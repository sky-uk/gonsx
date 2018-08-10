package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
	"os"
)

var virtualWireShowID string

func showVirtualWire(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if virtualWireShowID == "" {
		fmt.Println("Error virtualwire ID is required. Usage: -id virtualwire-XX")
		os.Exit(1)
	}

	readVirtualWire := virtualwire.NewGet(virtualWireShowID)
	err := client.Do(readVirtualWire)
	if err != nil {
		fmt.Printf("\nError reading virtual wire ID: %s. Error: %v", virtualWireShowID, err)
		os.Exit(2)
	}

	readVirtualWireResponse := readVirtualWire.GetResponse()

	if readVirtualWire.StatusCode() == http.StatusOK {
		row := map[string]interface{}{}
		row["VirtualWireID"] = readVirtualWireResponse.ObjectID
		row["Name"] = readVirtualWireResponse.Name
		row["TenantID"] = readVirtualWireResponse.TenantID
		row["ControlPlaneMode"] = readVirtualWireResponse.ControlPlaneMode
		row["Description"] = readVirtualWireResponse.Description
		PrettyPrintSingle(row)
	} else {
		fmt.Printf("\nError HTTP response code != 200. Response %v\n", readVirtualWireResponse)
		os.Exit(3)
	}
}

func init() {
	virtualWireShow := flag.NewFlagSet("virtualwire-show", flag.ExitOnError)
	virtualWireShow.StringVar(&virtualWireShowID, "id", "", "usage: -id virtualwire-XX")
	RegisterCliCommand("virtualwire-show", virtualWireShow, showVirtualWire)
}
