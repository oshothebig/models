package models

var ConfigObjectMap = map[string]ConfigObj{
	"Vlan":                                             &Vlan{},      // created before auto YANG
	"IPV4Route":                                        &IPV4Route{}, // created before auto YANG
	"IPV4RouteState":                                   &IPV4RouteState{},
	"IPV4EventState":                                   &IPV4EventState{},
	"IPv4Intf":                                         &IPv4Intf{},          // created before auto YANG
	"ArpConfig":                                        &ArpConfig{},         // created before auto YANG
	"ArpEntry":                                         &ArpEntry{},          // created before auto YANG
	"BGPGlobalConfig":                                  &BGPGlobalConfig{},   // created before auto YANG
	"BGPGlobalState":                                   &BGPGlobalState{},    // created before auto YANG
	"BGPNeighborConfig":                                &BGPNeighborConfig{}, // created before auto YANG
	"BGPNeighborState":                                 &BGPNeighborState{},  // created before auto YANG
	"BGPPeerGroup":                                     &BGPPeerGroup{},      // created before auto YANG
	"BGPRoute":                                         &BGPRoute{},          // created before auto YANG
	"PolicyDefinitionSetsPrefixSet":                    &PolicyDefinitionSetsPrefixSet{},
	"PolicyDefinitionStmtDstIpMatchPrefixSetCondition": &PolicyDefinitionStmtDstIpMatchPrefixSetCondition{},
	"PolicyDefinitionStmtMatchNeighborSetCondition":    &PolicyDefinitionStmtMatchNeighborSetCondition{},
	"PolicyDefinitionStmtMatchTagSetCondition":         &PolicyDefinitionStmtMatchTagSetCondition{},
	"AggregationLacpConfig":                            &AggregationLacpConfig{},
	"EthernetConfig":                                   &EthernetConfig{},
	"AggregationLacpMemberStateCounters":               &AggregationLacpMemberStateCounters{},
	"AggregationLacpState":                             &AggregationLacpState{},
	"BfdSessionState":                                  &BfdSessionState{},
	"BfdGlobalConfig":                                  &BfdGlobalConfig{},
	"BfdGlobalState":                                   &BfdGlobalState{},
	"BfdIntfConfig":                                    &BfdIntfConfig{},
	"DhcpRelayHostDhcpState":                           &DhcpRelayHostDhcpState{},
	"DhcpRelayIntfServerState":                         &DhcpRelayIntfServerState{},
	"DhcpRelayIntfConfig":                              &DhcpRelayIntfConfig{},
	"DhcpRelayIntfState":                               &DhcpRelayIntfState{},
	"DhcpRelayGlobalConfig":                            &DhcpRelayGlobalConfig{},
	"OspfNbrEntryState":                                &OspfNbrEntryState{},
	"OspfGlobalConfig":                                 &OspfGlobalConfig{},
	"OspfExtLsdbEntryState":                            &OspfExtLsdbEntryState{},
	"OspfAsLsdbEntryState":                             &OspfAsLsdbEntryState{},
	"OspfLocalLsdbEntryState":                          &OspfLocalLsdbEntryState{},
	"OspfNbrEntryConfig":                               &OspfNbrEntryConfig{},
	"OspfAreaEntryConfig":                              &OspfAreaEntryConfig{},
	"OspfAreaEntryState":                               &OspfAreaEntryState{},
	"OspfVirtIfEntryConfig":                            &OspfVirtIfEntryConfig{},
	"OspfLsdbEntryState":                               &OspfLsdbEntryState{},
	"OspfAreaRangeEntryConfig":                         &OspfAreaRangeEntryConfig{},
	"OspfIfMetricEntryConfig":                          &OspfIfMetricEntryConfig{},
	"OspfVirtLocalLsdbEntryState":                      &OspfVirtLocalLsdbEntryState{},
	"OspfAreaLsaCountEntryState":                       &OspfAreaLsaCountEntryState{},
	"OspfAreaAggregateEntryConfig":                     &OspfAreaAggregateEntryConfig{},
	"OspfHostEntryConfig":                              &OspfHostEntryConfig{},
	"OspfIfEntryConfig":                                &OspfIfEntryConfig{},
	"OspfVirtNbrEntryState":                            &OspfVirtNbrEntryState{},
	"OspfStubAreaEntryConfig":                          &OspfStubAreaEntryConfig{},
	"OspfIfEntryState":                                 &OspfIfEntryState{},
	"OspfGlobalState":                                  &OspfGlobalState{},
	"PolicyDefinitionStmtConfig":                       &PolicyDefinitionStmtConfig{},
	"PolicyDefinitionConditionState":                   &PolicyDefinitionConditionState{},
	"PolicyDefinitionStmtRouteDispositionAction":       &PolicyDefinitionStmtRouteDispositionAction{},
	"PolicyDefinitionStmtRedistributionAction":         &PolicyDefinitionStmtRedistributionAction{},
	"PolicyDefinitionState":                            &PolicyDefinitionState{},
	"PolicyDefinitionStmtState":                        &PolicyDefinitionStmtState{},
	"PolicyDefinitionActionState":                      &PolicyDefinitionActionState{},
	"PolicyDefinitionStmtMatchProtocolCondition":       &PolicyDefinitionStmtMatchProtocolCondition{},
	"PolicyDefinitionStmtIgpActions":                   &PolicyDefinitionStmtIgpActions{},
	"PolicyDefinitionStmtAdminDistanceAction":          &PolicyDefinitionStmtAdminDistanceAction{},
        "RouteDistanceState":                               &RouteDistanceState{},
	"PolicyDefinitionConfig":                           &PolicyDefinitionConfig{},
	"PortConfig":                                       &PortConfig{},
	"PortState":                                        &PortState{},
	"OspfHostEntryState":                               &OspfHostEntryState{},
	"UserConfig":                                       &UserConfig{},
	"IPV4AddressBlock":                                 &IPV4AddressBlock{},
	"UserState":                                        &UserState{},
	"Dot1dStpPortEntryConfig":                          &Dot1dStpPortEntryConfig{},
	"Dot1dStpBridgeState":                              &Dot1dStpBridgeState{},
	"Dot1dStpBridgeConfig":                             &Dot1dStpBridgeConfig{},
	"Dot1dStpPortEntryStateCounters":                   &Dot1dStpPortEntryStateCounters{},
}
