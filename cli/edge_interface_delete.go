package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/edgeinterface"
	"net/http"
	"os"
	"strconv"
)

var indexStr string

func deleteEdgeInterface(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	var index int
	index = 0
	var err error
	if indexStr != "" {
		index, err = strconv.Atoi(indexStr)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("An index has to be provided, usage: -index <index>")
		os.Exit(1)
	}

	deleteEdgeInterfaceAPI := edgeinterface.NewDelete(index, edgeid)
	err = client.Do(deleteEdgeInterfaceAPI)
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
	deleteEdgeInterfaceFlags.StringVar(&indexStr, "index", "", "usage: -index <index of the interface to delete")
	RegisterCliCommand("edgeinterface-delete", deleteEdgeInterfaceFlags, deleteEdgeInterface)
}
