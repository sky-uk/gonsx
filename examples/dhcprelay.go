package main

import (
	"fmt"
	"os"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/dhcprelay"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("syntax error\nUsages: %s [https://nsxmanager_address] [username] [password]\n\n", os.Args[0])
		os.Exit(1)
	}
	nsxManager := os.Args[1]
	nsxUser := os.Args[2]
	nsxPassword := os.Args[3]

	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, true)

	//
	// Get All DHCP Relay agents.
	//
	api := dhcprelay.NewGetAll("edge-50")
	// make the api call with nsxclient
	err := nsxclient.Do(api)
	// check if we err otherwise read response.
	var CurrentDHCPRelay *dhcprelay.DhcpRelay
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(api.GetResponse())
		CurrentDHCPRelay = api.GetResponse()
	}

	// Create New Relay Agent.
	dhcp_ip := CurrentDHCPRelay.RelayServer.IpAddress
	new_relay_agent := dhcprelay.RelayAgent{VnicIndex: "18", GiAddress: "10.152.163.1"}
	newRelayAgentsList := append(CurrentDHCPRelay.RelayAgents, new_relay_agent)

	update_api := dhcprelay.NewUpdate(dhcp_ip, "edge-50", newRelayAgentsList)

	err = nsxclient.Do(update_api)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(update_api.GetResponse())
	}



}
