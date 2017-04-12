[![Build Status](http://jenkins.paas.int.ovp.bskyb.com/buildStatus/icon?job=gonsx/build)](http://jenkins.paas.int.ovp.bskyb.com/job/gonsx/job/build/)
# gonsx client library

## Overview

This project is a NSXClient library for talking to NSX API.

### Features

| Feature                 | Create | Read  | Update  | Delete |
|-------------------------|--------|-------|---------|--------|
| DHCP Relay              |   N    |   Y   |    Y    |   Y    |
| Edge Interface          |   Y    |   Y   |    N    |   Y    |
| Security Group          |   Y    |   Y   |    Y    |   Y    |
| Security Policy         |   Y    |   Y   |    Y    |   Y    |
| Security Tag            |   Y    |   Y   |    N    |   Y    |
| Security Tag Attachment |   Y    |   Y   |    N    |   Y    |
| Service                 |   Y    |   Y   |    Y    |   Y    |
| Transport Zone          |   N    |   Y   |    N    |   N    |
| Virtual Wire            |   Y    |   Y   |    Y    |   Y    |

Implementation of CRUD in above features is partial in some cases.


### Security Policy features


| Feature                          | NSX API                                                          | GONSX Library |
|----------------------------------|------------------------------------------------------------------|---------------|
| End point                        | POST /2.0/services/policy/securitypolicy                         |       Y       |
| Status                           | GET /2.0/services/policy/securitypolicy/status                   |       N       |
| Alarms                           | GET /2.0/services/policy/securitypolicy/alarms/all               |       N       |
| Provider Firewall                | GET /2.0/services/policy/securitypolicy/serviceprovider/firewall |       N       |
| Provider Firewall                | PUT /2.0/services/policy/securitypolicy/serviceprovider/firewall |       N       |
| ID                               | GET /2.0/services/policy/securitypolicy/{ID}                     |       Y       |
| ID                               | PUT /2.0/services/policy/securitypolicy/{ID}                     |       Y       |
| ID                               | DELETE /2.0/services/policy/securitypolicy/{ID}                  |       Y       |
| Security Actions                 | GET /2.0/services/policy/securitypolicy/{ID}/securityactions     |       N       |
| Hierarchy                        | GET /2.0/services/policy/securitypolicy/hierarchy                |       N       |
| Hierarchy                        | POST /2.0/services/policy/securitypolicy/hierarchy               |       N       |
| Security Action VMs              | GET /2.0/services/policy/securityaction/category/virtualmachines |       N       |
| Security Action Security Groups  | GET /2.0/services/policy/securitygroup/{ID}/securityactions      |       N       |
| VM Security Actions              | GET /2.0/services/policy/virtualmachine/{ID}/securityactions     |       N       |
| Service Provider Firewall        | GET /2.0/services/policy/serviceprovider/firewall                |       N       |
| Security Group Security Policies | GET /2.0/services/policy/securitygroup/{ID}/securitypolicies     |       N       |


### Security Group Features

Apart from the basic CRUD functionality the Security Group supports the following

| Component               | Funcionality Name                                           | GONSX |
|-------------------------|-------------------------------------------------------------|-------|
| Security Group          | Add Multiple Dynamic Membership Criteria                    |   P   |
| Security Group          | Static Membership (always add these )                       |   N   |
| Security Group          | Static Exclusions (always exclude these )                   |   N   |
| Security Group          | Add More than one Dynamic Membership criteria               |   N   |
| Security Group          | Operators for more than one Dynamic Membership criteria     |   N   |

Criteria Details for dynamic membership

- OS Name
- Computer Name
- VM Name
- Security Tag - Implemented
- Entity :This could be
     - Another Security Group
     - Cluster
     - Logical Switch
     - Logical Port Group
     - vApp
     - Datacenter
     - IP Sets
     - Directory Group
     - MAC Sets
     - Security Tag
     - vNic
     - Virtual Machine
     - Resource Pool
     - Distributed Port Group



- Static Membership (always include )
     - Security Group
     - Cluster
     - Logical Switch
     - Logical Port Group
     - vApp
     - Datacenter
     - IP Sets
     - Directory Group
     - MAC Sets
     - Security Tag
     - vNic
     - Virtual Machine
     - Resource Pool
     - Distributed Port Group



 - Static Exclusions (never include)
      - Security Group
      - Cluster
      - Logical Switch
      - Logical Port Group
      - vApp
      - Datacenter
      - IP Sets
      - Directory Group
      - MAC Sets
      - Security Tag
      - vNic
      - Virtual Machine
      - Resource Pool
      - Distributed Port Group

NOTE:

Security Group - Add Multiple Dynamic Membership Criteria
This functionality is implemented partially, this is not accounted when creating , however we can do it while editing.

## Usage
### NSXClient

The NSXClient is the class used to send requests to the NSX host and pass through credentials.
 
Import the following files.
 
To create an NSX object run the following code, with the correct params. 

```
import (
    "github.com/sky-uk/gonsx"
)

nsxclient := gonsx.NewNSXClient(url, username, password, ignoreSSL, debug)
```
The params used:

* url: URL of NSX host
  
> E.G. https://nsxhost.com

* username: NSX username
* password: NSX password
* ignoreSSL: bool on whether to ignore ssl (default false)
* debug: bool on whether to debug output (default false)

The client is also used run the api calls once you have created the resource object.

```
nsxclient.Do(my_resource_obj)
```


### Virtual Wire(Logical Switch)

Virtual Wire resource. This resource will call the Virtual Wires api within NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/virtualwire
```

Create:

```
 api := virtualwire.NewCreate(name, desc, tennantID, scopeID)
 nsxclient.Do(api)
```

Read:
```
api := virtualwire.NewGetAll(scopeID)
nsxclient.Do(api)
resp := api.GetResponse().FilterByName(virtualWireName)
```

Update:
```
api := virtualwire.NewUpdate(name, desc, virtualwireID)
nsxclient.Do(api)
```

Delete:
```
api := virtualwire.NewDelete(virtualWireID)
nsxclient.Do(delete_api)
```


### Interface

Interface resource. This resource will call the interface api within NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/edgeinterface
```

Create:

```
 api := edgeinterface.NewCreate(edgeId, interfaceName, virtualWireId, gateway,
                                        		subnetMask, interfaceType, mtu)
 nsxclient.Do(api)
```

Read:
```
api := edgeinterface.NewGetAll(edgeID)
nsxclient.Do(api)
resp := api.GetResponse().FilterByName(interfaceName)
```

Update:
```
Not yet implemented
```

Delete:
```
api := edgeinterface.NewDelete(interfaceIndex, edgeId)
nsxclient.Do(delete_api)
```

### Dhcp Relay

DHCP resource. This resource will call the DHCP relay api within NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/dhcprelay
```

The Dhcp relay behaves differently in the API and as such it doesn't have a create and only an update and delete.
The delete function will remove the whole relay and all of its information. If you do not wish to do this and only
remove interfaces from the DHCP relay, then you must run an update instead. 

Read:
```
api := dhcp.NewGetAll(edgeId)
nsxclient.Do(api)
```

Update:
```
api := dhcprelay.NewUpdate(dhcpIpAddress, edgeId, relayAgentslist)
nsxclient.Do(api)
```

Delete:
```
api := dhcprelay.NewDelete(edgeId)
nsxclient.Do(api)
```

### Security Tag

Security tag resource. This resource will call the security tag api within NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/securitytag
```

Create:
```
api := securitytag.NewCreate(name, desc)
nsxclient.Do(api)
```

Read:
```
api := securitytag.NewGetAll()
nsxclient.Do(api)
```

Update:
```
Not yet implemented
```

Delete:
```
api := securitytag.NewDelete(securitytagID)
nsxclient.Do(api)
```

Detach:
```
api := securitytag.NewDetach(securityTagID, vmID)
nsxclient.Do(api)
```

Attach: 
```
api := securitytag.NewAttach(securityTagID, vmID)
nsxclient.Do(api)
```


### Service

Service resource. This resource will call the service api with NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/service
```

Create:
```
api := service.NewCreate(scopeID, name, desc, proto, ports)
nsxclient.Do(api)
```

Read:
```
api := service.NewGetAll(scopeID)
nsxclient.Do(api)
```

Update:
```
Not yet implemented
```

Delete:
```
api := service.NewDelete(serviceID)
nsxclient.Do(api)
```


### Security Group

Security Group resource. This resource will call the security group api with NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/securitygroup
```

Create:
```
api := securitygroup.NewCreate(scopeID, name, setOperator, criteriaOperator, criteriaKey, criteriaValue, criteria)
nsxclient.Do(api)
```

Read:
```
api := securitygroup.NewGetAll(scopeID)
nsxclient.Do(api)
```

Update:
```
Not yet implemented
```

Delete:
```
api := securitygroup.NewDelete(serviceID)
nsxclient.Do(api)
```


### Security Policy
Security Policy resource. This resource will call the security policy with NSX.
Import the following class:
```
github.com/sky-uk/gonsx/api/securitypolicy
```

Create:
```
api := securitypolicy.NewCreate(name, precendence, desc, securityGroupsIDs, actions)
nsxclient.Do(api)
```

Read:
```
api := securitypolicy.NewGetAll()
nsxclient.Do(api)
```

Update:
```
Not yet implemented
```

Delete:
```
api := securitypolicy.NewDelete(securityPolicyID, force)
nsxclient.Do(api)
```

AddOutboundFirewall:
```
securityPolicyToModify := getAllAPI.GetResponse().FilterByName(securityPolicyName)

// we will use a help function to add a firewall rule.
securityPolicyToModify.AddOutboundFirewallAction(
	firewallName,
	action,
	direction,
	secGroupObjectIDs,
	serviceIDs,
)

updateAPI := securitypolicy.NewUpdate(securityPolicyToModify.ObjectID, securityPolicyToModify)
nsxclient.Do(updateAPI)
```

RemoveFirewall:
```
securityPolicyToModify := getAllAPI.GetResponse().FilterByName(securityPolicyName)

securityPolicyToModify.RemoveFirewallActionByName(firewallName)

updateAPI := securitypolicy.NewUpdate(securityPolicyToModify.ObjectID, securityPolicyToModify)
nsxclient.Do(updateAPI)
```
