package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"os"
	"github.com/sky-uk/gonsx/api/distributedfw/fwrules"
	//"github.com/sky-uk/gonsx/api/distributedfw"
)


func RunDistributedFirewallExamples(nsxManager, nsxUser, nsxPassword string, debug bool) {

	// Example to get al the rules of type LAYER3, inside section id 1110
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
	AllRules  :=  fwrules.NewGetAll("LAYER3","1110")
	err := nsxclient.Do(AllRules)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(AllRules.GetResponse())
	if AllRules.StatusCode() == 200 {
		myrules := AllRules.GetResponse()
		//var section distributedfw.Section
		fmt.Println(" ==== Sections for Layer 3 Section 1110====")
		fmt.Println(myrules)
	} else {
		fmt.Println("could not find firewall rules")
	}

	// Example to get a single Rule
	fmt.Println(" ==== Getting a single RULE ====")
	thisnsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
	thisRule  :=  fwrules.NewGetSingle("1163","LAYER3","1110")
	thiserr := thisnsxclient.Do(thisRule)

	if thiserr != nil {
		fmt.Println("Error getting a single rule")
		os.Exit(1)
	}

	if thisRule.StatusCode() == 200 {
		singleRule := thisRule.GetResponse()
		fmt.Println(singleRule)
	}



}

