[![Build Status](http://jenkins.paas.int.ovp.bskyb.com/buildStatus/icon?job=gonsx/build)](http://jenkins.paas.int.ovp.bskyb.com/job/gonsx/job/build/)
# gonsx client library

## Overview

This project is a NSXClient library for talking to NSX API.

## Usage
### NSXClient

The NSXClient is the class used to send requests to the NSX host and pass through credentials.
 
To create an NSX object run the following code, with the correct params. 

```
nsxclient := client.NewNSXClient(url, username, password, ignoreSSL, debug)
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
git.devops.int.ovp.bskyb.com/paas/gonsx/client/api/virtualwire
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
Not yet implemented
```

Delete:
```
api := virtualwire.NewDelete(virtualWireID)
nsxclient.Do(delete_api)
```
