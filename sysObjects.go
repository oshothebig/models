package models

import ()

type SystemLogging struct {
	ConfigObj
	SRLogger string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", DESCRIPTION: "Global logging"`
	Logging  string `DESCRIPTION: "Global logging", DEFAULT: "on"`
}

type ComponentLogging struct {
	ConfigObj
	Module string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Module name to set logging level"`
	Level  string `DESCRIPTION: "Logging level", DEFAULT: "info"`
}

type IpTableAcl struct {
	ConfigObj
	Name         string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	PhysicalPort string `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: "all"`
	Action       string `DESCRIPTION: "ACCEPT or DROP"`
	IpAddr       string `DESCRIPTION: "ip address of subnet or host, e.g: 192.168.1.0/24, 192.168.1.1"`
	Protocol     string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP"`
	Port         string `DESCRITION: "port for protocol, e.g for dhcprelay port is 68", DEFAULT: "all"`
}

/*
type IpTableAclState struct {
	ConfigObj
	Name         string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	PhysicalPort string `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: "all"`
	Action       string `DESCRIPTION: "ACCEPT or DROP"`
	IpAddr       string `DESCRIPTION: "ip address of subnet or host, e.g: 192.168.1.0/24, 192.168.1.1"`
	Protocol     string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP"`
	Port         string `DESCRITION: "port for protocol, e.g for dhcprelay port is 68", DEFAULT: "all"`
}
*/

type Daemon struct {
	ConfigObj
	Name  string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Daemon name"`
	State string `DESCRIPTION: "State of the daemon, SELECTION: start/stop", DEFAULT: "start"`
}

type DaemonState struct {
	ConfigObj
	Name          string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Daemon name"`
	State         string `DESCRIPTION: "State of the daemon, running or stopped"`
	Reason        string `DESCRIPTION: "Reason for current state of the daemon"`
	KeepAlive     string `DESCRIPTION: "KeepAlive state of the daemon"`
	RestartCount  int32  `DESCRIPTION: "Number of times this daemon has been restarted"`
	RestartTime   string `DESCRIPTION: "Last restart time"`
	RestartReason string `DESCRIPTION: "Last restart reason"`
}
