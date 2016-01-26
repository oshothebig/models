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
	thriftobj.RequiredMinRxInterval = int32(dbobj.RequiredMinRxInterval)
	thriftobj.AuthData = string(dbobj.AuthData)
	thriftobj.DemandEnabled = bool(dbobj.DemandEnabled)
	thriftobj.SequenceNumber = int32(dbobj.SequenceNumber)
	thriftobj.AuthKeyId = int32(dbobj.AuthKeyId)
	thriftobj.DesiredMinTxInterval = int32(dbobj.DesiredMinTxInterval)
	thriftobj.AuthenticationEnabled = bool(dbobj.AuthenticationEnabled)
	thriftobj.Interface = int32(dbobj.Interface)
	thriftobj.RequiredMinEchoRxInterval = int32(dbobj.RequiredMinEchoRxInterval)
	thriftobj.Type = int32(dbobj.Type)
	thriftobj.LocalMultiplier = int32(dbobj.LocalMultiplier)
}

func ConvertThriftTobfddBfdIntfConfigObj(thriftobj *bfdd.BfdIntfConfig, dbobj *BfdIntfConfig) {
	dbobj.RequiredMinRxInterval = uint32(thriftobj.RequiredMinRxInterval)
	dbobj.AuthData = string(thriftobj.AuthData)
	dbobj.DemandEnabled = bool(thriftobj.DemandEnabled)
	dbobj.SequenceNumber = uint32(thriftobj.SequenceNumber)
	dbobj.AuthKeyId = uint32(thriftobj.AuthKeyId)
	dbobj.DesiredMinTxInterval = uint32(thriftobj.DesiredMinTxInterval)
	dbobj.AuthenticationEnabled = bool(thriftobj.AuthenticationEnabled)
	dbobj.Interface = int32(thriftobj.Interface)
	dbobj.RequiredMinEchoRxInterval = uint32(thriftobj.RequiredMinEchoRxInterval)
	dbobj.Type = uint32(thriftobj.Type)
	dbobj.LocalMultiplier = uint32(thriftobj.LocalMultiplier)
}
