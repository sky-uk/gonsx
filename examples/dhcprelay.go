package main

import (
	"fmt"
	"os"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/dhcprelay"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("syntax error\nUsages: %s [NSX Manager Address] [Username] [Password] [INput file]\n\n", os.Args[0])
		os.Exit(1)
	}
	nsxManager := os.Args[1]
	nsxUser := os.Args[2]
	nsxPassword := os.Args[3]

	// Create NSXClient object.
	nsxclient := gonsx.NewNSXClient("https://"+nsxManager, nsxUser, nsxPassword, true, true)

	// Create DHCPRelay API call reference.
	getAllDhcpRelay := dhcprelay.NewGetAll("edge-50")

	// Actually making the api call with nsxclient.
	err := nsxclient.Do(getAllDhcpRelay)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(getAllDhcpRelay.GetResponse())
	}

}
