package main

import (
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/ipset"
)

// RunIPSetExample  Implements Security Policy example.
func RunIPSetExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All SecurityPolicies.
	//
	// Create api object.
	fmt.Println("== Running Get All ==")
	getAllAPI := ipset.NewGetAll("globalroot-0")

	// make api call.
	err := nsxclient.Do(getAllAPI)
	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllIPSet := getAllAPI.GetResponse().IPSets
		for _, ipSet := range AllIPSet {
			fmt.Printf("objectId: %-20s name: %-20s\n", ipSet.ObjectID, ipSet.Name)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	//
	// Create a new IpSet
	//
	fmt.Println("== Running Create new IpSet ==")
	createAPI := ipset.NewCreate(
		"globalroot-0",
		&ipset.IPSet{Value: "10.50.0.0/8", Name: "CIDtestV", Description: "CIDR_fefENV_Description"},
	)

	createErr := nsxclient.Do(createAPI)
	if createErr != nil {
		fmt.Println("Error:", createErr)
	}
	// check if the status code.
	if createAPI.StatusCode() == 201 {
		fmt.Println("IpSet created.")
	} else {
		fmt.Println("Status code:", createAPI.StatusCode())
		fmt.Println("Response: ", createAPI.ResponseObject())
	}

	// Now finally call update security policy api call.

	getAPI := ipset.NewGet(createAPI.ResponseObject().(string))

	getErr := nsxclient.Do(getAPI)
	if getErr != nil {
		fmt.Println("Error:", getErr)
	}
	// check if the status code.
	if getAPI.StatusCode() == 201 {
		fmt.Println("IpSet created.")
	} else {
		fmt.Println("Status code:", getAPI.StatusCode())
		fmt.Println("Response: ", getAPI.ResponseObject().(*ipset.IPSet))
	}
	updateData := getAPI.ResponseObject().(*ipset.IPSet)
	fmt.Println(updateData)
	updateData.Description = "CIDR_ENV_DescriptionUPDATED"
	updateAPI := ipset.NewUpdate(createAPI.ResponseObject().(string), updateData)
	updateErr := nsxclient.Do(updateAPI)
	if updateErr != nil {
		fmt.Println("Update Error:", updateErr)
	}
	// check if the status code.
	if updateAPI.StatusCode() == 200 {
		fmt.Println("IpSet updated.")
	} else {
		fmt.Println("IpSet update  failure!!!")
		fmt.Println("Status code:", updateAPI.StatusCode())
		fmt.Println("Response: ", updateAPI.ResponseObject())
	}

	//Delete a IpSet

	fmt.Println("== Running Delete IpSet ==")

	// build the delete API call object.
	fmt.Println("Trying to delete: " + createAPI.ResponseObject().(string))
	deleteAPICall := ipset.NewDelete(createAPI.ResponseObject().(string))

	// make the call.
	deleteErr := nsxclient.Do(deleteAPICall)

	// check for errors.
	if deleteErr != nil {
		fmt.Println("Error:", deleteErr)
	}

	// check if the status code.
	if deleteAPICall.StatusCode() == 200 {
		fmt.Println("IpSet deleted.")
	} else {
		fmt.Println("Error Status code:", deleteAPICall.StatusCode())
	}

}
