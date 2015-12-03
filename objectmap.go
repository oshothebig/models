package models

var ConfigObjectMap = map[string]ConfigObj{
	"IPV4Route":             &IPV4Route{},         // manually merged from originional
	"Vlan":                  &Vlan{},              // manually added, no YANG defined
	"IPv4Intf":              &IPv4Intf{},          // manually added, no YANG defined
	"IPv4Neighbor":          &IPv4Neighbor{},      // manually added, no YANG defined
	"BGPGlobalConfig":       &BGPGlobalConfig{},   //manually added, no YANG defined
	"BGPNeighborConfig":     &BGPNeighborConfig{}, //manually added, no YANG defined
	"AggregationLacpConfig": &AggregationLacpConfig{},
	"EthernetConfig":        &EthernetConfig{},
	"AggregationConfig":     &AggregationConfig{},
}
