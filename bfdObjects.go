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
	Interface                 int32 `SNAPROUTE: "KEY"`
	LocalMultiplier           uint32
	DesiredMinTxInterval      uint32
	RequiredMinRxInterval     uint32
	RequiredMinEchoRxInterval uint32
	DemandEnabled             bool
	AuthenticationEnabled     bool
	AuthType                  uint32
	AuthKeyId                 uint32
	SequenceNumber            uint32
	AuthData                  string
}

/*
 * BFD Session state
 */
type BfdSessionState struct {
	BaseObj
	SessionId             int32
	LocalIpAddr           string
	RemoteIpAddr          string
	InterfaceId           int32
	RegisteredProtocols   string
	SessionState          int32
	RemoteSessionState    int32
	LocalDiscriminator    uint32
	RemoteDiscriminator   uint32
	LocalDiagType         int32
	DesiredMinTxInterval  int32
	RequiredMinRxInterval int32
	RemoteMinRxInterval   int32
	DetectionMultiplier   uint32
	DemandMode            bool
	RemoteDemandMode      bool
	AuthSeqKnown          bool
	AuthType              uint32
	ReceivedAuthSeq       uint32
	SentAuthSeq           uint32
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
