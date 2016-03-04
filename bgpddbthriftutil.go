package models

import (
	//"models"
	"bgpd"
)

func ConvertbgpdBGPGlobalConfigObjToThrift(dbobj *BGPGlobalConfig, thriftobj *bgpd.BGPGlobalConfig) {
	thriftobj.RouterId = string(dbobj.RouterId)
	thriftobj.ASNum = int32(dbobj.ASNum)
	thriftobj.IBGPMaxPaths = int32(dbobj.IBGPMaxPaths)
	thriftobj.EBGPMaxPaths = int32(dbobj.EBGPMaxPaths)
	thriftobj.UseMultiplePaths = bool(dbobj.UseMultiplePaths)
	thriftobj.EBGPAllowMultipleAS = bool(dbobj.EBGPAllowMultipleAS)
}

func ConvertThriftTobgpdBGPGlobalConfigObj(thriftobj *bgpd.BGPGlobalConfig, dbobj *BGPGlobalConfig) {
	dbobj.RouterId = string(thriftobj.RouterId)
	dbobj.ASNum = uint32(thriftobj.ASNum)
	dbobj.IBGPMaxPaths = uint32(thriftobj.IBGPMaxPaths)
	dbobj.EBGPMaxPaths = uint32(thriftobj.EBGPMaxPaths)
	dbobj.UseMultiplePaths = bool(thriftobj.UseMultiplePaths)
	dbobj.EBGPAllowMultipleAS = bool(thriftobj.EBGPAllowMultipleAS)
}

func ConvertbgpdBGPNeighborConfigObjToThrift(dbobj *BGPNeighborConfig, thriftobj *bgpd.BGPNeighborConfig) {
	thriftobj.RouteReflectorClusterId = int32(dbobj.RouteReflectorClusterId)
	thriftobj.RouteReflectorClient = bool(dbobj.RouteReflectorClient)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.PeerGroup = string(dbobj.PeerGroup)
	thriftobj.MultiHopTTL = int8(dbobj.MultiHopTTL)
	thriftobj.PeerAS = int32(dbobj.PeerAS)
	thriftobj.KeepaliveTime = int32(dbobj.KeepaliveTime)
	thriftobj.AuthPassword = string(dbobj.AuthPassword)
	thriftobj.AddPathsMaxTx = int8(dbobj.AddPathsMaxTx)
	thriftobj.MultiHopEnable = bool(dbobj.MultiHopEnable)
	thriftobj.AddPathsRx = bool(dbobj.AddPathsRx)
	thriftobj.NeighborAddress = string(dbobj.NeighborAddress)
	thriftobj.HoldTime = int32(dbobj.HoldTime)
	thriftobj.LocalAS = int32(dbobj.LocalAS)
	thriftobj.ConnectRetryTime = int32(dbobj.ConnectRetryTime)
}

func ConvertThriftTobgpdBGPNeighborConfigObj(thriftobj *bgpd.BGPNeighborConfig, dbobj *BGPNeighborConfig) {
	dbobj.RouteReflectorClusterId = uint32(thriftobj.RouteReflectorClusterId)
	dbobj.RouteReflectorClient = bool(thriftobj.RouteReflectorClient)
	dbobj.Description = string(thriftobj.Description)
	dbobj.PeerGroup = string(thriftobj.PeerGroup)
	dbobj.MultiHopTTL = uint8(thriftobj.MultiHopTTL)
	dbobj.PeerAS = uint32(thriftobj.PeerAS)
	dbobj.KeepaliveTime = uint32(thriftobj.KeepaliveTime)
	dbobj.AuthPassword = string(thriftobj.AuthPassword)
	dbobj.AddPathsMaxTx = uint8(thriftobj.AddPathsMaxTx)
	dbobj.MultiHopEnable = bool(thriftobj.MultiHopEnable)
	dbobj.AddPathsRx = bool(thriftobj.AddPathsRx)
	dbobj.NeighborAddress = string(thriftobj.NeighborAddress)
	dbobj.HoldTime = uint32(thriftobj.HoldTime)
	dbobj.LocalAS = uint32(thriftobj.LocalAS)
	dbobj.ConnectRetryTime = uint32(thriftobj.ConnectRetryTime)
}

func ConvertbgpdBGPPeerGroupObjToThrift(dbobj *BGPPeerGroup, thriftobj *bgpd.BGPPeerGroup) {
	thriftobj.RouteReflectorClusterId = int32(dbobj.RouteReflectorClusterId)
	thriftobj.RouteReflectorClient = bool(dbobj.RouteReflectorClient)
	thriftobj.Description = string(dbobj.Description)
	thriftobj.MultiHopTTL = int8(dbobj.MultiHopTTL)
	thriftobj.PeerAS = int32(dbobj.PeerAS)
	thriftobj.KeepaliveTime = int32(dbobj.KeepaliveTime)
	thriftobj.LocalAS = int32(dbobj.LocalAS)
	thriftobj.AddPathsMaxTx = int8(dbobj.AddPathsMaxTx)
	thriftobj.MultiHopEnable = bool(dbobj.MultiHopEnable)
	thriftobj.AddPathsRx = bool(dbobj.AddPathsRx)
	thriftobj.Name = string(dbobj.Name)
	thriftobj.HoldTime = int32(dbobj.HoldTime)
	thriftobj.AuthPassword = string(dbobj.AuthPassword)
	thriftobj.ConnectRetryTime = int32(dbobj.ConnectRetryTime)
}

func ConvertThriftTobgpdBGPPeerGroupObj(thriftobj *bgpd.BGPPeerGroup, dbobj *BGPPeerGroup) {
	dbobj.RouteReflectorClusterId = uint32(thriftobj.RouteReflectorClusterId)
	dbobj.RouteReflectorClient = bool(thriftobj.RouteReflectorClient)
	dbobj.Description = string(thriftobj.Description)
	dbobj.MultiHopTTL = uint8(thriftobj.MultiHopTTL)
	dbobj.PeerAS = uint32(thriftobj.PeerAS)
	dbobj.KeepaliveTime = uint32(thriftobj.KeepaliveTime)
	dbobj.LocalAS = uint32(thriftobj.LocalAS)
	dbobj.AddPathsMaxTx = uint8(thriftobj.AddPathsMaxTx)
	dbobj.MultiHopEnable = bool(thriftobj.MultiHopEnable)
	dbobj.AddPathsRx = bool(thriftobj.AddPathsRx)
	dbobj.Name = string(thriftobj.Name)
	dbobj.HoldTime = uint32(thriftobj.HoldTime)
	dbobj.AuthPassword = string(thriftobj.AuthPassword)
	dbobj.ConnectRetryTime = uint32(thriftobj.ConnectRetryTime)
}

func ConvertbgpdBGPAggregateObjToThrift(dbobj *BGPAggregate, thriftobj *bgpd.BGPAggregate) {
	thriftobj.SendSummaryOnly = bool(dbobj.SendSummaryOnly)
	thriftobj.GenerateASSet = bool(dbobj.GenerateASSet)
	thriftobj.IPPrefix = string(dbobj.IPPrefix)
}

func ConvertThriftTobgpdBGPAggregateObj(thriftobj *bgpd.BGPAggregate, dbobj *BGPAggregate) {
	dbobj.SendSummaryOnly = bool(thriftobj.SendSummaryOnly)
	dbobj.GenerateASSet = bool(thriftobj.GenerateASSet)
	dbobj.IPPrefix = string(thriftobj.IPPrefix)
}
