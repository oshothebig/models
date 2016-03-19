package models

var ConfigObjectMap = map[string]ConfigObj{
	"Vlan": &Vlan{}, // created before auto YANG
	//"IPv4Route":                                        &IPv4Route{}, // created before auto YANG
	//"IPv4RouteState":                                   &IPv4RouteState{},
	//"IPv4EventState":                                   &IPv4EventState{},
	"IPv4Intf":          &IPv4Intf{}, // created before auto YANG
	"LogicalIntfConfig": &LogicalIntfConfig{},
	"LogicalIntfState":  &LogicalIntfState{},
	//"ArpConfig":                                        &ArpConfig{},         // created before auto YANG
	//"ArpEntry":                                         &ArpEntry{},          // created before auto YANG
	"AggregationLacpConfig":              &AggregationLacpConfig{},
	"EthernetConfig":                     &EthernetConfig{},
	"AggregationLacpMemberStateCounters": &AggregationLacpMemberStateCounters{},
	"AggregationLacpState":               &AggregationLacpState{},
	//"BfdSession":                                       &BfdSession{},
	//"BfdSessionState":                                  &BfdSessionState{},
	//"BfdGlobal":                                        &BfdGlobal{},
	//"BfdGlobalState":                                   &BfdGlobalState{},
	//"BfdInterface":                                     &BfdInterface{},
	//"BfdInterfaceState":                                &BfdInterfaceState{},
	"DhcpRelayHostDhcpState":   &DhcpRelayHostDhcpState{},
	"DhcpRelayIntfServerState": &DhcpRelayIntfServerState{},
	"DhcpRelayIntfConfig":      &DhcpRelayIntfConfig{},
	"DhcpRelayIntfState":       &DhcpRelayIntfState{},
	"DhcpRelayGlobalConfig":    &DhcpRelayGlobalConfig{},
	//"PolicyDefinitionConfig":                           &PolicyDefinitionConfig{},
	//"PolicyDefinitionState":                            &PolicyDefinitionState{},
	//"PolicyStmtConfig":                                 &PolicyStmtConfig{},
	//"PolicyStmtState":                                  &PolicyStmtState{},
	//"PolicyConditionConfig":                            &PolicyConditionConfig{},
	//"PolicyConditionState":                             &PolicyConditionState{},
	//"PolicyActionConfig":                               &PolicyActionConfig{},
	//"PolicyActionState":                                &PolicyActionState{},
	//"RouteDistanceState":                               &RouteDistanceState{},
	"PortConfig":       &PortConfig{},
	"PortState":        &PortState{},
	"UserConfig":       &UserConfig{},
	"IPV4AddressBlock": &IPV4AddressBlock{},
	"UserState":        &UserState{},
}
