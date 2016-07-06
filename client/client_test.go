package client

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
	"git.devops.int.ovp.bskyb.com/paas/gonsx/client/api"
)

func TestHTTPResponse(t *testing.T) {
	user := "nsxUser"
	password := "nsxPass"
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<vdnScopes><vdnScope><objectId>vdnscope-19</objectId><objectTypeName>VdnScope</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>VdnScope</typeName></type><name>S3_OVP_Routed_Network_Slough</name><description></description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><id>vdnscope-19</id><clusters><cluster><cluster><objectId>domain-c212</objectId><objectTypeName>ClusterComputeResource</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>405</revision><type><typeName>ClusterComputeResource</typeName></type><name>HDESX38</name><scope><id>datacenter-21</id><objectTypeName>Datacenter</objectTypeName><name>S57 - Slough</name></scope><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision></cluster></cluster></clusters><virtualWireCount>7</virtualWireCount><controlPlaneMode>UNICAST_MODE</controlPlaneMode></vdnScope></vdnScopes>`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, xmlContent)
	}))
	defer server.Close()

	client := NewNSXClient(server.URL, user, password, true)

	api := api.NewGetTransportZone()

	client.get(api)
	assert.Equal(t, "vdnscope-19", api.GetResponse().NetworkScopeList[0].ObjectId)
}
