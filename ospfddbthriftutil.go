package models

import (
	//"models"
	"ospfd"
)

func ConvertospfdOspfAreaEntryConfigObjToThrift(dbobj *OspfAreaEntryConfig, thriftobj *ospfd.OspfAreaEntryConfig) {
	thriftobj.AreaIdKey = string(dbobj.AreaIdKey)
	thriftobj.AuthType = int32(dbobj.AuthType)
	thriftobj.ImportAsExtern = int32(dbobj.ImportAsExtern)
	thriftobj.AreaSummary = int32(dbobj.AreaSummary)
	thriftobj.AreaNssaTranslatorRole = int32(dbobj.AreaNssaTranslatorRole)
	thriftobj.AreaNssaTranslatorStabilityInterval = int32(dbobj.AreaNssaTranslatorStabilityInterval)
}

func ConvertThriftToospfdOspfAreaEntryConfigObj(thriftobj *ospfd.OspfAreaEntryConfig, dbobj *OspfAreaEntryConfig) {
	dbobj.AreaIdKey = string(thriftobj.AreaIdKey)
	dbobj.AuthType = int32(thriftobj.AuthType)
	dbobj.ImportAsExtern = int32(thriftobj.ImportAsExtern)
	dbobj.AreaSummary = int32(thriftobj.AreaSummary)
	dbobj.AreaNssaTranslatorRole = int32(thriftobj.AreaNssaTranslatorRole)
	dbobj.AreaNssaTranslatorStabilityInterval = int32(thriftobj.AreaNssaTranslatorStabilityInterval)
}

func ConvertospfdOspfStubAreaEntryConfigObjToThrift(dbobj *OspfStubAreaEntryConfig, thriftobj *ospfd.OspfStubAreaEntryConfig) {
	thriftobj.StubAreaIdKey = string(dbobj.StubAreaIdKey)
	thriftobj.StubTOSKey = int32(dbobj.StubTOSKey)
	thriftobj.StubMetric = int32(dbobj.StubMetric)
	thriftobj.StubMetricType = int32(dbobj.StubMetricType)
}

func ConvertThriftToospfdOspfStubAreaEntryConfigObj(thriftobj *ospfd.OspfStubAreaEntryConfig, dbobj *OspfStubAreaEntryConfig) {
	dbobj.StubAreaIdKey = string(thriftobj.StubAreaIdKey)
	dbobj.StubTOSKey = int32(thriftobj.StubTOSKey)
	dbobj.StubMetric = int32(thriftobj.StubMetric)
	dbobj.StubMetricType = int32(thriftobj.StubMetricType)
}

func ConvertospfdOspfAreaRangeEntryConfigObjToThrift(dbobj *OspfAreaRangeEntryConfig, thriftobj *ospfd.OspfAreaRangeEntryConfig) {
	thriftobj.AreaRangeMask = string(dbobj.AreaRangeMask)
	thriftobj.AreaRangeAreaIdKey = string(dbobj.AreaRangeAreaIdKey)
	thriftobj.AreaRangeEffect = int32(dbobj.AreaRangeEffect)
	thriftobj.AreaRangeNetKey = string(dbobj.AreaRangeNetKey)
}

func ConvertThriftToospfdOspfAreaRangeEntryConfigObj(thriftobj *ospfd.OspfAreaRangeEntryConfig, dbobj *OspfAreaRangeEntryConfig) {
	dbobj.AreaRangeMask = string(thriftobj.AreaRangeMask)
	dbobj.AreaRangeAreaIdKey = string(thriftobj.AreaRangeAreaIdKey)
	dbobj.AreaRangeEffect = int32(thriftobj.AreaRangeEffect)
	dbobj.AreaRangeNetKey = string(thriftobj.AreaRangeNetKey)
}

func ConvertospfdOspfHostEntryConfigObjToThrift(dbobj *OspfHostEntryConfig, thriftobj *ospfd.OspfHostEntryConfig) {
	thriftobj.HostIpAddressKey = string(dbobj.HostIpAddressKey)
	thriftobj.HostMetric = int32(dbobj.HostMetric)
	thriftobj.HostTOSKey = int32(dbobj.HostTOSKey)
	thriftobj.HostCfgAreaID = string(dbobj.HostCfgAreaID)
}

func ConvertThriftToospfdOspfHostEntryConfigObj(thriftobj *ospfd.OspfHostEntryConfig, dbobj *OspfHostEntryConfig) {
	dbobj.HostIpAddressKey = string(thriftobj.HostIpAddressKey)
	dbobj.HostMetric = int32(thriftobj.HostMetric)
	dbobj.HostTOSKey = int32(thriftobj.HostTOSKey)
	dbobj.HostCfgAreaID = string(thriftobj.HostCfgAreaID)
}

func ConvertospfdOspfIfEntryConfigObjToThrift(dbobj *OspfIfEntryConfig, thriftobj *ospfd.OspfIfEntryConfig) {
	thriftobj.IfMulticastForwarding = int32(dbobj.IfMulticastForwarding)
	thriftobj.IfRetransInterval = int32(dbobj.IfRetransInterval)
	thriftobj.IfType = int32(dbobj.IfType)
	thriftobj.IfAuthKey = string(dbobj.IfAuthKey)
	thriftobj.IfHelloInterval = int32(dbobj.IfHelloInterval)
	thriftobj.IfIpAddressKey = string(dbobj.IfIpAddressKey)
	thriftobj.IfRtrPriority = int32(dbobj.IfRtrPriority)
	thriftobj.AddressLessIfKey = int32(dbobj.AddressLessIfKey)
	thriftobj.IfAreaId = string(dbobj.IfAreaId)
	thriftobj.IfPollInterval = int32(dbobj.IfPollInterval)
	thriftobj.IfAuthType = int32(dbobj.IfAuthType)
	thriftobj.IfAdminStat = int32(dbobj.IfAdminStat)
	thriftobj.IfDemand = bool(dbobj.IfDemand)
	thriftobj.IfTransitDelay = int32(dbobj.IfTransitDelay)
	thriftobj.IfRtrDeadInterval = int32(dbobj.IfRtrDeadInterval)
}

func ConvertThriftToospfdOspfIfEntryConfigObj(thriftobj *ospfd.OspfIfEntryConfig, dbobj *OspfIfEntryConfig) {
	dbobj.IfMulticastForwarding = int32(thriftobj.IfMulticastForwarding)
	dbobj.IfRetransInterval = int32(thriftobj.IfRetransInterval)
	dbobj.IfType = int32(thriftobj.IfType)
	dbobj.IfAuthKey = string(thriftobj.IfAuthKey)
	dbobj.IfHelloInterval = int32(thriftobj.IfHelloInterval)
	dbobj.IfIpAddressKey = string(thriftobj.IfIpAddressKey)
	dbobj.IfRtrPriority = int32(thriftobj.IfRtrPriority)
	dbobj.AddressLessIfKey = int32(thriftobj.AddressLessIfKey)
	dbobj.IfAreaId = string(thriftobj.IfAreaId)
	dbobj.IfPollInterval = int32(thriftobj.IfPollInterval)
	dbobj.IfAuthType = int32(thriftobj.IfAuthType)
	dbobj.IfAdminStat = int32(thriftobj.IfAdminStat)
	dbobj.IfDemand = bool(thriftobj.IfDemand)
	dbobj.IfTransitDelay = int32(thriftobj.IfTransitDelay)
	dbobj.IfRtrDeadInterval = int32(thriftobj.IfRtrDeadInterval)
}

func ConvertospfdOspfIfMetricEntryConfigObjToThrift(dbobj *OspfIfMetricEntryConfig, thriftobj *ospfd.OspfIfMetricEntryConfig) {
	thriftobj.IfMetricTOSKey = int32(dbobj.IfMetricTOSKey)
	thriftobj.IfMetricValue = int32(dbobj.IfMetricValue)
	thriftobj.IfMetricAddressLessIfKey = int32(dbobj.IfMetricAddressLessIfKey)
	thriftobj.IfMetricIpAddressKey = string(dbobj.IfMetricIpAddressKey)
}

func ConvertThriftToospfdOspfIfMetricEntryConfigObj(thriftobj *ospfd.OspfIfMetricEntryConfig, dbobj *OspfIfMetricEntryConfig) {
	dbobj.IfMetricTOSKey = int32(thriftobj.IfMetricTOSKey)
	dbobj.IfMetricValue = int32(thriftobj.IfMetricValue)
	dbobj.IfMetricAddressLessIfKey = int32(thriftobj.IfMetricAddressLessIfKey)
	dbobj.IfMetricIpAddressKey = string(thriftobj.IfMetricIpAddressKey)
}

func ConvertospfdOspfVirtIfEntryConfigObjToThrift(dbobj *OspfVirtIfEntryConfig, thriftobj *ospfd.OspfVirtIfEntryConfig) {
	thriftobj.VirtIfRetransInterval = int32(dbobj.VirtIfRetransInterval)
	thriftobj.VirtIfHelloInterval = int32(dbobj.VirtIfHelloInterval)
	thriftobj.VirtIfNeighborKey = string(dbobj.VirtIfNeighborKey)
	thriftobj.VirtIfAuthType = int32(dbobj.VirtIfAuthType)
	thriftobj.VirtIfRtrDeadInterval = int32(dbobj.VirtIfRtrDeadInterval)
	thriftobj.VirtIfAreaIdKey = string(dbobj.VirtIfAreaIdKey)
	thriftobj.VirtIfTransitDelay = int32(dbobj.VirtIfTransitDelay)
	thriftobj.VirtIfAuthKey = string(dbobj.VirtIfAuthKey)
}

func ConvertThriftToospfdOspfVirtIfEntryConfigObj(thriftobj *ospfd.OspfVirtIfEntryConfig, dbobj *OspfVirtIfEntryConfig) {
	dbobj.VirtIfRetransInterval = int32(thriftobj.VirtIfRetransInterval)
	dbobj.VirtIfHelloInterval = int32(thriftobj.VirtIfHelloInterval)
	dbobj.VirtIfNeighborKey = string(thriftobj.VirtIfNeighborKey)
	dbobj.VirtIfAuthType = int32(thriftobj.VirtIfAuthType)
	dbobj.VirtIfRtrDeadInterval = int32(thriftobj.VirtIfRtrDeadInterval)
	dbobj.VirtIfAreaIdKey = string(thriftobj.VirtIfAreaIdKey)
	dbobj.VirtIfTransitDelay = int32(thriftobj.VirtIfTransitDelay)
	dbobj.VirtIfAuthKey = string(thriftobj.VirtIfAuthKey)
}

func ConvertospfdOspfNbrEntryConfigObjToThrift(dbobj *OspfNbrEntryConfig, thriftobj *ospfd.OspfNbrEntryConfig) {
	thriftobj.NbrIpAddrKey = string(dbobj.NbrIpAddrKey)
	thriftobj.NbrAddressLessIndexKey = int32(dbobj.NbrAddressLessIndexKey)
	thriftobj.NbrPriority = int32(dbobj.NbrPriority)
}

func ConvertThriftToospfdOspfNbrEntryConfigObj(thriftobj *ospfd.OspfNbrEntryConfig, dbobj *OspfNbrEntryConfig) {
	dbobj.NbrIpAddrKey = string(thriftobj.NbrIpAddrKey)
	dbobj.NbrAddressLessIndexKey = int32(thriftobj.NbrAddressLessIndexKey)
	dbobj.NbrPriority = int32(thriftobj.NbrPriority)
}

func ConvertospfdOspfAreaAggregateEntryConfigObjToThrift(dbobj *OspfAreaAggregateEntryConfig, thriftobj *ospfd.OspfAreaAggregateEntryConfig) {
	thriftobj.AreaAggregateNetKey = string(dbobj.AreaAggregateNetKey)
	thriftobj.AreaAggregateAreaIDKey = string(dbobj.AreaAggregateAreaIDKey)
	thriftobj.AreaAggregateMaskKey = string(dbobj.AreaAggregateMaskKey)
	thriftobj.AreaAggregateExtRouteTag = int32(dbobj.AreaAggregateExtRouteTag)
	thriftobj.AreaAggregateLsdbTypeKey = int32(dbobj.AreaAggregateLsdbTypeKey)
	thriftobj.AreaAggregateEffect = int32(dbobj.AreaAggregateEffect)
}

func ConvertThriftToospfdOspfAreaAggregateEntryConfigObj(thriftobj *ospfd.OspfAreaAggregateEntryConfig, dbobj *OspfAreaAggregateEntryConfig) {
	dbobj.AreaAggregateNetKey = string(thriftobj.AreaAggregateNetKey)
	dbobj.AreaAggregateAreaIDKey = string(thriftobj.AreaAggregateAreaIDKey)
	dbobj.AreaAggregateMaskKey = string(thriftobj.AreaAggregateMaskKey)
	dbobj.AreaAggregateExtRouteTag = uint32(thriftobj.AreaAggregateExtRouteTag)
	dbobj.AreaAggregateLsdbTypeKey = int32(thriftobj.AreaAggregateLsdbTypeKey)
	dbobj.AreaAggregateEffect = int32(thriftobj.AreaAggregateEffect)
}

func ConvertospfdOspfGlobalConfigObjToThrift(dbobj *OspfGlobalConfig, thriftobj *ospfd.OspfGlobalConfig) {
	thriftobj.DiscontinuityTime = int32(dbobj.DiscontinuityTime)
	thriftobj.StubRouterAdvertisement = int32(dbobj.StubRouterAdvertisement)
	thriftobj.RouterIdKey = string(dbobj.RouterIdKey)
	thriftobj.RestartSupport = int32(dbobj.RestartSupport)
	thriftobj.AsLsaCount = int32(dbobj.AsLsaCount)
	thriftobj.OpaqueLsaSupport = bool(dbobj.OpaqueLsaSupport)
	thriftobj.OriginateNewLsas = int32(dbobj.OriginateNewLsas)
	thriftobj.ExternLsaCount = int32(dbobj.ExternLsaCount)
	thriftobj.RestartExitReason = int32(dbobj.RestartExitReason)
	thriftobj.RxNewLsas = int32(dbobj.RxNewLsas)
	thriftobj.TOSSupport = bool(dbobj.TOSSupport)
	thriftobj.ExternLsaCksumSum = int32(dbobj.ExternLsaCksumSum)
	thriftobj.RestartStatus = int32(dbobj.RestartStatus)
	thriftobj.RestartAge = int32(dbobj.RestartAge)
	thriftobj.ExtLsdbLimit = int32(dbobj.ExtLsdbLimit)
	thriftobj.DemandExtensions = bool(dbobj.DemandExtensions)
	thriftobj.RestartInterval = int32(dbobj.RestartInterval)
	thriftobj.RFC1583Compatibility = bool(dbobj.RFC1583Compatibility)
	thriftobj.ReferenceBandwidth = int32(dbobj.ReferenceBandwidth)
	thriftobj.VersionNumber = int32(dbobj.VersionNumber)
	thriftobj.ASBdrRtrStatus = bool(dbobj.ASBdrRtrStatus)
	thriftobj.RestartStrictLsaChecking = bool(dbobj.RestartStrictLsaChecking)
	thriftobj.ExitOverflowInterval = int32(dbobj.ExitOverflowInterval)
	thriftobj.AdminStat = int32(dbobj.AdminStat)
	thriftobj.MulticastExtensions = int32(dbobj.MulticastExtensions)
	thriftobj.AreaBdrRtrStatus = bool(dbobj.AreaBdrRtrStatus)
	thriftobj.StubRouterSupport = bool(dbobj.StubRouterSupport)
	thriftobj.AsLsaCksumSum = int32(dbobj.AsLsaCksumSum)
}

func ConvertThriftToospfdOspfGlobalConfigObj(thriftobj *ospfd.OspfGlobalConfig, dbobj *OspfGlobalConfig) {
	dbobj.DiscontinuityTime = uint32(thriftobj.DiscontinuityTime)
	dbobj.StubRouterAdvertisement = int32(thriftobj.StubRouterAdvertisement)
	dbobj.RouterIdKey = string(thriftobj.RouterIdKey)
	dbobj.RestartSupport = int32(thriftobj.RestartSupport)
	dbobj.AsLsaCount = uint32(thriftobj.AsLsaCount)
	dbobj.OpaqueLsaSupport = bool(thriftobj.OpaqueLsaSupport)
	dbobj.OriginateNewLsas = uint32(thriftobj.OriginateNewLsas)
	dbobj.ExternLsaCount = uint32(thriftobj.ExternLsaCount)
	dbobj.RestartExitReason = int32(thriftobj.RestartExitReason)
	dbobj.RxNewLsas = uint32(thriftobj.RxNewLsas)
	dbobj.TOSSupport = bool(thriftobj.TOSSupport)
	dbobj.ExternLsaCksumSum = int32(thriftobj.ExternLsaCksumSum)
	dbobj.RestartStatus = int32(thriftobj.RestartStatus)
	dbobj.RestartAge = uint32(thriftobj.RestartAge)
	dbobj.ExtLsdbLimit = int32(thriftobj.ExtLsdbLimit)
	dbobj.DemandExtensions = bool(thriftobj.DemandExtensions)
	dbobj.RestartInterval = int32(thriftobj.RestartInterval)
	dbobj.RFC1583Compatibility = bool(thriftobj.RFC1583Compatibility)
	dbobj.ReferenceBandwidth = uint32(thriftobj.ReferenceBandwidth)
	dbobj.VersionNumber = int32(thriftobj.VersionNumber)
	dbobj.ASBdrRtrStatus = bool(thriftobj.ASBdrRtrStatus)
	dbobj.RestartStrictLsaChecking = bool(thriftobj.RestartStrictLsaChecking)
	dbobj.ExitOverflowInterval = int32(thriftobj.ExitOverflowInterval)
	dbobj.AdminStat = int32(thriftobj.AdminStat)
	dbobj.MulticastExtensions = int32(thriftobj.MulticastExtensions)
	dbobj.AreaBdrRtrStatus = bool(thriftobj.AreaBdrRtrStatus)
	dbobj.StubRouterSupport = bool(thriftobj.StubRouterSupport)
	dbobj.AsLsaCksumSum = uint32(thriftobj.AsLsaCksumSum)
}
