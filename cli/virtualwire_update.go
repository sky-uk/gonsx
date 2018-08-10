package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
	"os"
)

var virtualWireUpdateObject virtualwire.VirtualWire

func virtualWireUpdate(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if virtualWireUpdateObject.ObjectID == "" {
		fmt.Println("Virtual wire ID is required. Usage: -id virtualwire-XX")
		os.Exit(1)
	}
	updateVirtualWireAPI := virtualwire.NewUpdate(virtualWireUpdateObject)
	err := client.Do(updateVirtualWireAPI)
	if err != nil {
		fmt.Printf("\nError updating virtual wire ID %s: %v\n", virtualWireUpdateObject.ObjectID, err.Error())
		os.Exit(2)
	}
	if updateVirtualWireAPI.StatusCode() == http.StatusOK {
		fmt.Println("Successfully update virtual wire " + virtualWireUpdateObject.ObjectID)
	} else {
		fmt.Println("HTTP response code != 200")
		os.Exit(3)
	}
}

func init() {
	virtualWireUpdateFlags := flag.NewFlagSet("virtualwire-update", flag.ExitOnError)
	virtualWireUpdateFlags.StringVar(&virtualWireUpdateObject.ObjectID, "id", "", "usage: -id virtualwire-XX")
	virtualWireUpdateFlags.StringVar(&virtualWireUpdateObject.Name, "name", "", "usage: -name my_switch")
	virtualWireUpdateFlags.StringVar(&virtualWireUpdateObject.Description, "description", "", "usage: -description 'some_description'")
	virtualWireUpdateFlags.StringVar(&virtualWireUpdateObject.ControlPlaneMode, "controlplanemode", "", "usage: -controlplanemode UNICAST_MODE")
	RegisterCliCommand("virtualwire-update", virtualWireUpdateFlags, virtualWireUpdate)
}
