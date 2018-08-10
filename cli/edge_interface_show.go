package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/edgeinterface"
	"net/http"
	"os"
)

var index int

func getEdgeInterface(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	getEdgeInterfaceAPI := edgeinterface.NewGet(edgeid, index)
	err := client.Do(getEdgeInterfaceAPI)
	if err != nil {
		fmt.Println("error creating edge interfaces: " + err.Error())
		os.Exit(2)
	}

	if getEdgeInterfaceAPI.StatusCode() == http.StatusOK {
		interfaces := getEdgeInterfaceAPI.GetResponse()

		fmt.Println("------------ Edge Interface ------------")
		edgesAsXML, err := xml.MarshalIndent(interfaces, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(edgesAsXML))
	} else {
		fmt.Println("Error reading edge interfaces. StatusCode: ", getEdgeInterfaceAPI.StatusCode())
		os.Exit(3)
	}

}

func init() {
	getEdgeInterfaceFlags := flag.NewFlagSet("edgeinterface-show", flag.ExitOnError)
	getEdgeInterfaceFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id to attach interfaces to>")
	getEdgeInterfaceFlags.IntVar(&index, "index", 0, "usage: -index <the interface index>")
	RegisterCliCommand("edgeinterface-show", getEdgeInterfaceFlags, getEdgeInterface)
}
