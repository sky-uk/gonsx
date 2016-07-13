package service

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllServiceApi *GetAllServiceApi

func setupGetAll() {
	getAllServiceApi = NewGetAll("globalroot-0")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllServiceApi.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/services/application/globalroot-0", getAllServiceApi.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<list><application><objectId>application-5</objectId><objectTypeName>Application</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>Application</typeName></type><name>Service HTTP</name><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><element><applicationProtocol>TCP</applicationProtocol><value>8080</value></element></application><application><objectId>application-8</objectId><objectTypeName>Application</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>Application</typeName></type><name>HTTP 8181</name><scope><id>globalroot-0</id><objectTypeName>GlobalRoot</objectTypeName><name>Global</name></scope><element><applicationProtocol>TCP</applicationProtocol><value>8181</value></element></application></list>")

	xmlerr := xml.Unmarshal(xmlContent, getAllServiceApi.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllServiceApi.GetResponse().Applications, 2)
	assert.Equal(t, "8080", getAllServiceApi.GetResponse().Applications[0].Element[0].Value)
}
