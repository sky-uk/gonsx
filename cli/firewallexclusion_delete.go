package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/firewallexclusion"
	"net/http"
	"os"
)

func deleteFirewallExclusion(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if moid == "" {
		fmt.Println("moid must be set - usage: -moid vm-1234")
		os.Exit(1)
	}
	deleteFirewallExclusion := firewallexclusion.NewDelete(moid)
	err := client.Do(deleteFirewallExclusion)
	if err != nil {
		fmt.Println("Error creating firewall exclusion" + err.Error())
		os.Exit(1)
	}

	if deleteFirewallExclusion.StatusCode() == http.StatusOK {
		fmt.Println("Sucessfully deleted firewall exclusion")
		os.Exit(2)
	} else {
		fmt.Printf("Error deleting firewall exclusion, got HTTP %d\n", deleteFirewallExclusion.StatusCode())
		os.Exit(1)
	}
}

func init() {
	deleteFirewallExclusionFlags := flag.NewFlagSet("firewallexclusion-show-all", flag.ExitOnError)
	deleteFirewallExclusionFlags.StringVar(&moid, "moid", "", "usage: -moid vm-1234")
	RegisterCliCommand("firewallexclusion-delete", deleteFirewallExclusionFlags, deleteFirewallExclusion)
}
