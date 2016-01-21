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
	thriftobj.MultiHopEnable = bool(dbobj.MultiHopEnable)
	thriftobj.RouteReflectorClient = bool(dbobj.RouteReflectorClient)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.MultiHopTTL = int8(dbobj.MultiHopTTL)
	thriftobj.NeighborAddress = string(dbobj.NeighborAddress)
	thriftobj.PeerAS = int32(dbobj.PeerAS)
	thriftobj.LocalAS = int32(dbobj.LocalAS)
	thriftobj.AuthPassword = string(dbobj.AuthPassword)
}

func ConvertThriftTobgpdBGPNeighborConfigObj(thriftobj *bgpd.BGPNeighborConfig, dbobj *BGPNeighborConfig) {
	dbobj.RouteReflectorClusterId = uint32(thriftobj.RouteReflectorClusterId)
	dbobj.MultiHopEnable = bool(thriftobj.MultiHopEnable)
	dbobj.RouteReflectorClient = bool(thriftobj.RouteReflectorClient)
	dbobj.Description = string(thriftobj.Description)
	dbobj.MultiHopTTL = uint8(thriftobj.MultiHopTTL)
	dbobj.NeighborAddress = string(thriftobj.NeighborAddress)
	dbobj.PeerAS = uint32(thriftobj.PeerAS)
	dbobj.LocalAS = uint32(thriftobj.LocalAS)
	dbobj.AuthPassword = string(thriftobj.AuthPassword)
}
