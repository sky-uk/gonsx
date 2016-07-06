package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
)

var api *GetTransportZoneApi

func setup() {
	api = NewGetTransportZone()
}

func TestMethod(t *testing.T) {
	setup()
	assert.Equal(t, "GET", api.GetMethod())
}

func TestEndpoint(t *testing.T) {
	setup()
	assert.Equal(t, "/api/2.0/vdn/scopes", api.GetEndpoint())
}

func TestUnMarshalling(t *testing.T) {
	setup()
	xmlContent := []byte("<vdnScopes><vdnScope><objectId>vdnscope-19</objectId><objectTypeName>VdnScope</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>VdnScope</typeName></type><name>S3_OVP_Routed_Network_Slough</name><description></description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><id>vdnscope-19</id><clusters><cluster><cluster><objectId>domain-c212</objectId><objectTypeName>ClusterComputeResource</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>405</revision><type><typeName>ClusterComputeResource</typeName></type><name>HDESX38</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></cluster></cluster></clusters><virtualWireCount>7</virtualWireCount><controlPlaneMode>UNICAST_MODE</controlPlaneMode></vdnScope></vdnScopes>")
	xmlerr := xml.Unmarshal(xmlContent, api.GetResponseObject())
	assert.Nil(t, xmlerr)
	assert.Equal(t, "vdnscope-19", api.GetResponse().NetworkScopeList[0].ObjectId)
}

// TODO: add failure scenarios