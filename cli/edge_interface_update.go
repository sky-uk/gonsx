package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/edgeinterface"
	"io/ioutil"
	"net/http"
	"os"
)

var edge edgeinterface.EdgeInterface

func updateEdgeInterface(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	if configFile == "" {
		fmt.Println("configFile must be set - usage: -configFile <a XML-encoded config file with interfaces configuration >")
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(dat, &edge)
	if err != nil {
		fmt.Println("Error decoding edge profile: ", err.Error())
		os.Exit(4)
	}

	updateEdgeInterfaceAPI := edgeinterface.NewUpdate(edgeid, index, edge)
	err = client.Do(updateEdgeInterfaceAPI)
	if err != nil {
		fmt.Println("error deleting interface: " + err.Error())
		os.Exit(2)
	}

	if updateEdgeInterfaceAPI.StatusCode() == http.StatusNoContent {
		fmt.Printf("Edge interface %d successfully updated\n", index)
	} else {
		fmt.Printf("Error updating edge interface. StatusCode: %d\n", updateEdgeInterfaceAPI.StatusCode())
		os.Exit(3)
	}

}

func init() {
	updateEdgeInterfaceFlags := flag.NewFlagSet("edgeinterface-update", flag.ExitOnError)
	updateEdgeInterfaceFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id to attach interfaces to>")
	updateEdgeInterfaceFlags.IntVar(&index, "index", 0, "usage: -index <index of the interface to update")
	updateEdgeInterfaceFlags.StringVar(&configFile, "configFile", "", "usage: -configFile <an XML file path with interface profile")
	RegisterCliCommand("edgeinterface-update", updateEdgeInterfaceFlags, updateEdgeInterface)
}
