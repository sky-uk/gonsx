package dhcprelay


import (
	"github.com/sky-uk/paas-gonsx/api"
	"net/http"
)

type CreateDhcpRelayApi struct {
	*api.BaseApi
}

func NewCreate(interfaceId, dhcpRelayIp string){
	
}

func (this CreateDhcpRelayApi) GetResponse() *DhcpRelay {
	return this.RequestObject().(*DhcpRelay)
}