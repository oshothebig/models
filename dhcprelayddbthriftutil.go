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
	thriftobj.IfIndex = int32(dbobj.IfIndex)
	thriftobj.Enable = bool(dbobj.Enable)

	for _, data2 := range dbobj.ServerIp {
		thriftobj.ServerIp = append(thriftobj.ServerIp, string(data2))
	}
}

func ConvertThriftTodhcprelaydDhcpRelayIntfConfigObj(thriftobj *dhcprelayd.DhcpRelayIntfConfig, dbobj *DhcpRelayIntfConfig) {
	dbobj.IfIndex = int32(thriftobj.IfIndex)
	dbobj.Enable = bool(thriftobj.Enable)

	for _, data2 := range thriftobj.ServerIp {
		dbobj.ServerIp = append(dbobj.ServerIp, string(data2))
	}
}
