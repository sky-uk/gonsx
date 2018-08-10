package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/firewallexclusion"
	"net/http"
	"os"
)

func readAllFirewallExclusion(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	readAllFirewallExclusion := firewallexclusion.NewGetAll()
	err := client.Do(readAllFirewallExclusion)
	if err != nil {
		fmt.Println("error retrieving a list of all firewall exclusions" + err.Error())
		os.Exit(1)
	}

	if readAllFirewallExclusion.StatusCode() == http.StatusOK {
		firewallExclusionShowAllResponse := readAllFirewallExclusion.GetResponse()
		rows := []map[string]interface{}{}
		headers := []string{"MOID", "Name"}

		for _, firewallExcludedMember := range firewallExclusionShowAllResponse.Members {
			row := map[string]interface{}{}
			row["MOID"] = firewallExcludedMember.MOID
			row["Name"] = firewallExcludedMember.Name
			rows = append(rows, row)
		}
		PrettyPrintMany(headers, rows)
	} else {
		fmt.Printf("\nError firewallexclusion-show-all HTTP return code != 200: %d", readAllFirewallExclusion.StatusCode())
		os.Exit(2)
	}
}

func init() {
	readAllFirewallExclusionFlags := flag.NewFlagSet("firewallexclusion-show-all", flag.ExitOnError)
	RegisterCliCommand("firewallexclusion-show-all", readAllFirewallExclusionFlags, readAllFirewallExclusion)
}
