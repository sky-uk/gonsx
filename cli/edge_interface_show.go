package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/edgeinterface"
	"net/http"
	"os"
)

func getAllEdgeInterfaces(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	getAllEdgeInterfaceAPI := edgeinterface.NewGetAll(edgeid)
	err := client.Do(getAllEdgeInterfaceAPI)
	if err != nil {
		fmt.Println("error creating edge interfaces: " + err.Error())
		os.Exit(2)
	}

	if getAllEdgeInterfaceAPI.StatusCode() == http.StatusOK {
		interfaces := getAllEdgeInterfaceAPI.GetResponse()
		fmt.Println("------------ Edge Interfaces ------------")
		edgesAsXML, err := xml.MarshalIndent(interfaces, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(edgesAsXML))
	} else {
		fmt.Println("Error reading edge interfaces. StatusCode: ", getAllEdgeInterfaceAPI.StatusCode())
		os.Exit(3)
	}

}

func init() {
	getAllEdgeInterfaceFlags := flag.NewFlagSet("edgeinterface-show", flag.ExitOnError)
	getAllEdgeInterfaceFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id to attach interfaces to>")
	RegisterCliCommand("edgeinterface-show", getAllEdgeInterfaceFlags, getAllEdgeInterfaces)
}
