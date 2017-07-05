package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("syntax error\nUsages: %s [https://nsxmanager_address] [username] [password] [Example Name] [true]\n\n", os.Args[0])
		os.Exit(1)
	}

	nsxManager := os.Args[1]
	nsxUser := os.Args[2]
	nsxPassword := os.Args[3]
	exampleName := os.Args[4]
	debug := false

	if len(os.Args) == 6 && os.Args[5] == "true" {
		debug = true
	}

	fmt.Println("running "+ exampleName +" with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
	switch exampleName {
	case "dhcprelay":
		RunDhcpRelayExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "edgeinterface":
		RunEdgeinterfaceExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "virtualwire":
		RunVirtualWireExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "service":
		RunServiceExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "securitygroup":
		RunSecurityGroupExample(nsxManager, nsxUser, nsxPassword, debug)
	case "securitytag":
		RunSecurityTagExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "securitypolicy":
		RunSecurityPolicyExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "ipset":
		RunIpSetExample(nsxManager, nsxUser, nsxPassword, debug)
		return

	}

	fmt.Println("Example not implemented.")

}
