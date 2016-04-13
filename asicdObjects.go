package models

type Vlan struct {
	BaseObj
	VlanId           int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	IfIndexList      string `DESCRIPTION: "List of system assigned interface id's for tagged ports on this vlan"`
	UntagIfIndexList string `DESCRIPTION: "List of system assigned interface id's for untagged ports on this vlan"`
}

type VlanState struct {
	BaseObj
	VlanId    int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	VlanName  string `DESCRIPTION: "System assigned vlan name"`
	OperState string `DESCRIPTION: "Operational state of vlan interface"`
	IfIndex   int32  `DESCRIPTION: "System assigned interface id for this vlan interface"`
}

type IPv4Intf struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface"`
	IfIndex int32  `DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
}

type IPv4IntfState struct {
	BaseObj
	IpAddr            string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
	OperState         string `DESCRIPTION: "Operational state of this IP interface"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	L2IntfType        string `DESCRIPTION: "Type of L2 interface on which IP has been configured (Port/Lag/Vlan)"`
	L2IntfId          int32  `DESCRIPTION: "Id of the L2 interface. Port number/lag id/vlan id."`
}

type Port struct {
	BaseObj
	PortNum     int32  `SNAPROUTE: "KEY", ACCESS:"rw", DESCRIPTION: "Front panel port number"`
	Description string `DESCRIPTION: "User provided string description", DEFAULT: "FP Port"`
	PhyIntfType string `DESCRIPTION: "Type of internal phy interface"`
	AdminState  string `DESCRIPTION: "Administrative state of this port"`
	MacAddr     string `DESCRIPTION: "Mac address associated with this port"`
	Speed       int32  `DESCRIPTION: "Port speed in Mbps"`
	Duplex      string `DESCRIPTION: "Duplex setting for this port"`
	Autoneg     string `DESCRIPTION: "Autonegotiation setting for this port"`
	MediaType   string `DESCRIPTION: "Type of media inserted into this port"`
	Mtu         int32  `DESCRIPTION: "Maximum transmission unit size for this port"`
}

type PortState struct {
	BaseObj
	PortNum           int32  `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Front panel port number"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this port"`
	Name              string `DESCRIPTION: "System assigned vlan name"`
	OperState         string `DESCRIPTION: "Operational state of front panel port"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	IfInOctets        int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts     int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards      int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors        int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets       int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts    int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards     int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors       int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
	ErrDisableReason  string `DESCRIPTION: "Reason explaining why port has been disabled by protocol code"`
}

type MacTableEntryState struct {
	BaseObj
	MacAddr string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "MAC Address"`
	VlanId  int32  `DESCRIPTION: "Vlan id corresponding to which mac was learned", DEFAULT:0`
	Port    int32  `DESCRIPTION: "Port number on which mac was learned", DEFAULT:0`
}

type IPv4RouteHwState struct {
	BaseObj
	DestinationNw    string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route in CIDR format"`
	NextHopIps       string `DESCRIPTION: "next hop ip list for the route"`
	RouteCreatedTime string `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime string `DESCRIPTION :"Time when the route was last updated"`
}

type ArpEntryHwState struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", QPARAM: "optional" ,DESCRIPTION: "Neighbor's IP Address"`
	MacAddr string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address", QPARAM: "optional" `
	Vlan    string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to", QPARAM: "optional" `
	Port    string `DESCRIPTION: "Router Interface to which neighbor is attached to", QPARAM: "optional" `
}

type LogicalIntf struct {
	BaseObj
	Name string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Name of logical interface"`
	Type string `DESCRIPTION: "Type of logical interface (e.x. loopback)"`
}

type LogicalIntfState struct {
	BaseObj
	Name              string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Name of logical interface"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this logical interface"`
	SrcMac            string `DESCRIPTION: "Source Mac assigned to the interface"`
	OperState         string `DESCRIPTION: "Operational state of logical interface"`
	IfInOctets        int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts     int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards      int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors        int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets       int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts    int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards     int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors       int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
}

type SubIPv4Intf struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Ip Address for the interface"`
	IfIndex int32  `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"System generated id for the ipv4Intf where sub interface is to be configured"`
	Type    string `DESCRIPTION:"Type of interface, e.g. Secondary or Virtual"`
	MacAddr string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv4Intf mac address will be used`
	Enable  bool   `DESCRIPTION:"Enable or disable this interface", DEFAULT:false`
}
