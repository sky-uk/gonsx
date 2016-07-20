package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("syntax error\nUsages: %s [https://nsxmanager_address] [username] [password] [Example Name]\n\n", os.Args[0])
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

	switch exampleName {
	case "dhcprelay":
		fmt.Println("running dhcprelay with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunDhcpRelayExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "edgeinterface":
		fmt.Println("running edge interface with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunEdgeinterfaceExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "virtualwire":
		fmt.Println("running virtualwire example with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunVirtualWireExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "securitytag":
		fmt.Println("running securitytag example with:", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunSecurityTagExample(nsxManager, nsxUser, nsxPassword, debug)
                return
	case "service":
		fmt.Println("running service example with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunServiceExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "securitygroup":
		fmt.Println("running service example with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunSecurityGroupExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	case "securitypolicy":
		fmt.Println("running service example with: ", nsxManager, nsxUser, nsxPassword, exampleName, debug)
		RunSecurityPolicyExample(nsxManager, nsxUser, nsxPassword, debug)
		return
	}


	fmt.Println("Example not implemented.")

}
