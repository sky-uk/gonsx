package main

import (
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/edge"
)

// RunEdgesExample Implementes edge example.
func RunEdgesExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	//
	// Create NSXClient object.
	//
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	//
	// Get All Edges
	//
	var edges []string
	api := edge.NewGetAll(100, 0)
	err := nsxclient.Do(api)
	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if api.StatusCode() == 200 {
		AllEdges := api.GetResponse().EdgePage
		for _, edge := range AllEdges.EdgeSummary {
			fmt.Printf("objectId: %-20s name: %-20s\n", edge.ObjectID, edge.Name)
			edges = append(edges, edge.ObjectID)
		}
	} else {
		fmt.Println("Status code:", api.StatusCode())
		fmt.Println("Response: ", api.ResponseObject())
	}

	//
	// Get Single Edges
	//
	apiEdge := edge.NewGet(edges[0])
	errEdge := nsxclient.Do(apiEdge)
	// check if there were any errors
	if errEdge != nil {
		fmt.Println("Error: ", errEdge)
	}

	// check the status code and proceed accordingly.
	if apiEdge.StatusCode() == 200 {
		edge := apiEdge.GetResponse()

		fmt.Printf("Found edge fqdn: %s", edge.FQDN)

	} else {
		fmt.Println("Status code:", apiEdge.StatusCode())
		fmt.Println("Response: ", apiEdge.ResponseObject())
	}
}
