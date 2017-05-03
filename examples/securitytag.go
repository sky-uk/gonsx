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

func detatchSecurityTag(securityTagID string, vmID string, nsxclient *gonsx.NSXClient)error{
	api := securitytag.NewDetach(securityTagID,vmID)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Detaching security tag with id "+ securityTagID +" from "+ vmID)
	return nil
}

func  updateAttachedSecurityTags(vmID string, securityTagPayload *securitytag.SecurityTagAttachmentList, nsxclient *gonsx.NSXClient) (string,error){

	var vmsAttached *securitytag.SecurityTags
	vmsAttached,_ = getAllAttachedToVM(vmID,nsxclient)


	var tagsToRemove []string
	tagsToRemove = securityTagPayload.VerifyAttachments(vmsAttached)
	fmt.Print(tagsToRemove)

	for _, objectID := range tagsToRemove {
		detatchSecurityTag(objectID,vmID,nsxclient)
	}


	api := securitytag.NewUpdateAttachedTags(vmID, securityTagPayload)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	return api.GetResponse(), nil

}

func getAllAttachedToVM(vmID string, nsxclient *gonsx.NSXClient) (*securitytag.SecurityTags, error) {
	api := securitytag.NewGetAllAttachedToVM(vmID)
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
		err := deleteSecurityTag(ID, nsxclient)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}



}

func main()  {
	nsxclient := gonsx.NewNSXClient("https://apnsxa30", "SVC-OTT-PAAS-DEPLOY", "(U+m)y(3£T£R+R", true, true)
	vmID := "vm-426"

	securityTagAttachmentOne := securitytag.SecurityTagAttachment{ObjectID: "securitytag-127"}
	securityTagAttachmentTwo := securitytag.SecurityTagAttachment{ObjectID: "securitytag-128"}
	//securityTagAttachmentThree := securitytag.SecurityTagAttachment{ObjectID: "securitytag-129"}
	securityTagAttachmentList := new(securitytag.SecurityTagAttachmentList)
	securityTagAttachmentList.AddSecurityTagToAttachmentList(securityTagAttachmentOne)
	securityTagAttachmentList.AddSecurityTagToAttachmentList(securityTagAttachmentTwo)
	//securityTagAttachmentList.AddSecurityTagToAttachmentList(securityTagAttachmentThree)
/*
	var vmsAttached *securitytag.SecurityTags
	vmsAttached,_ = getAllAttachedToVM(vmID,nsxclient)


	var tagsToRemove []string
	tagsToRemove = securityTagAttachmentList.UpdateAttachments(vmsAttached)
	fmt.Print(tagsToRemove)

	for _, objectID := range tagsToRemove {
		detatchSecurityTag(objectID,vmID,nsxclient)
	}
*/

 	updateAttachedSecurityTags(vmID,securityTagAttachmentList,nsxclient)

}
