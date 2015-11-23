package models

var ConfigObjectMap = map[string]ConfigObj{
	"IPV4Route":             &IPV4Route{},    // manually merged from originional
	"Vlan":                  &Vlan{},         // manually merged from originional
	"IPv4Intf":              &IPv4Intf{},     // manually merged from originional
	"IPv4Neighbor":          &IPv4Neighbor{}, // manually merged from originional
	"AggregationLacpConfig": &AggregationLacpConfig{},
	"AggregationConfig":     &AggregationConfig{},
	"BgpGlobalConfig":       &BgpGlobalConfig{},
	"BgpNeighborConfig":     &BgpNeighborConfig{},
	"VlanConfig":            &VlanConfig{},
}
