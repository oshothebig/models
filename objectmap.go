package models

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
var ConfigObjectMap = map[string]ConfigObj{
	"Vlan":                                          &Vlan{},      // created before auto YANG
	"IPV4Route":                                     &IPV4Route{}, // created before auto YANG
	"IPV4RouteState":                                &IPV4RouteState{},
	"IPv4Intf":                                      &IPv4Intf{},          // created before auto YANG
	"ArpConfig":                                     &ArpConfig{},         // created before auto YANG
	"ArpEntry":                                      &ArpEntry{},          // created before auto YANG
	"BGPGlobalConfig":                               &BGPGlobalConfig{},   // created before auto YANG
	"BGPGlobalState":                                &BGPGlobalState{},    // created before auto YANG
	"BGPNeighborConfig":                             &BGPNeighborConfig{}, // created before auto YANG
	"BGPNeighborState":                              &BGPNeighborState{},  // created before auto YANG
	"PolicyDefinitionSetsPrefixSet":                 &PolicyDefinitionSetsPrefixSet{},
	"PolicyDefinitionStmtMatchPrefixSetCondition":   &PolicyDefinitionStmtMatchPrefixSetCondition{},
	"PolicyDefinitionStmtMatchNeighborSetCondition": &PolicyDefinitionStmtMatchNeighborSetCondition{},
	"PolicyDefinitionStmtMatchTagSetCondition":      &PolicyDefinitionStmtMatchTagSetCondition{},
	"PolicyDefinitionStmtMatchProtocolCondition":    &PolicyDefinitionStmtMatchProtocolCondition{},
	"PolicyDefinitionStmtIgpActions":                &PolicyDefinitionStmtIgpActions{},
	"PolicyDefinitionStmtRouteDispositionAction":    &PolicyDefinitionStmtRouteDispositionAction{},
	"PolicyDefinitionStmtRedistributionAction":      &PolicyDefinitionStmtRedistributionAction{},
	"PolicyDefinitionStmtConfig":                    &PolicyDefinitionStmtConfig{},
	"PolicyDefinitionStmtState":                     &PolicyDefinitionStmtState{},
	"PolicyDefinitionConditionState":                &PolicyDefinitionConditionState{},
	"PolicyDefinitionActionState":                   &PolicyDefinitionActionState{},
	"PolicyDefinitionConfig":                        &PolicyDefinitionConfig{},
	"PolicyDefinitionState":                         &PolicyDefinitionState{},
	"AggregationLacpConfig":                         &AggregationLacpConfig{},
	"EthernetConfig":                                &EthernetConfig{},
	"AggregationLacpMemberStateCounters":            &AggregationLacpMemberStateCounters{},
	"AggregationLacpState":                          &AggregationLacpState{},
	"PortConfig":                                    &PortConfig{},
	"PortState":                                     &PortState{},
	"UserConfig":                                    &UserConfig{},
	"UserState":                                     &UserState{},
	"IPV4AddressBlock":                              &IPV4AddressBlock{},
	"OspfAreaEntryConfig":                           &OspfAreaEntryConfig{},
	"OspfAreaEntryState":                            &OspfAreaEntryState{},
	"OspfStubAreaEntryConfig":                       &OspfStubAreaEntryConfig{},
	"OspfLsdbEntryState":                            &OspfLsdbEntryState{},
	"OspfAreaRangeEntryConfig":                      &OspfAreaRangeEntryConfig{},
	"OspfHostEntryConfig":                           &OspfHostEntryConfig{},
	"OspfHostEntryState":                            &OspfHostEntryState{},
	"OspfIfEntryConfig":                             &OspfIfEntryConfig{},
	"OspfIfEntryState":                              &OspfIfEntryState{},
	"OspfIfMetricEntryConfig":                       &OspfIfMetricEntryConfig{},
	"OspfVirtIfEntryConfig":                         &OspfVirtIfEntryConfig{},
	"OspfNbrEntryConfig":                            &OspfNbrEntryConfig{},
	"OspfNbrEntryState":                             &OspfNbrEntryState{},
	"OspfVirtNbrEntryState":                         &OspfVirtNbrEntryState{},
	"OspfExtLsdbEntryState":                         &OspfExtLsdbEntryState{},
	"OspfAreaAggregateEntryConfig":                  &OspfAreaAggregateEntryConfig{},
	"OspfLocalLsdbEntryState":                       &OspfLocalLsdbEntryState{},
	"OspfVirtLocalLsdbEntryState":                   &OspfVirtLocalLsdbEntryState{},
	"OspfAsLsdbEntryState":                          &OspfAsLsdbEntryState{},
	"OspfAreaLsaCountEntryState":                    &OspfAreaLsaCountEntryState{},
	"OspfGlobalConfig":                              &OspfGlobalConfig{},
	"OspfGlobalState":                               &OspfGlobalState{},
	"Login":                                         &Login{},
	"Logout":                                        &Logout{},
	"Dot1dStpPortEntryConfig":                       &Dot1dStpPortEntryConfig{},
	"Dot1dStpBridgeState":                           &Dot1dStpBridgeState{},
	"Dot1dStpBridgeConfig":                          &Dot1dStpBridgeConfig{},
	"Dot1dStpPortEntryState":                        &Dot1dStpPortEntryState{},
	"DhcpRelayIntfConfig":                           &DhcpRelayIntfConfig{},
	"DhcpRelayGlobalConfig":                         &DhcpRelayGlobalConfig{},
	"BfdGlobalConfig":                               &BfdGlobalConfig{},
	"BfdIntfConfig":                                 &BfdIntfConfig{},
	"BfdGlobalState":                                &BfdGlobalState{},
	"BfdSessionState":                               &BfdSessionState{},
}
