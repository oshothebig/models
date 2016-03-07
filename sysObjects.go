package models

import ()

type SystemLoggingConfig struct {
	BaseObj
	SRLogger      string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", DESCRIPTION: "Global logging"`
	SystemLogging string `DESCRIPTION: "Global logging", DEFAULT: "on"`
}

type ComponentLoggingConfig struct {
	BaseObj
	Module string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Module name to set logging level"`
	Level  string `DESCRIPTION: "Logging level", DEFAULT: "info"`
}

type IpTableAclConfig struct {
	BaseObj
	Name     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	IfIndex  int32  `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: 0`
	Access   string `DESCRIPTION: "permit or deny"`
	IpAddr   string `DESCRIPTION: "ip address of subnet or host"`
	NetMask  string `DESCRIPTION: "netmask for ip addr", DEFAULT: "0.0.0.0"`
	Protocol string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP"`
	Port     string `DESCRITION: "port for protocol"`
}
