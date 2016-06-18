package events

type PortKey struct {
	IntfRef string
}

type VlanKey struct {
	VlanId int32
}

const (
	PortOperStateUp   EventId = 1
	PortOperStateDown EventId = 2
	PortSpeedChange   EventId = 3
	VlanOperStateUp   EventId = 4
	VlanStateDown     EventId = 5
)

var AsicdEventKeyMap KeyMap = KeyMap{
	"Port": PortKey{},
	"Vlan": VlanKey{},
}
