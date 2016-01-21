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
	thriftobj.ServerIp = string(dbobj.ServerIp)
	thriftobj.AgentSubType = int32(dbobj.AgentSubType)
}

func ConvertThriftTodhcprelaydDhcpRelayIntfConfigObj(thriftobj *dhcprelayd.DhcpRelayIntfConfig, dbobj *DhcpRelayIntfConfig) {
	dbobj.Netmask = string(thriftobj.Netmask)
	dbobj.Enable = bool(thriftobj.Enable)
	dbobj.IpSubnet = string(thriftobj.IpSubnet)
	dbobj.IfIndex = string(thriftobj.IfIndex)
	dbobj.ServerIp = string(thriftobj.ServerIp)
	dbobj.AgentSubType = int32(thriftobj.AgentSubType)
}
