package models

import (
	"encoding/json"
	"fmt"
)

/*
 * Global DataStructure for DHCP RELAY
 */
type DhcpRelayGlobalConfig struct {
	BaseObj
	// This will tell whether DHCP RELAY is enabled/disabled
	// on the box right now or not.
	DhcpRelay string `SNAPROUTE: "KEY"`
	Enable    bool
}

/*
 * This DS will be used while adding/deleting Relay Agent.
 */
type DhcpRelayIntfConfig struct {
	BaseObj
	IpSubnet     string `SNAPROUTE: "KEY"`
	Netmask      string `SNAPROUTE: "KEY"`
	IfIndex      string `SNAPROUTE: "KEY"`
	AgentSubType int32
	Enable       bool
	// To make life easy for testing first pass lets have only 1 server
	ServerIp []string
	//ServerIp string
}

type DhcpRelayHostDhcpState struct {
	BaseObj
	MacAddr        string
	ServerIp       string
	OfferedIp      string
	GatewayIp      string
	AcceptedIp     string
	LeaseDuration  string
	ClientRequest  string
	ClientResponse string
	ServerRequest  string
	ServerResponse string
}

type DhcpRelayIntfState struct {
	BaseObj
	IntfId            int32
	TotalDrops        int32
	TotalDhcpClientRx int32
	TotalDhcpClientTx int32
	TotalDhcpServerRx int32
	TotalDhcpServerTx int32
}

type DhcpRelayIntfServerState struct {
	BaseObj
	IntfId    int32
	ServerIp  string
	Request   int32
	Responses int32
}

func (obj DhcpRelayHostDhcpState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf DhcpRelayHostDhcpState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling DhcpRelayGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

func (obj DhcpRelayIntfState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf DhcpRelayIntfState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling DhcpRelayGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

func (obj DhcpRelayIntfServerState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf DhcpRelayIntfServerState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling DhcpRelayGlobalConfig from Json", body)
		}
	}
	return gConf, err
}
