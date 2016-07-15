package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/dhcprelay"
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
	CurrentDHCPRelay, err := getAllDhcpRelays("edge-50", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}

	//
	// Add New Relay Agent into existing list.
	//
	newRelayAgent:= dhcprelay.RelayAgent{VnicIndex: "16", GiAddress: "10.152.165.1"}
	newRelayAgentsList := append(CurrentDHCPRelay.RelayAgents, newRelayAgent)

	updateAPI:= dhcprelay.NewUpdate("10.152.160.10", "edge-50", newRelayAgentsList)

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
	// Add One More New Relay Agent into existing list.
	//

	// Get All current DHCP Relays.
	CurrentDHCPRelay, err = getAllDhcpRelays("edge-50", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}

	newRelayAgent = dhcprelay.RelayAgent{VnicIndex: "17", GiAddress: "10.152.164.1"}
	newRelayAgentsList = append(CurrentDHCPRelay.RelayAgents, newRelayAgent)

	updateAPI = dhcprelay.NewUpdate("10.152.160.10", "edge-50", newRelayAgentsList)

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
	CurrentDHCPRelay, err = getAllDhcpRelays("edge-50", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}

	//  see if our vnic exists in relay agents list and this is the only one there.
	if CurrentDHCPRelay.CheckByVnicIndex("16") && (len(CurrentDHCPRelay.RelayAgents) == 1) {
		deleteAPI := dhcprelay.NewDelete("edge-50")
		err = nsxclient.Do(deleteAPI)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("DHCP Relay agent deleted.")
		}
	} else {
		// if we got more than one relay agents, then we have to call update after removing
		// the entry we want to remove.
		fmt.Println("There are other DHCP Relay agents, only removing single entry with update.")
		newRelayAgentsList := CurrentDHCPRelay.RemoveByVnicIndex("16").RelayAgents

		updateAPI := dhcprelay.NewUpdate("10.152.160.10", "edge-50", newRelayAgentsList)
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

	//
	// Delete DHCP Relay Agent
	//
	// First get current dhcp relay agent list, we are using the objet from above Get All here.
	CurrentDHCPRelay, err = getAllDhcpRelays("edge-50", nsxclient)
	if err != nil {
		fmt.Println("Failed to get all DHCP relays.")
		os.Exit(1)
	}
	//  see if our vnic exists in relay agents list and this is the only one there.
	if CurrentDHCPRelay.CheckByVnicIndex("17") && (len(CurrentDHCPRelay.RelayAgents) == 1) {
		fmt.Println("Last dhcp relay agent, removing the whole DHCP Relay.")
		deleteAPI := dhcprelay.NewDelete("edge-50")
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
		// if we got more than one relay agents, then we have to call update after removing
		// the entry we want to remove.
		fmt.Println("There are other DHCP Relay agents, only removing single entry with update.")
		newRelayAgentsList := CurrentDHCPRelay.RemoveByVnicIndex("16").RelayAgents

		updateAPI := dhcprelay.NewUpdate("10.152.160.10", "edge-50", newRelayAgentsList)
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
