package main

import (
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/dhcprelay"
	"os"
)

func getAllDhcpRelays(edgeID string, nsxclient *gonsx.NSXClient) (*dhcprelay.DhcpRelay, error) {
	//
	// Get All DHCP Relay agents.
	//
	api := dhcprelay.NewGetAll(edgeID)
	// make the api call with nsxclient
	err := nsxclient.Do(api)
	// check if we err otherwise read response.
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Println("Get All Response: ", api.GetResponse())
	return api.GetResponse(), nil
}

// RunDhcpRelayExample ...Runs the DHCPRelay example.
func RunDhcpRelayExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	// Get All current DHCP Relays.
	CurrentDHCPRelay, err := getAllDhcpRelays("edge-5", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}

	//
	// Add New Relay Agent into existing list.
	//
	var createDhcp dhcprelay.DhcpRelay
	newRelayAgent := dhcprelay.RelayAgent{VnicIndex: "9", GiAddress: "10.72.232.200"}
	newRelayAgentsList := append(CurrentDHCPRelay.RelayAgents, newRelayAgent)
	createDhcp.RelayAgents = newRelayAgentsList

	createDhcp.RelayServer.IPAddress = []string{"10.152.160.10"}

	CreateAPI := dhcprelay.NewCreate("edge-5", createDhcp)

	err = nsxclient.Do(CreateAPI)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if CreateAPI.StatusCode() == 204 {
			fmt.Println("Created DHCP Relay.")
			fmt.Println(CreateAPI.GetResponse())
		} else {
			fmt.Println("Failed to update the DHCP relay 9")
			fmt.Println(CreateAPI.GetResponse())
		}
	}

	//
	// Add One More New Relay Agent into existing list.
	//

	// Get All current DHCP Relays.
	CurrentDHCPRelay, err = getAllDhcpRelays("edge-5", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}
	var updateDhcp dhcprelay.DhcpRelay
	newRelayAgent = dhcprelay.RelayAgent{VnicIndex: "1", GiAddress: "10.72.6.177"}
	newRelayAgentsList = append(CurrentDHCPRelay.RelayAgents, newRelayAgent)
	updateDhcp.RelayAgents = newRelayAgentsList
	updateDhcp.RelayServer.IPAddress = CurrentDHCPRelay.RelayServer.IPAddress

	updateAPI := dhcprelay.NewUpdate("edge-5", updateDhcp)

	err = nsxclient.Do(updateAPI)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if updateAPI.StatusCode() == 204 {
			fmt.Println("Updated DHCP Relay.")
			fmt.Println(updateAPI.GetResponse())
		} else {
			fmt.Println("Failed to update the DHCP relay")
			fmt.Println(updateAPI.GetResponse())
		}
	}

	//
	// Delete DHCP Relay Agent
	//
	// First get current dhcp relay agent list, we are using the objet from above Get All here.

	CurrentDHCPRelay, err = getAllDhcpRelays("edge-5", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}

	//  see if our vnic exists in relay agents list and this is the only one there.
	if CurrentDHCPRelay.CheckByVnicIndex("9") && (len(CurrentDHCPRelay.RelayAgents) == 1) {
		deleteAPI := dhcprelay.NewDelete("edge-5")
		err = nsxclient.Do(deleteAPI)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("DHCP Relay agent deleted.")
		}
	} else {
		var deleteDhcp dhcprelay.DhcpRelay
		// if we got more than one relay agents, then we have to call update after removing
		// the entry we want to remove.
		fmt.Println("There are other DHCP Relay agents, only removing single entry with update.")
		newRelayAgentsList := CurrentDHCPRelay.RemoveByVnicIndex("1").RelayAgents
		deleteDhcp.RelayServer.IPAddress = CurrentDHCPRelay.RelayServer.IPAddress
		deleteDhcp.RelayAgents = newRelayAgentsList

		updateAPI := dhcprelay.NewUpdate("edge-5", deleteDhcp)
		err = nsxclient.Do(updateAPI)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if updateAPI.StatusCode() == 204 {
				fmt.Println("Updated DHCP Relay.")
				fmt.Println(updateAPI.GetResponse())
			} else {
				fmt.Println("Failed to remove  the DHCP relay 1 ")
				fmt.Println(updateAPI.GetResponse())
			}
		}

	}

	//
	// Delete DHCP Relay Agent
	//
	// First get current dhcp relay agent list, we are using the objet from above Get All here.
	CurrentDHCPRelay, err = getAllDhcpRelays("edge-5", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}
	//  see if our vnic exists in relay agents list and this is the only one there.
	if CurrentDHCPRelay.CheckByVnicIndex("9") && (len(CurrentDHCPRelay.RelayAgents) == 1) {
		fmt.Println("Last dhcp relay agent, removing the whole DHCP Relay.")
		deleteAPI := dhcprelay.NewDelete("edge-5")
		err = nsxclient.Do(deleteAPI)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if updateAPI.StatusCode() == 204 {
				fmt.Println("Updated DHCP Relay.")
				fmt.Println(updateAPI.GetResponse())
			} else {
				fmt.Println("Failed to update the DHCP relay")
				fmt.Println(updateAPI.GetResponse())
			}
		}
	} else {
		var deleteDhcp2 dhcprelay.DhcpRelay
		// if we got more than one relay agents, then we have to call update after removing
		// the entry we want to remove.
		fmt.Println("There are other DHCP Relay agents, only removing single entry with update.")
		newRelayAgentsList := CurrentDHCPRelay.RemoveByVnicIndex("9").RelayAgents
		deleteDhcp2.RelayAgents = newRelayAgentsList
		deleteDhcp2.RelayServer.IPAddress = CurrentDHCPRelay.RelayServer.IPAddress
		updateAPI := dhcprelay.NewUpdate("edge-5", deleteDhcp2)
		err = nsxclient.Do(updateAPI)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if updateAPI.StatusCode() == 204 {
				fmt.Println("Updated DHCP Relay.")
				fmt.Println(updateAPI.GetResponse())
			} else {
				fmt.Println("Failed to update the DHCP relay")
				fmt.Println(updateAPI.GetResponse())
			}
		}

	}

}
