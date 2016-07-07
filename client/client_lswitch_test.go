package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api/lswitch"
)

func TestClientGetAllLogicalSwitches(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdsContexts><vdsContext><switch><objectId>dvs-61</objectId><objectTypeName>VmwareDistributedVirtualSwitch</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>420</revision><type><typeName>VmwareDistributedVirtualSwitch</typeName></type><name>vds-slu-d-ott-2x10-03</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></switch><mtu>1600</mtu><teaming>LOADBALANCE_SRCID</teaming><uplinkPortName>dvUplink1</uplinkPortName><uplinkPortName>dvUplink2</uplinkPortName><promiscuousMode>false</promiscuousMode></vdsContext></vdsContexts>`
	setup(200, xmlContent)
	defer server.Close()

	api := lswitch.NewGetAll()
	nsxClient.Do(api)

	assert.Len(t, api.GetResponse().VdsContextList, 1)
	assert.Equal(t, "dvs-61", api.GetResponse().VdsContextList[0].Switch.ObjectId)
}

func TestClientGetAllLogicalSwitchesFiltered(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdsContexts><vdsContext><switch><objectId>dvs-61</objectId><objectTypeName>VmwareDistributedVirtualSwitch</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>420</revision><type><typeName>VmwareDistributedVirtualSwitch</typeName></type><name>vds-slu-d-ott-2x10-03</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></switch><mtu>1600</mtu><teaming>LOADBALANCE_SRCID</teaming><uplinkPortName>dvUplink1</uplinkPortName><uplinkPortName>dvUplink2</uplinkPortName><promiscuousMode>false</promiscuousMode></vdsContext></vdsContexts>`
	setup(200, xmlContent)
	defer server.Close()

	api := lswitch.NewGetAll()
	nsxClient.Do(api)
	actualLogicalSwitch := api.GetResponse().FilterByName("vds-slu-d-ott-2x10-03")

	assert.Equal(t, "dvs-61", actualLogicalSwitch.Switch.ObjectId)
}

func TestClientGetLogicalSwitch(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdsContext><switch><objectId>dvs-61</objectId><objectTypeName>VmwareDistributedVirtualSwitch</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>420</revision><type><typeName>VmwareDistributedVirtualSwitch</typeName></type><name>vds-slu-d-ott-2x10-03</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></switch><mtu>1600</mtu><teaming>LOADBALANCE_SRCID</teaming><uplinkPortName>dvUplink1</uplinkPortName><uplinkPortName>dvUplink2</uplinkPortName><promiscuousMode>false</promiscuousMode></vdsContext>`
	setup(200, xmlContent)
	defer server.Close()

	api := lswitch.NewGet("dvs-61")
	nsxClient.Do(api)
	actualLogicalSwitch := api.GetResponse().Switch

	assert.Equal(t, "dvs-61", actualLogicalSwitch.ObjectId)
}

func TestClientCreateLogicalSwitch(t *testing.T) {
	// TODO: add actual response
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>`
	setup(200, xmlContent)
	defer server.Close()

	api := lswitch.NewCreate()
	nsxClient.Do(api)

	// TODO: add assertions
}
