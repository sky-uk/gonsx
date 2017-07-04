package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"os"
	"github.com/sky-uk/gonsx/api/distributedfirewall"

)


/*func getAllFirewallRules(contextID string) (*distributedfirewall.Rule, error){
	return nil
}*/


func RunDistributedFirewallExamples(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
	AllRules  :=  distributedfirewall.NewGetAll("globalroot-0")
	err := nsxclient.Do(AllRules)

	if err != nil {
		os.Exit(1)
	}
	if AllRules.StatusCode() == 200 {
		myrules := AllRules.GetResponse()
		var section distributedfirewall.Section
		fmt.Println("Sections for Layer:")
		fmt.Println("======================================================================================")
		for _, section =  range myrules.Layer3Sections.Sections {
			fmt.Println("===============================================================================")
			fmt.Println("Section ID : ",section.Id)
			fmt.Println("Section Name : ",section.Name)
			fmt.Println("Section Type : ",section.Type)
			fmt.Println("Rules for :",section.Name)
			var l3Rule distributedfirewall.Rule

			for _, l3Rule = range section.Rules{
				fmt.Println(l3Rule)
			}

			//fmt.Println(section.Rules)

		}
		fmt.Println("Sections and Rules for Layer 2")
		fmt.Println("======================================================================================")
		for _, section =  range myrules.Layer2Sections.Sections {
			fmt.Println("Section Name : ",section.Name)
			fmt.Println("Section Type : ",section.Type)
			//fmt.Println(section.Rules)

		}
		//fmt.Println(AllRules.RawResponse())
	} else {
		fmt.Println("could not find firewall rules")
	}



}

