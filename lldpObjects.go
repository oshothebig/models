package models

type LLDPIntf struct {
	BaseObj
	IfIndex int32 `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1",DESCRIPTION: "IfIndex where lldp needs to be configured"`
	Enable  bool  `DESCRIPTION: "Enable/Disable lldp config"`
}

type LLDPIntfState struct {
	BaseObj
	IfIndex int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1",DESCRIPTION: "IfIndex where lldp needs to be configured"`
	Enable  bool   `DESCRIPTION: "Enable/Disable lldp config"`
	PeerMac string `DESCRIPTION: "Mac address of direct connection"`
	PortId  string `DESCRIPTION: "Port id of direct connection"`
	TTL     int32  `DESCRIPTION: "Validity of the peer information"`
}
