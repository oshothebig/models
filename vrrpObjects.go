package models

/*
 * This DS will be used while Created/Deleting Vrrp Intf Config
 */
type VrrpIntfConfig struct {
	BaseObj
	IfIndex                 int32  `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*"`
	VRID                    int32  `SNAPROUTE: "KEY"` // no default for VRID
	Priority                int32  // default value is 100
	VirtualIPv4Addr         string // No Default for Virtual IPv4 addr.. Can support one or more
	AdvertisementInterval   int32  // Default is 100 centiseconds which is 1 SEC
	PreemptMode             bool   // False to prohibit preemption. Default is True.
	AcceptMode              bool   // The default is False.
	VirtualRouterMACAddress string // MAC address used for the source MAC address in VRRP advertisements
}

type VrrpIntfState struct {
	BaseObj
	IfIndex                 int32 `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*"`
	VRID                    int32 `SNAPROUTE: "KEY"`
	IntfIpAddr              string
	Priority                int32
	VirtualIPv4Addr         string
	AdvertisementInterval   int32
	PreemptMode             bool
	VirtualRouterMACAddress string
	SkewTime                int32
	MasterDownInterval      int32
}
