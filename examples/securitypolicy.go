package main

import (
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/securitypolicy"
)

// RunSecurityPolicyExample  Implements Security Policy example.
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
	fmt.Println("== Running Create new SecurityPolicy with name 'ovp_test_security_policy' ==")
	createAPI := securitypolicy.NewCreate(
		"ovp_test_security_policy",
		"50002",
		"long description",
		[]string{"securitygroup-304", "securitygroup-305"},
		[]securitypolicy.Action{},
	)
	//
	// * make sure above securitygroup names exist or you will be creating them
	// as part of terraform manifest and reference the security group names.
	// function expects a list of string, I'm create a list inplace in the function
	// call but this can be passed from terraform manifest.
	//
	// * precendence needs to be unique for security policies.  NSX doesn't allows
	//   creating new policy with same precedence.
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
	// Add Firewall Rule.
	//
	//
	fmt.Println("== Running Add Firewall Rule to new SecurityPolicy with name 'ovp_test_security_policy' ==")
	// Refresh the response of getAllAPI because we just created a new security policy which won't be
	// there in the getAllAPI response which we have from earlier request.
	err = nsxclient.Do(getAllAPI)
	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// now search for the name and get the objectID of security policy we want to modify.
	securityPolicyToModify := getAllAPI.GetResponse().FilterByName("ovp_test_security_policy")

	// we will use a help function to add a firewall rule.
	securityPolicyToModify.AddOutboundFirewallAction(
		"DummyAllowOutboundRule",
		"allow",
		"outbound",
		[]string{"securitygroup-304"},
		[]string{"application-212", "application-66"},
	)

	securityPolicyToModify.AddInboundFirewallAction(
		"DummyAllowInbouncRule",
		"allow",
		"inbound",
		[]string{"application-212", "application-66"},
	)

	// Now finally call update security policy api call.
	updateAPI := securitypolicy.NewUpdate(securityPolicyToModify.ObjectID, securityPolicyToModify)
	updateErr := nsxclient.Do(updateAPI)
	if updateErr != nil {
		fmt.Println("Update Error:", updateErr)
	}
	// check if the status code.
	if updateAPI.StatusCode() == 200 {
		fmt.Println("SecurityPolicy updated.")
	} else {
		fmt.Println("SecurityPolicy update  failure!!!")
		fmt.Println("Status code:", updateAPI.StatusCode())
		fmt.Println("Response: ", updateAPI.ResponseObject())
	}

	//
	// Delete Firewall Rule.
	//
	fmt.Println("== Running Remove Firewall Rule to new SecurityPolicy with name 'ovp_test_security_policy' ==")
	// Refresh the response of getAllAPI because we just created a new security policy which won't be
	// there in the getAllAPI response which we have from earlier request.
	getAllAPI2 := securitypolicy.NewGetAll()
	err = nsxclient.Do(getAllAPI2)
	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// now search for the name and get the objectID of security policy we want to modify.
	securityPolicyToModify2 := getAllAPI2.GetResponse().FilterByName("ovp_test_security_policy")

	// we will use a help function to add a firewall rule.
	securityPolicyToModify2.RemoveFirewallActionByName("DummyRule")

	// Now finally call update security policy api call.
	updateAPI2 := securitypolicy.NewUpdate(securityPolicyToModify2.ObjectID, securityPolicyToModify2)
	updateErr = nsxclient.Do(updateAPI2)
	if updateErr != nil {
		fmt.Println("Update Error:", updateErr)
	}
	// check if the status code.
	if updateAPI2.StatusCode() == 200 {
		fmt.Println("SecurityPolicy updated.")
	} else {
		fmt.Println("SecurityPolicy update failure!!!")
		fmt.Println("Status code:", updateAPI2.StatusCode())
		fmt.Println("Response: ", updateAPI2.ResponseObject())
	}

	//Delete a SecurityPolicy

	fmt.Println("== Running Delete SecurityPolicy with name 'ovp_test_security_policy' ==")

	// Refresh the response of getAllAPI because we just created a new security policy which won't be
	// there in the getAllAPI response which we have from earlier request.
	err = nsxclient.Do(getAllAPI)
	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// now search for the name and get the objectID of security policy we want to delete.
	securityPolicyToDelete := getAllAPI.GetResponse().FilterByName("ovp_test_security_policy")
	fmt.Println(" securityPolicy ObjectID:", securityPolicyToDelete.ObjectID)
	// build the delete API call object.
	deleteAPICall := securitypolicy.NewDelete(securityPolicyToDelete.ObjectID, false)

	// make the call.
	deleteErr := nsxclient.Do(deleteAPICall)

	// check for errors.
	if deleteErr != nil {
		fmt.Println("Error:", deleteErr)
	}

	// check if the status code.
	if deleteAPICall.StatusCode() == 204 {
		fmt.Println("SecurityPolicy deleted.")
	} else {
		fmt.Println("Status code:", deleteAPICall.StatusCode())
		fmt.Println("Response: ", deleteAPICall.ResponseObject())
	}
	// END

}
