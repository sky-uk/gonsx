package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/distributedfw/fwrules"
	"github.com/sky-uk/gonsx/api/distributedfw/sections"
	"os"

)


func CreateNewSource(name, value,sourceType string , valid bool ) fwrules.Source {
	var newSource fwrules.Source
	newSource.Name = name
	newSource.Value = value
	newSource.Type = sourceType
	newSource.IsValid = &valid
	return newSource
}

// RunDistributedFirewallExamples - Runs examples
func RunDistributedFirewallExamples(nsxManager, nsxUser, nsxPassword string, debug bool) {

		//example to create a new firewall rule
		var newrule fwrules.Rule

		var newDestination fwrules.Destination
		var newService fwrules.Service
		var newApplied fwrules.AppliedTo
	        var newSourceList  fwrules.SourceList
		var serviceList fwrules.SvcList
		creatensxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
		newrule.Name = "My Test Rule"
		newrule.Action = "ALLOW"
		newrule.RuleType = "LAYER3"
		newrule.Direction = "inout"
		newrule.SectionID = 1003
		newrule.PacketType = "any"
		newrule.Logged = "false"

		//newrule.Sources = append(newrule.Sources, newSource)
		newDestination.Name = "sandbox_private_sg"
		newDestination.Value = "securitygroup-714"
		newDestination.Type = "SecurityGroup"
		//newrule.Destinations = append(newrule.Destinations, newDestination)
		newService.Name = "SSH"
		newService.Value = "application-305"
		newService.Type = "Application"
		newService.DestinationPort = 80
		newService.Protocol = 6
		newSourceList.Source = append(newSourceList.Source, CreateNewSource("sandbox_private_sg","securitygroup-713","SecurityGroup",true))
		newSourceList.Excluded = "false"
		newrule.Sources = &newSourceList
		serviceList.Services = append(serviceList.Services, newService)
		newrule.Services = &serviceList
		newApplied.Name = "DISTRIBUTED_FIREWALL"
		newApplied.Value = "DISTRIBUTED_FIREWALL"
		newApplied.Type = "DISTRIBUTED_FIREWALL"
		newApplied.IsValid = true
		newrule.AppliedToList = append(newrule.AppliedToList, newApplied)
		newRuleAPI := fwrules.NewCreate(newrule)

		sectionTimestamp := sections.GetSectionTimestamp(newrule.SectionID, newrule.RuleType)
		sectsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
		sectErr := sectsxclient.Do(sectionTimestamp)
		if sectErr != nil {
			 fmt.Println("Error getting timestamp")
		}
		fmt.Println(sectionTimestamp.GetResponse().Timestamp)
		creatensxclient.SetHeader("If-Match",sectionTimestamp.GetResponse().Timestamp)
		errCreate := creatensxclient.Do(newRuleAPI)
		if errCreate != nil {
			fmt.Println("could not create")
		}


		return


	//Example to get all the sections
	sectionnsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
	AllSection := sections.NewGetAll()
	sectionsErr := sectionnsxclient.Do(AllSection)
	if sectionsErr != nil {
		os.Exit(1)
	}
	fmt.Println(AllSection.GetResponse())

	return

	// Example to get al the rules of type LAYER3, inside section id 1110
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)
	AllRules := fwrules.NewGetAll("LAYER3", "1110")
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
	thisRule := fwrules.NewGetSingle("1163", "LAYER3", "1110")
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
