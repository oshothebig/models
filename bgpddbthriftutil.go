package models

import (
	//"models"
	"bgpd"
)

func ConvertbgpdBGPGlobalConfigObjToThrift(dbobj *BGPGlobalConfig, thriftobj *bgpd.BGPGlobalConfig) {
	thriftobj.RouterId = string(dbobj.RouterId)
	thriftobj.ASNum = int32(dbobj.ASNum)
}

func ConvertThriftTobgpdBGPGlobalConfigObj(thriftobj *bgpd.BGPGlobalConfig, dbobj *BGPGlobalConfig) {
	dbobj.RouterId = string(thriftobj.RouterId)
	dbobj.ASNum = uint32(thriftobj.ASNum)
}

func ConvertbgpdBGPNeighborConfigObjToThrift(dbobj *BGPNeighborConfig, thriftobj *bgpd.BGPNeighborConfig) {
	thriftobj.RouteReflectorClusterId = int32(dbobj.RouteReflectorClusterId)
	thriftobj.RouteReflectorClient = bool(dbobj.RouteReflectorClient)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.MultiHopTTL = int8(dbobj.MultiHopTTL)
	thriftobj.PeerAS = int32(dbobj.PeerAS)
	thriftobj.KeepaliveTime = int32(dbobj.KeepaliveTime)
	thriftobj.AuthPassword = string(dbobj.AuthPassword)
	thriftobj.MultiHopEnable = bool(dbobj.MultiHopEnable)
	thriftobj.NeighborAddress = string(dbobj.NeighborAddress)
	thriftobj.HoldTime = int32(dbobj.HoldTime)
	thriftobj.LocalAS = int32(dbobj.LocalAS)
	thriftobj.ConnectRetryTime = int32(dbobj.ConnectRetryTime)
}

func ConvertThriftTobgpdBGPNeighborConfigObj(thriftobj *bgpd.BGPNeighborConfig, dbobj *BGPNeighborConfig) {
	dbobj.RouteReflectorClusterId = uint32(thriftobj.RouteReflectorClusterId)
	dbobj.RouteReflectorClient = bool(thriftobj.RouteReflectorClient)
	dbobj.Description = string(thriftobj.Description)
	dbobj.MultiHopTTL = uint8(thriftobj.MultiHopTTL)
	dbobj.PeerAS = uint32(thriftobj.PeerAS)
	dbobj.KeepaliveTime = uint32(thriftobj.KeepaliveTime)
	dbobj.AuthPassword = string(thriftobj.AuthPassword)
	dbobj.MultiHopEnable = bool(thriftobj.MultiHopEnable)
	dbobj.NeighborAddress = string(thriftobj.NeighborAddress)
	dbobj.HoldTime = uint32(thriftobj.HoldTime)
	dbobj.LocalAS = uint32(thriftobj.LocalAS)
	dbobj.ConnectRetryTime = uint32(thriftobj.ConnectRetryTime)
}
