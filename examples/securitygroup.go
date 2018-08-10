package main

import (
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/securitygroup"
)

// RunSecurityGroupExample  Implementes Service example.
func RunSecurityGroupExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All Services.
	//
	// Create api object.
	fmt.Println("== Running Get All ==")
	getAllAPI := securitygroup.NewGetAll("globalroot-0")

	// make api call.
	err := nsxclient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllSecurityGroups := getAllAPI.GetResponse().SecurityGroups
		for _, secGroup := range AllSecurityGroups {
			fmt.Printf("objectId: %-20s name: %-20s\n", secGroup.ObjectID, secGroup.Name)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	//
	// Create single service.
	//
	fmt.Println("== Running Create new SecurityGroup with name 'OVP_sg_test' ==")

	dynamicCriteria := securitygroup.NewDynamicCriteria("OR", "VM.SECURITY_TAG", "ovp_test_app4", "contains")
	dynamicCriteriaList := []securitygroup.DynamicCriteria{*dynamicCriteria}

	dynamicSet := securitygroup.NewDynamicSet("OR", dynamicCriteriaList)
	dynamicSetList := []securitygroup.DynamicSet{*dynamicSet}

	dynamicMemberDefinition := securitygroup.NewDynamicMemberDefinition(dynamicSetList)

	createAPI := securitygroup.NewCreate("globalroot-0", "OVP_sg_test", dynamicMemberDefinition)
	err = nsxclient.Do(createAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if createAPI.StatusCode() == 201 {
		newSecurityGroupID := createAPI.ResponseObject()
		fmt.Println("SecurityGroup created successfully.")
		fmt.Println("objectId:", newSecurityGroupID)
	} else {
		fmt.Println("Failed to create the securitygroup!")
		fmt.Println(createAPI.ResponseObject())
	}

	//
	// Get Single SecurityGroup
	//
	// Get All (we need to make the getAllAPI call again to get a fresh list of securitygroups. )
	// check the status code and proceed accordingly.
	fmt.Println("== Running Get Single Security Group with name 'OVP_sg_test' ==")
	// make api call.
	err = nsxclient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if getAllAPI.StatusCode() == 200 {
		service := getAllAPI.GetResponse().FilterByName("OVP_sg_test")
		if service.ObjectID != "" {
			fmt.Println(service)
		} else {
			fmt.Println("Not found!")
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	//
	// Update SecurityGroup
	//
	securityGroupToUpdate := getAllAPI.GetResponse().FilterByName("OVP_sg_test")

	newDynamicCriteria := securitygroup.DynamicCriteria{
		Operator: "AND",
		Key:      "VM.NAME",
		Value:    "test_vm_name2",
		Criteria: "contains",
	}
	newDynamicCriteriaList := []securitygroup.DynamicCriteria{newDynamicCriteria}
	securityGroupToUpdate.AddDynamicMemberDefinitionSet("OR", newDynamicCriteriaList)

	updateAPI := securitygroup.NewUpdate(securityGroupToUpdate.ObjectID, securityGroupToUpdate)

	err = nsxclient.Do(updateAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if updateAPI.StatusCode() == 200 {
		fmt.Println("Security Group object updated successfully.")
		response := updateAPI.GetResponse()
		fmt.Println(response)
	} else {
		fmt.Println("Failed to update the security group!")
		fmt.Println("StatusCode:", updateAPI.StatusCode())
		fmt.Println("ResponseObject:", updateAPI.ResponseObject())
	}

	//
	// Delete single SecurityGroup with objectId
	//
	// Get All (note that we're re-utilizing the GetAll object from above here )
	// check the status code and proceed accordingly.
	fmt.Println("== Running Delete Single Security Group with name 'OVP_sg_test' ==")
	securityGroup := getAllAPI.GetResponse().FilterByName("OVP_sg_test")
	if securityGroup.ObjectID != "" {
		fmt.Println(securityGroup)
	} else {
		fmt.Println("Not found!")
	}

	deleteAPI := securitygroup.NewDelete(securityGroup.ObjectID)
	err = nsxclient.Do(deleteAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if deleteAPI.StatusCode() == 200 {
		fmt.Println(deleteAPI.ResponseObject())
		fmt.Println("Security group deleted successfully.")
	} else {
		fmt.Println("Failed to delete the security group.!")
		fmt.Println("StatusCode:", deleteAPI.StatusCode())
		fmt.Println("ResponseObject:", deleteAPI.ResponseObject())
	}

}
