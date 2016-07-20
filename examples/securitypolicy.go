package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/securitypolicy"
)

// RunSecurityPolicyExample  Implementes Service example.
func RunSecurityPolicyExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)


	//
	// Get All Services.
	//
	// Create api object.
	fmt.Println("== Running Get All ==")
	getAllAPI := securitypolicy.NewGetAll()

	// make api call.
	err := nsxclient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllSecurityPolicies := getAllAPI.GetResponse().SecurityPolicies
		for _, secPolicy := range AllSecurityPolicies {
			fmt.Printf("objectId: %-20s name: %-20s\n", secPolicy.ObjectID, secPolicy.Name)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

}
