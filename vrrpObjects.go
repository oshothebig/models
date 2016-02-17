package models

/*
 * This DS will be used while Created/Deleting Vrrp Intf Config
 */
type VrrpIntfConfig struct {
	BaseObj
	IfIndex                 int32  `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*"`
	VRID                    int32  // no default for VRID
	Priority                int32  // default value is 100
	IPv4Addr                string // No Default for IPv4 addr.. Can support one or more IPv4 addresses
	AdvertisementInterval   int32  // Default is 100 centiseconds which is 1 SEC
	PreemptMode             bool   // False to prohibit preemption. Default is True.
	AcceptMode              bool   // The default is False.
	VirtualRouterMACAddress string // MAC address used for the source MAC address in VRRP advertisements
}
