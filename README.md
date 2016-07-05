# gonsx client library

## Overview

This project is a NSXClient library for talking to NSX API.

## Testing

There is a test.go file and nsxservices.csv file.

```
go get -insecure git.devops.int.ovp.bskyb.com/paas/gonsx.git
```

the above seems to doesn't work with our gitlab so try following.

```
go to src into your GOPATH
mkdir -p git.devops.int.ovp.bskyb.com/paas/
cd into above directory and do a clone there manually.
```

Then you can run
```
go run test.go apnsx020 SVC-APP-OVP-DEPLOY <nsx password> ./nsx_services.csv
```