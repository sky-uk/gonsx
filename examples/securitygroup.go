package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/securitygroup"
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
	createAPI := securitygroup.NewCreate("globalroot-0", "OVP_sg_test", "OR", "OR", "VM.SECURITY_TAG", "ovp_test_app4", "contains")
	err = nsxclient.Do(createAPI)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if createAPI.StatusCode() == 201 {
		newSecurityGroupID := createAPI.ResponseObject()
		fmt.Println("Service created successfully.")
		fmt.Println("objectId:", newSecurityGroupID)
	} else {
		fmt.Println("Failed to created the service!")
		fmt.Println(createAPI.ResponseObject())
	}


	//
	// Get Single SecurityGroup
	//
	// Get All (note that we're re-utilizing the GetAll object from above here )
	// check the status code and proceed accordingly.
	fmt.Println("== Running Get Single Security Group with name 'OVP_sg_test' ==")
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




}
