package virtualwire

func (v VirtualWires) FilterByName(name string) *VirtualWire {
	var virtualWireFound VirtualWire
	for _, virtualWire := range v.DataPage.VirtualWires {
		if(virtualWire.Name == name) {
			virtualWireFound = virtualWire
			break
		}
	}
	return &virtualWireFound
}
