package models

var ConfigObjectMap = map[string]ConfigObj{
	"Vlan":                                    &Vlan{},      // created before auto YANG
	"IPV4Route":                               &IPV4Route{}, // created before auto YANG
	"IPV4RouteState":                          &IPV4RouteState{},
	"IPV4EventState":                          &IPV4EventState{},
	"IPv4Intf":                                &IPv4Intf{},          // created before auto YANG
	"ArpConfig":                               &ArpConfig{},         // created before auto YANG
	"ArpEntry":                                &ArpEntry{},          // created before auto YANG
	"BGPGlobalConfig":                         &BGPGlobalConfig{},   // created before auto YANG
	"BGPGlobalState":                          &BGPGlobalState{},    // created before auto YANG
	"BGPNeighborConfig":                       &BGPNeighborConfig{}, // created before auto YANG
	"BGPNeighborState":                        &BGPNeighborState{},  // created before auto YANG
	"BGPPeerGroup":                            &BGPPeerGroup{},      // created before auto YANG
	"BGPRoute":                                &BGPRoute{},          // created before auto YANG
	"AggregationLacpConfig":                   &AggregationLacpConfig{},
	"EthernetConfig":                          &EthernetConfig{},
	"AggregationLacpMemberStateCounters":      &AggregationLacpMemberStateCounters{},
	"AggregationLacpState":                    &AggregationLacpState{},
	"BfdSessionConfig":                        &BfdSessionConfig{},
	"BfdSessionState":                         &BfdSessionState{},
	"BfdGlobalConfig":                         &BfdGlobalConfig{},
	"BfdGlobalState":                          &BfdGlobalState{},
	"BfdIntfConfig":                           &BfdIntfConfig{},
	"BfdIntfState":                            &BfdIntfState{},
	"DhcpRelayHostDhcpState":                  &DhcpRelayHostDhcpState{},
	"DhcpRelayIntfServerState":                &DhcpRelayIntfServerState{},
	"DhcpRelayIntfConfig":                     &DhcpRelayIntfConfig{},
	"DhcpRelayIntfState":                      &DhcpRelayIntfState{},
	"DhcpRelayGlobalConfig":                   &DhcpRelayGlobalConfig{},
	"PolicyDefinitionConfig":                  &PolicyDefinitionConfig{},
	"PolicyDefinitionState":                   &PolicyDefinitionState{},
	"PolicyStmtConfig":                        &PolicyStmtConfig{},
	"PolicyStmtState":                         &PolicyStmtState{},
	"PolicyPrefixSet":                         &PolicyPrefixSet{},
	"PolicyConditionConfig":                   &PolicyConditionConfig{},
	"PolicyConditionState":                    &PolicyConditionState{},
	"PolicyActionConfig":                      &PolicyActionConfig{},
	"PolicyActionState":                       &PolicyActionState{},
	"RouteDistanceState":                      &RouteDistanceState{},
	"PortConfig":                              &PortConfig{},
	"PortState":                               &PortState{},
	"UserConfig":                              &UserConfig{},
	"IPV4AddressBlock":                        &IPV4AddressBlock{},
	"UserState":                               &UserState{},
	"Dot1dStpPortEntryConfig":                 &Dot1dStpPortEntryConfig{},
	"Dot1dStpBridgeState":                     &Dot1dStpBridgeState{},
	"Dot1dStpBridgeConfig":                    &Dot1dStpBridgeConfig{},
	"Dot1dStpPortEntryStateCountersFsmStates": &Dot1dStpPortEntryStateCountersFsmStates{},
}
