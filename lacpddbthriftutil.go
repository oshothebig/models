package models

import (
	//"models"
	"lacpd"
)

func ConvertlacpdEthernetConfigObjToThrift(dbobj *EthernetConfig, thriftobj *lacpd.EthernetConfig) {
	thriftobj.MacAddress = string(dbobj.MacAddress)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.AggregateId = string(dbobj.AggregateId)
	thriftobj.NameKey = string(dbobj.NameKey)
	thriftobj.Enabled = bool(dbobj.Enabled)
	thriftobj.Speed = string(dbobj.Speed)
	thriftobj.Mtu = int16(dbobj.Mtu)
	thriftobj.DuplexMode = int32(dbobj.DuplexMode)
	thriftobj.EnableFlowControl = bool(dbobj.EnableFlowControl)
	thriftobj.Auto = bool(dbobj.Auto)
	thriftobj.Type = string(dbobj.Type)
}

func ConvertThriftTolacpdEthernetConfigObj(thriftobj *lacpd.EthernetConfig, dbobj *EthernetConfig) {
	dbobj.MacAddress = string(thriftobj.MacAddress)
	dbobj.Description = string(thriftobj.Description)
	dbobj.AggregateId = string(thriftobj.AggregateId)
	dbobj.NameKey = string(thriftobj.NameKey)
	dbobj.Enabled = bool(thriftobj.Enabled)
	dbobj.Speed = string(thriftobj.Speed)
	dbobj.Mtu = uint16(thriftobj.Mtu)
	dbobj.DuplexMode = int32(thriftobj.DuplexMode)
	dbobj.EnableFlowControl = bool(thriftobj.EnableFlowControl)
	dbobj.Auto = bool(thriftobj.Auto)
	dbobj.Type = string(thriftobj.Type)
}

func ConvertlacpdAggregationLacpConfigObjToThrift(dbobj *AggregationLacpConfig, thriftobj *lacpd.AggregationLacpConfig) {
	thriftobj.Description = string(dbobj.Description)
	thriftobj.MinLinks = int16(dbobj.MinLinks)
	thriftobj.SystemPriority = int16(dbobj.SystemPriority)
	thriftobj.NameKey = string(dbobj.NameKey)
	thriftobj.Interval = int32(dbobj.Interval)
	thriftobj.Enabled = bool(dbobj.Enabled)
	thriftobj.Mtu = int16(dbobj.Mtu)
	thriftobj.SystemIdMac = string(dbobj.SystemIdMac)
	thriftobj.LagType = int32(dbobj.LagType)
	thriftobj.LagHash = int32(dbobj.LagHash)
	thriftobj.Type = string(dbobj.Type)
	thriftobj.LacpMode = int32(dbobj.LacpMode)
}

func ConvertThriftTolacpdAggregationLacpConfigObj(thriftobj *lacpd.AggregationLacpConfig, dbobj *AggregationLacpConfig) {
	dbobj.Description = string(thriftobj.Description)
	dbobj.MinLinks = uint16(thriftobj.MinLinks)
	dbobj.SystemPriority = uint16(thriftobj.SystemPriority)
	dbobj.NameKey = string(thriftobj.NameKey)
	dbobj.Interval = int32(thriftobj.Interval)
	dbobj.Enabled = bool(thriftobj.Enabled)
	dbobj.Mtu = uint16(thriftobj.Mtu)
	dbobj.SystemIdMac = string(thriftobj.SystemIdMac)
	dbobj.LagType = int32(thriftobj.LagType)
	dbobj.LagHash = int32(thriftobj.LagHash)
	dbobj.Type = string(thriftobj.Type)
	dbobj.LacpMode = int32(thriftobj.LacpMode)
}
