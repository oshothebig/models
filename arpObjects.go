package models

/*
 * This DS will be used while Created/Deleting Arp Config
 */
type ArpConfig struct {
        BaseObj
        // placeholder to create a key
        ArpConfigKey string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "Arp config"`
        Timeout      int32  `DESCRIPTION: "Global Arp entry timeout value. Default value: 600 seconds, Minimum Possible Value: 300 seconds, Unit: second"`
}

type ArpEntry struct {
        BaseObj
        IpAddr         string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Neighbor's IP Address"`
        MacAddr        string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address"`
        Vlan           int32  `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to"`
        Intf           string `DESCRIPTION: "Router Interface to which neighbor is attached to"`
        ExpiryTimeLeft string `DESCRIPTION: "Time left before entry expires in case neighbor departs"`
}
