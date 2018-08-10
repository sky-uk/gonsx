package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/virtualwire"
	"net/http"
	"os"
)

var virtualWireSpec virtualwire.CreateSpec
var virtualWireCreateScopeID string

func createVirtualWire(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if virtualWireSpec.TenantID == "" {
		fmt.Println("tenantID must be set - usage: -tenantid a_tenant id")
		os.Exit(1)
	}
	if virtualWireCreateScopeID == "" {
		fmt.Println("scopeid must be set - usage: -scopeid vdnscope-1")
		os.Exit(1)
	}

	createVirtualWireAPI := virtualwire.NewCreate(virtualWireSpec, virtualWireCreateScopeID)
	err := client.Do(createVirtualWireAPI)
	if err != nil {
		fmt.Println("error creating new virtual wire: " + err.Error())
		os.Exit(2)
	}

	if createVirtualWireAPI.StatusCode() == http.StatusCreated {
		fmt.Println("Virtual wire ID " + createVirtualWireAPI.GetResponse() + " successfully created")
	} else {
		fmt.Println("error creating virtual wire. Received http response code != 201: " + createVirtualWireAPI.GetResponse())
		os.Exit(3)
	}

}

func init() {
	createVirtualWireFlags := flag.NewFlagSet("virtualwire-create", flag.ExitOnError)
	createVirtualWireFlags.StringVar(&virtualWireSpec.Name, "name", "", "usage: -name my_logical_switch")
	createVirtualWireFlags.StringVar(&virtualWireSpec.Description, "description", "", "usage: -description 'some description'")
	createVirtualWireFlags.StringVar(&virtualWireSpec.ControlPlaneMode, "controlplanemode", "", "usage: -controlplanemode 'UNICAST_MODE'")
	createVirtualWireFlags.StringVar(&virtualWireSpec.TenantID, "tenantid", "", "usage: -tenantid a_tenant_id")
	createVirtualWireFlags.StringVar(&virtualWireCreateScopeID, "scopeid", "", "usage: -scopeid vdnscope-1")
	RegisterCliCommand("virtualwire-create", createVirtualWireFlags, createVirtualWire)
}
