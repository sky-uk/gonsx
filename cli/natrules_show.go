package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/nat"
	"net/http"
	"os"
)

func getNatRulesAPI(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	getAllNatAPI := nat.NewGetAll(edgeid)
	err := client.Do(getAllNatAPI)
	if err != nil {
		fmt.Println("error getting edge: " + err.Error())
		os.Exit(2)
	}

	if getAllNatAPI.StatusCode() == http.StatusOK {
		interfaces := getAllNatAPI.GetResponse()

		fmt.Println("------------ Edge ------------")
		edgesAsXML, err := xml.MarshalIndent(interfaces, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(edgesAsXML))
	} else {
		fmt.Println("Error reading nat rules. StatusCode: ", getAllNatAPI.StatusCode())
		os.Exit(3)
	}

}

func init() {
	getNatRuleAPIFlags := flag.NewFlagSet("natrules-show", flag.ExitOnError)
	getNatRuleAPIFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id>")
	RegisterCliCommand("natrules-show", getNatRuleAPIFlags, getNatRulesAPI)
}
