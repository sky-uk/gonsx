package gonsx

import (
	"github.com/tadaweb/gonsx/api/tzone"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClientGetAllTransportZones(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdnScopes><vdnScope><objectId>vdnscope-19</objectId><objectTypeName>VdnScope</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>VdnScope</typeName></type><name>S3_OVP_Routed_Network_Slough</name><description></description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><id>vdnscope-19</id><clusters><cluster><cluster><objectId>domain-c212</objectId><objectTypeName>ClusterComputeResource</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>405</revision><type><typeName>ClusterComputeResource</typeName></type><name>HDESX38</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></cluster></cluster></clusters><virtualWireCount>7</virtualWireCount><controlPlaneMode>UNICAST_MODE</controlPlaneMode></vdnScope></vdnScopes>`
	setup(http.StatusOK, xmlContent)
	defer server.Close()

	api := tzone.NewGetAll()
	nsxClient.Do(api)

	assert.Len(t, api.GetResponse().NetworkScopeList, 1)
	assert.Equal(t, "vdnscope-19", api.GetResponse().NetworkScopeList[0].ObjectID)
}

func TestClientGetAllTransportZonesFiltered(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdnScopes><vdnScope><objectId>vdnscope-19</objectId><objectTypeName>VdnScope</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>VdnScope</typeName></type><name>S3_OVP_Routed_Network_Slough</name><description></description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><id>vdnscope-19</id><clusters><cluster><cluster><objectId>domain-c212</objectId><objectTypeName>ClusterComputeResource</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>405</revision><type><typeName>ClusterComputeResource</typeName></type><name>HDESX38</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></cluster></cluster></clusters><virtualWireCount>7</virtualWireCount><controlPlaneMode>UNICAST_MODE</controlPlaneMode></vdnScope></vdnScopes>`
	setup(http.StatusOK, xmlContent)
	defer server.Close()

	api := tzone.NewGetAll()
	nsxClient.Do(api)
	actualTransportZone := api.GetResponse().FilterByName("S3_OVP_Routed_Network_Slough")

	assert.Equal(t, "vdnscope-19", actualTransportZone.ObjectID)
}

func TestClientGetTransportZone(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdnScope><objectId>vdnscope-19</objectId><objectTypeName>VdnScope</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>VdnScope</typeName></type><name>S3_OVP_Routed_Network_Slough</name><description></description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><id>vdnscope-19</id><clusters><cluster><cluster><objectId>domain-c212</objectId><objectTypeName>ClusterComputeResource</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>405</revision><type><typeName>ClusterComputeResource</typeName></type><name>HDESX38</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></cluster></cluster></clusters><virtualWireCount>7</virtualWireCount><controlPlaneMode>UNICAST_MODE</controlPlaneMode></vdnScope>`
	setup(http.StatusOK, xmlContent)
	defer server.Close()

	api := tzone.NewGet("vdnscope-19")
	nsxClient.Do(api)
	actualTransportZone := api.GetResponse()

	assert.Equal(t, "vdnscope-19", actualTransportZone.ObjectID)
}
