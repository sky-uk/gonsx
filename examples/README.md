# GoNSX client library examples

## Overview

Examples folder shows some examples how the gonsx client library can be 
utilized.

## Usage


```

go run examples/*.go  <nsxmanager address> <username> <password> <example name>

i.e.
go run examples/*.go  https://apnsx010 <username> <password> dhcprelay 


Get All Response:  &{{ relay} DhcpRelayServer ipAddress: []}
2016/07/13 14:01:35 <relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>16</vnicIndex><giAddress>10.152.165.1</giAddress></relayAgent></relayAgents></relay>
Updated DHCP Relay.

Get All Response:  &{{ relay} DhcpRelayServer ipAddress:10.152.160.10 [DhcpRelayAgent VnicIndex:16, GiAddress:10.152.165.1]}
2016/07/13 14:01:41 <relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>16</vnicIndex><giAddress>10.152.165.1</giAddress></relayAgent><relayAgent><vnicIndex>17</vnicIndex><giAddress>10.152.164.1</giAddress></relayAgent></relayAgents></relay>
Updated DHCP Relay.

Get All Response:  &{{ relay} DhcpRelayServer ipAddress:10.152.160.10 [DhcpRelayAgent VnicIndex:16, GiAddress:10.152.165.1 DhcpRelayAgent VnicIndex:17, GiAddress:10.152.164.1]}
There are other DHCP Relay agents, only removing single entry with update.
2016/07/13 14:01:49 <relay><relayServer><ipAddress>10.152.160.10</ipAddress></relayServer><relayAgents><relayAgent><vnicIndex>17</vnicIndex><giAddress>10.152.164.1</giAddress></relayAgent></relayAgents></relay>
Updated DHCP Relay.

Get All Response:  &{{ relay} DhcpRelayServer ipAddress:10.152.160.10 [DhcpRelayAgent VnicIndex:17, GiAddress:10.152.164.1]}
Last dhcp relay agent, removing the whole DHCP Relay.
DHCP Relay agent deleted.
Example not implemented.


```

