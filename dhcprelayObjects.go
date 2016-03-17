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
	DhcpRelay string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1",DESCRIPTION: "Global Dhcp Relay Agent Information"`
	Enable    bool   `DESCRIPTION: "Global Config stating whether DHCP Relay Agent is enabled on the box or not"`
}

/*
 * This DS will be used while adding/deleting Relay Agent.
 */
type DhcpRelayIntfConfig struct {
	BaseObj
	//IpSubnet     string `SNAPROUTE: "KEY"`
	//Netmask      string `SNAPROUTE: "KEY"`
	IfIndex int32 `SNAPROUTE: "KEY",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION:"Interface index for which Relay Agent Config needs to be done"`
	//AgentSubType int32
	Enable bool `DESCRIPTION: "Enabling/Disabling relay agent per interface"`
	// To make life easy for testing first pass lets have only 1 server
	ServerIp []string `DESCRIPTION: "Dhcp Server(s) where relay agent can relay client dhcp requests"`
}

type DhcpRelayHostDhcpState struct {
	BaseObj
	MacAddr         string `ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Host Hardware/Mac Address"`
	ServerIp        string `DESCRIPTION: "Configured Dhcp Server"`
	OfferedIp       string `DESCRIPTION: "Ip Address offered by Dhcp Server"`
	GatewayIp       string `DESCRIPTION: "Ip Address which client needs to use"`
	AcceptedIp      string `DESCRIPTION: "Ip Address which client accepted"`
	RequestedIp     string `DESCRIPTION: "Ip Address request from client"`
	ClientDiscover  string `DESCRIPTION: "Most recent time stamp of client discover message to dhcp server"`
	ClientRequest   string `DESCRIPTION: "Most recent time stamp of client request message"`
	ClientRequests  int32  `DESCRIPTION: "Total Number of client request message relayed to server"`
	ClientResponses int32  `DESCRIPTION: "Total Number of server response relayed to client"`
	ServerOffer     string `DESCRIPTION: "Most recent time stamp of server offer message"`
	ServerAck       string `DESCRIPTION: "Most recent time stamp of server ack message"`
	ServerRequests  int32  `DESCRIPTION: "Total Number of requests relayed to server"`
	ServerResponses int32  `DESCRIPTION: "Total Number of responses received from server"`
}

type DhcpRelayIntfState struct {
	BaseObj
	IntfId            int32 `ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface Index for which state is required to be collected"`
	TotalDrops        int32 `DESCRIPTION: "Total number of Dhcp Packets dropped by relay agent"`
	TotalDhcpClientRx int32 `DESCRIPTION: "Total number of client requests that camde to relay agent"`
	TotalDhcpClientTx int32 `DESCRIPTION: "Total number of client responses send out by relay agent"`
	TotalDhcpServerRx int32 `DESCRIPTION: "Total number of server requests made by relay agent"`
	TotalDhcpServerTx int32 `DESCRIPTION: "Total number of server responses received by relay agent"`
}

type DhcpRelayIntfServerState struct {
	BaseObj
	IntfId    int32  `ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface Index for which state is required to be collected"`
	ServerIp  string `DESCRIPTION: "Information about any one of configured Dhcp server"`
	Request   int32  `DESCRIPTION: "Total number of requests to Server"`
	Responses int32  `DESCRIPTION: "Total number of responses from Server"`
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
