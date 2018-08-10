package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/firewallexclusion"
	"net/http"
	"os"
)

var moid string

func createFirewallExclusion(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if moid == "" {
		fmt.Println("moid must be set - usage: -moid vm-1234")
		os.Exit(1)
	}
	createFirewallExclusion := firewallexclusion.NewCreate(moid)
	err := client.Do(createFirewallExclusion)
	if err != nil {
		fmt.Println("Error creating firewall exclusion" + err.Error())
		os.Exit(1)
	}

	if createFirewallExclusion.StatusCode() == http.StatusOK {
		fmt.Println("Sucessfully created firewall exclusion")
		os.Exit(2)
	} else {
		fmt.Printf("Error creating firewall exclusion, got HTTP %d\n", createFirewallExclusion.StatusCode())
		os.Exit(1)
	}

}

func init() {
	createFirewallExclusionFlags := flag.NewFlagSet("firewallexclusion-create", flag.ExitOnError)
	createFirewallExclusionFlags.StringVar(&moid, "moid", "", "usage: -moid vm-1234")
	RegisterCliCommand("firewallexclusion-create", createFirewallExclusionFlags, createFirewallExclusion)
}
