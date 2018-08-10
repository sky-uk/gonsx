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

var edgeInterfaces edgeinterface.EdgeInterfaces
var edgeid string
var configFile string

func createEdgeInterface(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

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

	err = xml.Unmarshal(dat, &edgeInterfaces)
	if err != nil {
		panic(err)
	}

	api := edgeinterface.NewCreate(&edgeInterfaces, edgeid)
	err = client.Do(api)
	if err != nil {
		fmt.Println("Error creating edge interfaces: " + err.Error())
		os.Exit(2)
	}

	if api.StatusCode() == http.StatusOK {
		fmt.Println("Edge interface successfully created")
		fmt.Println("-------------- Response -----------------")
		fmt.Printf("%s\n", api.RawResponse())

	} else {
		fmt.Println("Error creating edge interface. StatusCode: ", api.StatusCode())
		os.Exit(3)
	}

}

func init() {
	createEdgeInterfaceFlags := flag.NewFlagSet("edgeinterface-create", flag.ExitOnError)
	createEdgeInterfaceFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id to attach interfaces to>")
	createEdgeInterfaceFlags.StringVar(&configFile, "configFile", "", "usage: -configFile <an XML file path with interfaces configuration")
	RegisterCliCommand("edgeinterface-create", createEdgeInterfaceFlags, createEdgeInterface)
}
