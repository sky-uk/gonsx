package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
	"os"
)

var virtualWireShowAllScopeID string

func readAllVirtualWires(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if virtualWireShowAllScopeID == "" {
		fmt.Println("Error scopeid must be specified. Usage -scopeid vdnscope-1")
		os.Exit(3)
	}

	readAllVirtualWires := virtualwire.NewGetAll(virtualWireShowAllScopeID)
	err := client.Do(readAllVirtualWires)
	if err != nil {
		fmt.Println("error retrieving a list of all virtual wires" + err.Error())
		os.Exit(1)
	}
	if readAllVirtualWires.StatusCode() == http.StatusOK {
		virtualWireShowAllResponse := readAllVirtualWires.GetResponse()
		rows := []map[string]interface{}{}
		headers := []string{"VirtualWireID", "Name", "TenantID", "ControlPlaneMode"}

		for _, virtualWire := range virtualWireShowAllResponse.DataPage.VirtualWires {
			row := map[string]interface{}{}
			row["VirtualWireID"] = virtualWire.ObjectID
			row["Name"] = virtualWire.Name
			row["TenantID"] = virtualWire.TenantID
			row["ControlPlaneMode"] = virtualWire.ControlPlaneMode
			rows = append(rows, row)
		}
		PrettyPrintMany(headers, rows)
	} else {
		fmt.Printf("\nError virtualwire-show-all HTTP return code != 200: %v", readAllVirtualWires.GetResponse())
		os.Exit(2)
	}
}

func init() {
	readAllVirtualWiresFlags := flag.NewFlagSet("virtualwire-show-all", flag.ExitOnError)
	readAllVirtualWiresFlags.StringVar(&virtualWireShowAllScopeID, "scopeid", "", "usage: -scopeid vdnscope-1")
	RegisterCliCommand("virtualwire-show-all", readAllVirtualWiresFlags, readAllVirtualWires)
}
