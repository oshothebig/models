package models

/*
 * This DS will be used while Created/Deleting Arp Config
 */
type ArpConfig struct {
	ConfigObj
	// placeholder to create a key
	ArpConfigKey string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "Arp config"`
	Timeout      int32  `DESCRIPTION: "Global Arp entry timeout value. Default value: 600 seconds, Minimum Possible Value: 300 seconds, Unit: second"`
}

type ArpEntryState struct {
	ConfigObj
	IpAddr         string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Neighbor's IP Address"`
	MacAddr        string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address"`
	Vlan           string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to"`
	Intf           string `DESCRIPTION: "Router Interface to which neighbor is attached to"`
	ExpiryTimeLeft string `DESCRIPTION: "Time left before entry expires in case neighbor departs"`
}

type ArpDeleteByIPv4Addr struct {
	ConfigObj
	IpAddr string `SNAPROUTE: "KEY", ACCESS:"x", MULTIPLICITY:"1", DESCRIPTION: "Neighbor's IP Address for which corresponding Arp entry needed to be re-learned"`
}

type ArpDeleteByIfName struct {
	ConfigObj
	IfName string `SNAPROUTE: "KEY", ACCESS:"x", MULTIPLICITY:"1", DESCRIPTION: "All the Arp learned on given IfName will be deleted and re-learned"`
}

type ArpLinuxEntryState struct {
	ConfigObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Neighbor's IP Address"`
	HWType  string `DESCRIPTION: "Hardware Type"`
	MacAddr string `DESCRIPTION: "MAC address of neighbor"`
	IfName  string `DESCRIPTION: "Interface name"`
}
