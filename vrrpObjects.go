package models

/*
 * This DS will be used while Created/Deleting Vrrp Intf Config
 */
type VrrpIntfConfig struct {
	BaseObj
	IfIndex int32 `SNAPROUTE: "KEY"`
	// no default for VRID
	VRID int32
	// default value is 100
	Priority int32
	// No Default for IPv4 addr.. Can support one or more IPv4 addresses
	IPv4Addr []string
	// IPv6Addr... will add later when we decide to support IPv6

	// Default is 100 centiseconds which is 1 SEC
	AdvertisementInterval int32
	// False to prohibit preemption. Default is True.
	PreemptMode bool
	// The default is False.
	AcceptMode bool
	// MAC address used for the source MAC address in VRRP advertisements
	VirtualRouterMACAddress string
}
