package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/edgeinterface"
	"net/http"
	"os"
)

func deleteEdgeInterface(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	deleteEdgeInterfaceAPI := edgeinterface.NewDelete(edgeid, index)
	err := client.Do(deleteEdgeInterfaceAPI)
	if err != nil {
		fmt.Println("error deleting interface: " + err.Error())
		os.Exit(2)
	}

	if deleteEdgeInterfaceAPI.StatusCode() == http.StatusNoContent {
		fmt.Printf("Edge interface %d successfully deleted\n", index)
	} else {
		fmt.Printf("Error deleting edge interface. StatusCode: %d\n", deleteEdgeInterfaceAPI.StatusCode())
		os.Exit(3)
	}

}

func init() {
	deleteEdgeInterfaceFlags := flag.NewFlagSet("edgeinterface-delete", flag.ExitOnError)
	deleteEdgeInterfaceFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id to attach interfaces to>")
	deleteEdgeInterfaceFlags.IntVar(&index, "index", 0, "usage: -index <index of the interface to delete")
	RegisterCliCommand("edgeinterface-delete", deleteEdgeInterfaceFlags, deleteEdgeInterface)
}
