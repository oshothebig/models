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
	IpSubnet string `SNAPROUTE: "KEY"`
	Netmask  string `SNAPROUTE: "KEY"`
	//@TODO: Need to check if_index type
	IfIndex string `SNAPROUTE: "KEY"`
	// Use below field for agent sub-type
	AgentSubType int32
	Enable       bool
	// To make life easy for testing first pass lets have only 1 server
	//ServerIp     []string
	ServerIp string
}

func (obj DhcpRelayGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf DhcpRelayGlobalConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling DhcpRelayGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

func (obj DhcpRelayIntfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf DhcpRelayIntfConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling DhcpRelayIntfConfig from Json", body)
		}
	}
	return gConf, err
}
