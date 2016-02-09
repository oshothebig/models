package models

import (
	//"models"
	"dhcprelayd"
)

func ConvertdhcprelaydDhcpRelayGlobalConfigObjToThrift(dbobj *DhcpRelayGlobalConfig, thriftobj *dhcprelayd.DhcpRelayGlobalConfig) {
	thriftobj.Enable = bool(dbobj.Enable)
	thriftobj.DhcpRelay = string(dbobj.DhcpRelay)
}

func ConvertThriftTodhcprelaydDhcpRelayGlobalConfigObj(thriftobj *dhcprelayd.DhcpRelayGlobalConfig, dbobj *DhcpRelayGlobalConfig) {
	dbobj.Enable = bool(thriftobj.Enable)
	dbobj.DhcpRelay = string(thriftobj.DhcpRelay)
}

func ConvertdhcprelaydDhcpRelayIntfConfigObjToThrift(dbobj *DhcpRelayIntfConfig, thriftobj *dhcprelayd.DhcpRelayIntfConfig) {
	thriftobj.Netmask = string(dbobj.Netmask)
	thriftobj.Enable = bool(dbobj.Enable)
	thriftobj.IpSubnet = string(dbobj.IpSubnet)
	thriftobj.IfIndex = string(dbobj.IfIndex)

	for idx, data4 := range dbobj.ServerIp {
		thriftobj.ServerIp[idx] = string(data4)
	}
	thriftobj.AgentSubType = int32(dbobj.AgentSubType)
}

func ConvertThriftTodhcprelaydDhcpRelayIntfConfigObj(thriftobj *dhcprelayd.DhcpRelayIntfConfig, dbobj *DhcpRelayIntfConfig) {
	dbobj.Netmask = string(thriftobj.Netmask)
	dbobj.Enable = bool(thriftobj.Enable)
	dbobj.IpSubnet = string(thriftobj.IpSubnet)
	dbobj.IfIndex = string(thriftobj.IfIndex)

	for _, data4 := range thriftobj.ServerIp {
		dbobj.ServerIp = append(dbobj.ServerIp, string(data4))
	}
	dbobj.AgentSubType = int32(thriftobj.AgentSubType)
}
