package models

import (
	//"models"
	"bgpd"
)

func ConvertbgpdBGPGlobalConfigObjToThrift(dbobj *BGPGlobalConfig, thriftobj *bgpd.BGPGlobal) {
	thriftobj.RouterId = string(dbobj.RouterId)
	thriftobj.AS = int32(dbobj.ASNum)
}

func ConvertThriftTobgpdBGPGlobalConfigObj(thriftobj *bgpd.BGPGlobal, dbobj *BGPGlobalConfig) {
	dbobj.RouterId = string(thriftobj.RouterId)
	dbobj.ASNum = uint32(thriftobj.AS)
}

func ConvertbgpdBGPNeighborConfigObjToThrift(dbobj *BGPNeighborConfig, thriftobj *bgpd.BGPNeighbor) {
	thriftobj.RouteReflectorClusterId = int32(dbobj.RouteReflectorClusterId)
	thriftobj.RouteReflectorClient = bool(dbobj.RouteReflectorClient)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.NeighborAddress = string(dbobj.NeighborAddress)
	thriftobj.PeerAS = int32(dbobj.PeerAS)
	thriftobj.LocalAS = int32(dbobj.LocalAS)
	thriftobj.AuthPassword = string(dbobj.AuthPassword)
}

func ConvertThriftTobgpdBGPNeighborConfigObj(thriftobj *bgpd.BGPNeighbor, dbobj *BGPNeighborConfig) {
	dbobj.RouteReflectorClusterId = uint32(thriftobj.RouteReflectorClusterId)
	dbobj.RouteReflectorClient = bool(thriftobj.RouteReflectorClient)
	dbobj.Description = string(thriftobj.Description)
	dbobj.NeighborAddress = string(thriftobj.NeighborAddress)
	dbobj.PeerAS = uint32(thriftobj.PeerAS)
	dbobj.LocalAS = uint32(thriftobj.LocalAS)
	dbobj.AuthPassword = string(thriftobj.AuthPassword)
}
