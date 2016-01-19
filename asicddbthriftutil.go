package models

import (
	//"models"
	"asicdServices"
)

func ConvertasicdPortIntfConfigObjToThrift(dbobj *PortIntfConfig, thriftobj *asicdServices.PortConfig) {
	thriftobj.OperState = string(dbobj.OperState)
	thriftobj.MacAddr = string(dbobj.MacAddr)
	thriftobj.PortNum = int32(dbobj.PortNum)
	thriftobj.Name = string(dbobj.Name)
	thriftobj.Duplex = string(dbobj.Duplex)
	thriftobj.Type = string(dbobj.Type)
	thriftobj.MediaType = string(dbobj.MediaType)
	thriftobj.Mtu = int32(dbobj.Mtu)
	thriftobj.AdminState = string(dbobj.AdminState)
	thriftobj.Autoneg = string(dbobj.Autoneg)
	thriftobj.Speed = int32(dbobj.Speed)
	thriftobj.Description = string(dbobj.Description)
}

func ConvertThriftToasicdPortIntfConfigObj(thriftobj *asicdServices.PortConfig, dbobj *PortIntfConfig) {
	dbobj.OperState = string(thriftobj.OperState)
	dbobj.MacAddr = string(thriftobj.MacAddr)
	dbobj.PortNum = int32(thriftobj.PortNum)
	dbobj.Name = string(thriftobj.Name)
	dbobj.Duplex = string(thriftobj.Duplex)
	dbobj.Type = string(thriftobj.Type)
	dbobj.MediaType = string(thriftobj.MediaType)
	dbobj.Mtu = int32(thriftobj.Mtu)
	dbobj.AdminState = string(thriftobj.AdminState)
	dbobj.Autoneg = string(thriftobj.Autoneg)
	dbobj.Speed = int32(thriftobj.Speed)
	dbobj.Description = string(thriftobj.Description)
}
