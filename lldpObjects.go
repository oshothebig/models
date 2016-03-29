package models

type LLDPIntf struct {
	BaseObj
	IfIndex int32 `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1",DESCRIPTION: "IfIndex where lldp needs to be configured"`
	Enable  bool  `DESCRIPTION: "Enable/Disable lldp config"`
}
