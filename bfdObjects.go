package models

import (
	"encoding/json"
	"fmt"
)

/*
 * Global Config for BFD
 */
type BfdGlobalConfig struct {
	BaseObj
	Bfd    string `SNAPROUTE: "KEY"`
	Enable bool
}

/*
 * Global State
 */
type BfdGlobalState struct {
	BaseObj
	Enable               bool
	NumInterfaces        uint32
	NumTotalSessions     uint32
	NumUpSessions        uint32
	NumDownSessions      uint32
	NumAdminDownSessions uint32
}

func (obj BfdGlobalState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BfdGlobalState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BfdGlobalState from Json", body)
		}
	}
	return gConf, err
}

/*
 * BFD Interface config
 */
type BfdIntfConfig struct {
	BaseObj
	IfIndex                   int32 `SNAPROUTE: "KEY"`
	LocalMultiplier           uint32
	DesiredMinTxInterval      uint32
	RequiredMinRxInterval     uint32
	RequiredMinEchoRxInterval uint32
	DemandEnabled             bool
	AuthenticationEnabled     bool
	AuthType                  string
	AuthKeyId                 uint32
	AuthData                  string
}

/*
 * BFD Interface state
 */
type BfdIntfState struct {
	BaseObj
	IfIndex                   int32
	Enabled                   bool
	NumSessions               int32
	LocalMultiplier           int32
	DesiredMinTxInterval      string
	RequiredMinRxInterval     string
	RequiredMinEchoRxInterval string
	DemandEnabled             bool
	AuthenticationEnabled     bool
	AuthenticationType        string
	AuthenticationKeyId       int32
	AuthenticationData        string
}

func (obj BfdIntfState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BfdIntfState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BfdIntfState from Json", body)
		}
	}
	return gConf, err
}

/*
 * BFD Session config
 */
type BfdSessionConfig struct {
	BaseObj
	IpAddr    string `SNAPROUTE: "KEY"`
	PerLink   bool
	Owner     string
	Operation string
}

/*
 * BFD Session state
 */
type BfdSessionState struct {
	BaseObj
	SessionId             int32
	LocalIpAddr           string
	RemoteIpAddr          string
	IfIndex               int32
	RegisteredProtocols   string
	SessionState          string
	RemoteSessionState    string
	LocalDiscriminator    uint32
	RemoteDiscriminator   uint32
	LocalDiagType         string
	DesiredMinTxInterval  string
	RequiredMinRxInterval string
	RemoteMinRxInterval   string
	DetectionMultiplier   uint32
	DemandMode            bool
	RemoteDemandMode      bool
	AuthSeqKnown          bool
	AuthType              string
	ReceivedAuthSeq       uint32
	SentAuthSeq           uint32
	NumTxPackets          uint32
	NumRxPackets          uint32
}

func (obj BfdSessionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BfdSessionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BfdSessionState from Json", body)
		}
	}
	return gConf, err
}
