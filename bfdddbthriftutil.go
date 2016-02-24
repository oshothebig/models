package models

import (
	//"models"
	"bfdd"
)

func ConvertbfddBfdGlobalConfigObjToThrift(dbobj *BfdGlobalConfig, thriftobj *bfdd.BfdGlobalConfig) {
	thriftobj.Enable = bool(dbobj.Enable)
	thriftobj.Bfd = string(dbobj.Bfd)
}

func ConvertThriftTobfddBfdGlobalConfigObj(thriftobj *bfdd.BfdGlobalConfig, dbobj *BfdGlobalConfig) {
	dbobj.Enable = bool(thriftobj.Enable)
	dbobj.Bfd = string(thriftobj.Bfd)
}

func ConvertbfddBfdIntfConfigObjToThrift(dbobj *BfdIntfConfig, thriftobj *bfdd.BfdIntfConfig) {
	thriftobj.DesiredMinTxInterval = int32(dbobj.DesiredMinTxInterval)
	thriftobj.AuthenticationEnabled = bool(dbobj.AuthenticationEnabled)
	thriftobj.DemandEnabled = bool(dbobj.DemandEnabled)
	thriftobj.RequiredMinRxInterval = int32(dbobj.RequiredMinRxInterval)
	thriftobj.IfIndex = int32(dbobj.IfIndex)
	thriftobj.AuthKeyId = int32(dbobj.AuthKeyId)
	thriftobj.RequiredMinEchoRxInterval = int32(dbobj.RequiredMinEchoRxInterval)
	thriftobj.AuthData = string(dbobj.AuthData)
	thriftobj.AuthType = string(dbobj.AuthType)
	thriftobj.LocalMultiplier = int32(dbobj.LocalMultiplier)
}

func ConvertThriftTobfddBfdIntfConfigObj(thriftobj *bfdd.BfdIntfConfig, dbobj *BfdIntfConfig) {
	dbobj.DesiredMinTxInterval = uint32(thriftobj.DesiredMinTxInterval)
	dbobj.AuthenticationEnabled = bool(thriftobj.AuthenticationEnabled)
	dbobj.DemandEnabled = bool(thriftobj.DemandEnabled)
	dbobj.RequiredMinRxInterval = uint32(thriftobj.RequiredMinRxInterval)
	dbobj.IfIndex = int32(thriftobj.IfIndex)
	dbobj.AuthKeyId = uint32(thriftobj.AuthKeyId)
	dbobj.RequiredMinEchoRxInterval = uint32(thriftobj.RequiredMinEchoRxInterval)
	dbobj.AuthData = string(thriftobj.AuthData)
	dbobj.AuthType = string(thriftobj.AuthType)
	dbobj.LocalMultiplier = uint32(thriftobj.LocalMultiplier)
}

func ConvertbfddBfdSessionConfigObjToThrift(dbobj *BfdSessionConfig, thriftobj *bfdd.BfdSessionConfig) {
	thriftobj.Owner = string(dbobj.Owner)
	thriftobj.Operation = string(dbobj.Operation)
	thriftobj.IpAddr = string(dbobj.IpAddr)
	thriftobj.PerLink = bool(dbobj.PerLink)
}

func ConvertThriftTobfddBfdSessionConfigObj(thriftobj *bfdd.BfdSessionConfig, dbobj *BfdSessionConfig) {
	dbobj.Owner = string(thriftobj.Owner)
	dbobj.Operation = string(thriftobj.Operation)
	dbobj.IpAddr = string(thriftobj.IpAddr)
	dbobj.PerLink = bool(thriftobj.PerLink)
}
