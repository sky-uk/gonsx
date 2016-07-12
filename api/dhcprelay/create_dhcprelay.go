package dhcprelay


import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

type CreateDhcpRelayApi struct {
	*api.BaseApi
}

func NewCreate(dhcpIpAddress, edgeId string, relayAgentslist []RelayAgent) *CreateDhcpRelayApi {
	this := new(CreateDhcpRelayApi)
	requestPayload := new(DhcpRelay)
	requestPayload.RelayServer.IpAddress = dhcpIpAddress
	requestPayload.RelayAgents = relayAgentslist

	this.BaseApi = api.NewBaseApi(http.MethodPut, "/api/4.0/edges/" + edgeId +"/dhcp/config/relay", requestPayload, new(string))
	return this
}


func (this CreateDhcpRelayApi) GetResponse() string{
	return this.ResponseObject().(string)
}
