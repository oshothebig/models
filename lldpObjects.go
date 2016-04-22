package models

type LLDPIntf struct {
	BaseObj
	IfIndex int32 `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1",DESCRIPTION: "IfIndex where lldp needs to be configured"`
	Enable  bool  `DESCRIPTION: "Enable/Disable lldp config"`
}

type LLDPIntfState struct {
	BaseObj
	IfIndex      int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1",DESCRIPTION: "IfIndex where lldp needs to be configured"`
	Enable       bool   `DESCRIPTION: "Enable/Disable lldp config"`
	LocalPort    string `DESCRIPTION: "Local interface"`
	PeerMac      string `DESCRIPTION: "Mac address of direct connection"`
	Port         string `DESCRIPTION: "Name of directtly connected pors"`
	HoldTime     string `DESCRIPTION: "Validity of the peer information"`
	Capabilities string `DESCRIPTION: "Capabilities of the peer port"`
}
