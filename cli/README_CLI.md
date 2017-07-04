# GONSX CLI

## Overview
Calling 'make' in the root of the repo will build the binary 'gonsx-cli'. The gonsx-cli will allow a user to interact with the NSX API.  
  


  
### Authentication  
  
NSX API credentials may be passed as options on the command line or set as shell environment variables.  
E.g.

```
$ ./gonsx-cli -server https://nsx-server-name -username nsx_user_name -password 'password' virtualwire-show-all -scopeid vdnscope-1
```
  
**OR**    
  
```
$ export NSX_SERVER=https://nsx-server-name
$ export NSX_USERNAME=nsx_user_name
$ export NSX_PASSWORD='password'
$ ./gonsx-cli virtualwire-show-all -scopeid vdnscope-1
```

### Help
```
$ ./gonsx-cli -help
```
  
**Help per option**  
```
$ ./gonsx-cli virtualwire-create -help
```
Replace virtualwire-create with the desired option.  


### Virtual Wire
The gonsx-cli binary allows for creating, readind, updating and deleting a virtual wire. 

**Attributes**  
|---------------------------------------------------------------|
| Attribute        | CLI Option        | Create | Read | Update |
|------------------|-------------------|--------|------|--------|
| Name             | -name             |    Y   |   Y  |    Y   |
| Description      | -description      |    Y   |   Y  |    Y   |
| ControlPlaneMode | -controlplanemode |    Y   |   Y  |    Y   |
| TenantID         | -tenantid         |    Y   |   Y  |    N   |
| ObjectID*        | -id               |    N   |   Y  |    N   |
| ScopeID**        | -scopeid          |    N   |   N  |    N   |
-----------------------------------------------------------------
*ObjectID is automatically set on creation and may only be read.  
**ScopeID is the scope in which to create a logical switch.  
  
#### Create a Virtual Wire
```
$ ./gonsx-cli virtualwire-create -name LogicalSwitch1 -description 'My First Logical Switch' -controlplanemode UNICAST_MODE -tenantid tenantID1 -scopeid vdnscope-1
Virtual wire ID virtualwire-1 successfully created
```
  
#### Read All Virtual Wires
```
$ ./gonsx-cli virtualwire-show-all -scopeid vdnscope-1  

|------------------------------------------------------------------------------------------------------|
| VirtualWireID  | Name                                        | TenantID            | ControlPaneMode |
|------------------------------------------------------------------------------------------------------|
| virtualwire-1  | LogicalSwitch1                              | tenantID1           | UNICAST_MODE    |
|------------------------------------------------------------------------------------------------------|
| virtualwire-2  | LogicalSwitch2                              | tenantID2           | HYBRID_MODE     |
|------------------------------------------------------------------------------------------------------|
```
  
#### Read a Virtual Wire
```
$ ./gonsx-cli virtualwire-show -id virtualwire-91

|-------------------------------------------|
| VirtualWireID   | virtualwire-1           |
|-------------------------------------------|
| Name            | LogicalSwitch1          |
|-------------------------------------------|
| TenantID        | tenantID1               |
|-------------------------------------------|
| ControlPaneMode | UNICAST_MODE            |
|-------------------------------------------|
| Description     | My First Logical Switch |
|-------------------------------------------|
```
  
#### Update a Virtual Wire
```
$ ./gonsx-cli virtualwire-update -id virtualwire-1 -controlplanemode UNICAST_MODE -description 'Updated my first logical switch' -name LogicalSwitch1
Successfully update virtual wire virtualwire-1
```
*Note: when updating a virtual wire all existing parameters need to be sent along with the attribute being changed.    
  
#### Delete a Virtual Wire
```
$ ./gonsx-cli virtualwire-delete -id virtualwire-1
Successfully deleted virtualwire virtualwire-1
```
  