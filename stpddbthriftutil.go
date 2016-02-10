package models

import (
	//"models"
	"stpd"
)

func ConvertstpdDot1dStpPortEntryConfigObjToThrift(dbobj *Dot1dStpPortEntryConfig, thriftobj *stpd.Dot1dStpPortEntryConfig) {
	thriftobj.Dot1dStpPortAdminPointToPoint = int32(dbobj.Dot1dStpPortAdminPointToPoint)
	thriftobj.Dot1dStpPortKey = int32(dbobj.Dot1dStpPortKey)
	thriftobj.Dot1dBrgIfIndex = int32(dbobj.Dot1dBrgIfIndex)
	thriftobj.Dot1dStpPortPriority = int32(dbobj.Dot1dStpPortPriority)
	thriftobj.Dot1dStpPortAdminPathCost = int32(dbobj.Dot1dStpPortAdminPathCost)
	thriftobj.Dot1dStpPortPathCost = int32(dbobj.Dot1dStpPortPathCost)
	thriftobj.Dot1dStpPortPathCost32 = int32(dbobj.Dot1dStpPortPathCost32)
	thriftobj.Dot1dStpPortProtocolMigration = int32(dbobj.Dot1dStpPortProtocolMigration)
	thriftobj.Dot1dStpPortAdminEdgePort = int32(dbobj.Dot1dStpPortAdminEdgePort)
	thriftobj.Dot1dStpPortEnable = int32(dbobj.Dot1dStpPortEnable)
}

func ConvertThriftTostpdDot1dStpPortEntryConfigObj(thriftobj *stpd.Dot1dStpPortEntryConfig, dbobj *Dot1dStpPortEntryConfig) {
	dbobj.Dot1dStpPortAdminPointToPoint = int32(thriftobj.Dot1dStpPortAdminPointToPoint)
	dbobj.Dot1dStpPortKey = int32(thriftobj.Dot1dStpPortKey)
	dbobj.Dot1dBrgIfIndex = int32(thriftobj.Dot1dBrgIfIndex)
	dbobj.Dot1dStpPortPriority = int32(thriftobj.Dot1dStpPortPriority)
	dbobj.Dot1dStpPortAdminPathCost = int32(thriftobj.Dot1dStpPortAdminPathCost)
	dbobj.Dot1dStpPortPathCost = int32(thriftobj.Dot1dStpPortPathCost)
	dbobj.Dot1dStpPortPathCost32 = int32(thriftobj.Dot1dStpPortPathCost32)
	dbobj.Dot1dStpPortProtocolMigration = int32(thriftobj.Dot1dStpPortProtocolMigration)
	dbobj.Dot1dStpPortAdminEdgePort = int32(thriftobj.Dot1dStpPortAdminEdgePort)
	dbobj.Dot1dStpPortEnable = int32(thriftobj.Dot1dStpPortEnable)
}

func ConvertstpdDot1dStpBridgeConfigObjToThrift(dbobj *Dot1dStpBridgeConfig, thriftobj *stpd.Dot1dStpBridgeConfig) {
	thriftobj.Dot1dStpBridgeForceVersion = int32(dbobj.Dot1dStpBridgeForceVersion)
	thriftobj.Dot1dStpPriorityKey = int32(dbobj.Dot1dStpPriorityKey)
	thriftobj.Dot1dStpBridgeHelloTime = int32(dbobj.Dot1dStpBridgeHelloTime)
	thriftobj.Dot1dStpBridgeTxHoldCount = int32(dbobj.Dot1dStpBridgeTxHoldCount)
	thriftobj.Dot1dStpBridgeForwardDelay = int32(dbobj.Dot1dStpBridgeForwardDelay)
	thriftobj.Dot1dBridgeAddressKey = string(dbobj.Dot1dBridgeAddressKey)
	thriftobj.Dot1dStpVlan = int16(dbobj.Dot1dStpVlan)
	thriftobj.Dot1dStpBridgeMaxAge = int32(dbobj.Dot1dStpBridgeMaxAge)
}

func ConvertThriftTostpdDot1dStpBridgeConfigObj(thriftobj *stpd.Dot1dStpBridgeConfig, dbobj *Dot1dStpBridgeConfig) {
	dbobj.Dot1dStpBridgeForceVersion = int32(thriftobj.Dot1dStpBridgeForceVersion)
	dbobj.Dot1dStpPriorityKey = int32(thriftobj.Dot1dStpPriorityKey)
	dbobj.Dot1dStpBridgeHelloTime = int32(thriftobj.Dot1dStpBridgeHelloTime)
	dbobj.Dot1dStpBridgeTxHoldCount = int32(thriftobj.Dot1dStpBridgeTxHoldCount)
	dbobj.Dot1dStpBridgeForwardDelay = int32(thriftobj.Dot1dStpBridgeForwardDelay)
	dbobj.Dot1dBridgeAddressKey = string(thriftobj.Dot1dBridgeAddressKey)
	dbobj.Dot1dStpVlan = uint16(thriftobj.Dot1dStpVlan)
	dbobj.Dot1dStpBridgeMaxAge = int32(thriftobj.Dot1dStpBridgeMaxAge)
}
