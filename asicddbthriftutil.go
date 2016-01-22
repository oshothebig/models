package models

import (
	//"models"
	"asicdServices"
)

func ConvertasicdVlanConfigObjToThrift(dbobj *VlanConfig, thriftobj *asicdServices.VlanConfig) {
	thriftobj.IfIndexList = string(dbobj.IfIndexList)
	thriftobj.UntagIfIndexList = string(dbobj.UntagIfIndexList)
	thriftobj.VlanId = int32(dbobj.VlanId)
}

func ConvertThriftToasicdVlanConfigObj(thriftobj *asicdServices.VlanConfig, dbobj *VlanConfig) {
	dbobj.IfIndexList = string(thriftobj.IfIndexList)
	dbobj.UntagIfIndexList = string(thriftobj.UntagIfIndexList)
	dbobj.VlanId = int32(thriftobj.VlanId)
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
	thriftobj.OperState = string(dbobj.OperState)
	thriftobj.MacAddr = string(dbobj.MacAddr)
	thriftobj.Name = string(dbobj.Name)
	thriftobj.Duplex = string(dbobj.Duplex)
	thriftobj.Type = string(dbobj.Type)
	thriftobj.MediaType = string(dbobj.MediaType)
	thriftobj.Mtu = int32(dbobj.Mtu)
	thriftobj.AdminState = string(dbobj.AdminState)
	thriftobj.IfIndex = int32(dbobj.IfIndex)
	thriftobj.Autoneg = string(dbobj.Autoneg)
	thriftobj.Speed = int32(dbobj.Speed)
	thriftobj.Description = string(dbobj.Description)
}

func ConvertThriftToasicdPortConfigObj(thriftobj *asicdServices.PortConfig, dbobj *PortConfig) {
	dbobj.OperState = string(thriftobj.OperState)
	dbobj.MacAddr = string(thriftobj.MacAddr)
	dbobj.Name = string(thriftobj.Name)
	dbobj.Duplex = string(thriftobj.Duplex)
	dbobj.Type = string(thriftobj.Type)
	dbobj.MediaType = string(thriftobj.MediaType)
	dbobj.Mtu = int32(thriftobj.Mtu)
	dbobj.AdminState = string(thriftobj.AdminState)
	dbobj.IfIndex = int32(thriftobj.IfIndex)
	dbobj.Autoneg = string(thriftobj.Autoneg)
	dbobj.Speed = int32(thriftobj.Speed)
	dbobj.Description = string(thriftobj.Description)
}
