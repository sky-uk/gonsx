package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
	"encoding/csv"
	"git.devops.int.ovp.bskyb.com/paas/gonsx/nsxclient"
)

func main() {
	if(len(os.Args) != 5) {
		fmt.Printf("syntax error\nUsages: %s [NSX Manager Address] [Username] [Password] [INput file]\n\n", os.Args[0])
		os.Exit(1)
	}

	nsxManager := os.Args[1]
	nsxUser := os.Args[2]
	nsxPassword := os.Args[3]

	nsx_client, err := nsxclient.NewNsxClient(nsxManager, nsxUser, nsxPassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inputFileName := os.Args[4]
	readFilePath, err := filepath.Abs(inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	csvfile, err := os.Open(readFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 4
	reader.Comment = '#'
	reader.Comma = ';'
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, chunk := range rawCSVdata {
		resp, err := nsx_client.CreateService(chunk[0], chunk[1], chunk[2], chunk[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		fmt.Println("HTTP Response is: " + resp.Status)
		if(resp.StatusCode != 201) {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Error in execurint query - got following error:")
			fmt.Println(string(b))
			os.Exit(1)

		}
		loc, err := resp.Location()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		l := strings.Split(loc.String(), "/")
		fmt.Println("Successfully created application Object ID: " + l[len(l)-1])
	}

}