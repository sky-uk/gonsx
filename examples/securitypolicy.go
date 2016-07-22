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
	// Get All SecurityPolicies.
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


	//
	// Create a new SecurityPolicy
	//
	createAPI := securitypolicy.NewCreate("ovp_test_security_policy", "50002",
		"long description", []string{"securitygroup-177", "securitygroup-197"})
	// * make sure above securitygroup names exist or you will be creating them
	// as part of terraform manifest and reference the security group names.
	// function expects a list of string, I'm create a list inplace in the function
	// call but this can be passed from terraform manifest.
	//
	// * precendene needs to be unique for security policies.  NSX doesn't allows
	// creating new policy with same precedence.
	//
	createErr := nsxclient.Do(createAPI)
	if createErr != nil {
		fmt.Println("Error:", createErr)
	}
	// check if the status code.
	if createAPI.StatusCode() == 201 {
		fmt.Println("SecurityPolicy created.")
	} else {
		fmt.Println("Status code:", createAPI.StatusCode())
		fmt.Println("Response: ", createAPI.ResponseObject())
	}

	//
	// Delete a SecurityPolicy
	//


}
