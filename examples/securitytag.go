package main

import (
	"errors"
	"fmt"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/securitytag"
	"os"
)

// getAllSecurityTags - gets all securitytags
func getAllSecurityTags(nsxclient *gonsx.NSXClient) (*securitytag.SecurityTags, error) {
	api := securitytag.NewGetAll()
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	if api.StatusCode() == 200 {
		return api.GetResponse(), nil
	}

	return nil, errors.New(string(api.RawResponse()))
}

// createSecurityTag - creates securitytags
func createSecurityTag(name, desc string, nsxclient *gonsx.NSXClient) (string, error) {
	api := securitytag.NewCreate(name, desc)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}

	fmt.Println("Creating security tag with id ", api.GetResponse())
	return api.GetResponse(), nil

}

// deleteSecurityTag - deletes securitytags
func deleteSecurityTag(ID string, nsxclient *gonsx.NSXClient) error {
	api := securitytag.NewDelete(ID)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Deleting security tag with id", ID)
	return nil

}

// updateSecurityTag - Updates the tag
func updateSecurityTag(ID,name,description string, nsxclient *gonsx.NSXClient) error {
	api := securitytag.NewUpdate(ID, name, description)
	err := nsxclient.Do(api)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Updating security tag with id", ID)
	return nil
}

// RunSecurityTagExample - runs securitytag example
func RunSecurityTagExample(nsxManager, nsxUser, nsxPassword string, debug bool) {
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	getTags, err := getAllSecurityTags(nsxclient)
	if err != nil {
		fmt.Println("Failed to get tags. error response:", err)
		os.Exit(1)
	}

	if !getTags.CheckByName("test") {
		_, err := createSecurityTag("test", "t", nsxclient)
		if err != nil {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Tag already exists")
	}
	getTags, err = getAllSecurityTags(nsxclient)
	if err != nil {
		fmt.Println("Failed to get tags")
		os.Exit(1)
	}

	if getTags.CheckByName("test") {
		ID := getTags.FilterByName("test").ObjectID
		fmt.Println("Trying to update Tag",ID)
		updateerr := updateSecurityTag(ID,"test2","testing the update function", nsxclient)
		if updateerr != nil {
			fmt.Println("Unable to update tag " , updateerr)

		}

	}
	getTags, err = getAllSecurityTags(nsxclient)
	if getTags.CheckByName("test2") {
		fmt.Println("tag updated")
	}
	if getTags.CheckByName("test2") {
		ID := getTags.FilterByName("test2").ObjectID
		err := deleteSecurityTag(ID, nsxclient)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
