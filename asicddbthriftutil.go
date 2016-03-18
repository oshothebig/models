package models

import (
	//"models"
	"asicdServices"
)

func ConvertasicdVlanObjToThrift(dbobj *Vlan, thriftobj *asicdServices.Vlan) {
	thriftobj.IfIndexList = string(dbobj.IfIndexList)
	thriftobj.VlanName = string(dbobj.VlanName)
	thriftobj.UntagIfIndexList = string(dbobj.UntagIfIndexList)
	thriftobj.IfIndex = int32(dbobj.IfIndex)
	thriftobj.OperState = string(dbobj.OperState)
	thriftobj.VlanId = int32(dbobj.VlanId)
}

func ConvertThriftToasicdVlanObj(thriftobj *asicdServices.Vlan, dbobj *Vlan) {
	dbobj.IfIndexList = string(thriftobj.IfIndexList)
	dbobj.VlanName = string(thriftobj.VlanName)
	dbobj.UntagIfIndexList = string(thriftobj.UntagIfIndexList)
	dbobj.IfIndex = int32(thriftobj.IfIndex)
	dbobj.OperState = string(thriftobj.OperState)
	dbobj.VlanId = int32(thriftobj.VlanId)
}
func ConvertasicdLogicalIntfConfigObjToThrift(dbobj *LogicalIntfConfig, thriftobj *asicdServices.LogicalIntfConfig) {
	thriftobj.Name = string(dbobj.Name)
	thriftobj.Type = string(dbobj.Type)
}

func ConvertThriftToasicdLogicalIntfConfigObj(thriftobj *asicdServices.LogicalIntfConfig, dbobj *LogicalIntfConfig) {
	dbobj.Name = string(thriftobj.Name)
	dbobj.Type = string(thriftobj.Type)
}

func ConvertasicdIPv4IntfObjToThrift(dbobj *IPv4Intf, thriftobj *asicdServices.IPv4Intf) {
	thriftobj.IfIndex = int32(dbobj.IfIndex)
	thriftobj.IpAddr = string(dbobj.IpAddr)
}

func ConvertThriftToasicdIPv4IntfObj(thriftobj *asicdServices.IPv4Intf, dbobj *IPv4Intf) {
	dbobj.IfIndex = int32(thriftobj.IfIndex)
	dbobj.IpAddr = string(thriftobj.IpAddr)
}

func ConvertasicdPortConfigObjToThrift(dbobj *PortConfig, thriftobj *asicdServices.PortConfig) {
	thriftobj.PhyIntfType = string(dbobj.PhyIntfType)
	thriftobj.AdminState = string(dbobj.AdminState)
	thriftobj.MacAddr = string(dbobj.MacAddr)
	thriftobj.PortNum = int32(dbobj.PortNum)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.Duplex = string(dbobj.Duplex)
	thriftobj.Autoneg = string(dbobj.Autoneg)
	thriftobj.Speed = int32(dbobj.Speed)
	thriftobj.MediaType = string(dbobj.MediaType)
	thriftobj.Mtu = int32(dbobj.Mtu)
}

func ConvertThriftToasicdPortConfigObj(thriftobj *asicdServices.PortConfig, dbobj *PortConfig) {
	dbobj.PhyIntfType = string(thriftobj.PhyIntfType)
	dbobj.AdminState = string(thriftobj.AdminState)
	dbobj.MacAddr = string(thriftobj.MacAddr)
	dbobj.PortNum = int32(thriftobj.PortNum)
	dbobj.Description = string(thriftobj.Description)
	dbobj.Duplex = string(thriftobj.Duplex)
	dbobj.Autoneg = string(thriftobj.Autoneg)
	dbobj.Speed = int32(thriftobj.Speed)
	dbobj.MediaType = string(thriftobj.MediaType)
	dbobj.Mtu = int32(thriftobj.Mtu)
}
