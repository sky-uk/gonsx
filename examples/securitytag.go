package main

import (
	"fmt"
	"os"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/securitytags"
	"errors"
)

func getAllSecurityTags(nsxclient *gonsx.NSXClient) (*securitytags.SecurityTags, error){
	api := securitytags.NewGetAll()
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

func createSecurityTag(name, desc string, nsxclient *gonsx.NSXClient) (string, error){
	api := securitytags.NewCreate(name, desc)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}

	fmt.Println("Creating security tag with id ", api.GetResponse())
	return api.GetResponse(), nil

}

func deleteSecurityTag(ID string, nsxclient *gonsx.NSXClient)(error){
	api := securitytags.NewDelete(ID)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Deleting security tag with id" , ID)
	return nil

}

func RunSecurityTagExample(nsxManager, nsxUser, nsxPassword string, debug bool){
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
	} else{
		fmt.Println("Tag already exists")
	}

	getTags, err = getAllSecurityTags(nsxclient)
	if err != nil {
		fmt.Println("Failed to get tags")
		os.Exit(1)
	}

	if getTags.CheckByName("test") {
		ID := getTags.FilterByName("test").ObjectID
		err := deleteSecurityTag(ID, nsxclient)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}


}

