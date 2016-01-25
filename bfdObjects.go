package models

import ()

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
	NumTotalSessions     uint32
	NumUpSessions        uint32
	NumDownSessions      uint32
	NumAdminDownSessions uint32
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
	Type                      uint32
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
	ReqisteredProtocols   string
	SessionState          int
	RemoteSessionState    int
	LocalDicriminator     uint32
	RemoteDiscriminator   uint32
	LocalDiagType         int
	DesiredMinTxInterval  int
	RequiredMinRxInterval int
	RemoteMinRxInterval   int
	DetectionMultiplier   uint32
	DemandMode            bool
	RemoteDemandMode      bool
	AuthSeqKnown          bool
	AuthType              uint32
	ReceivedAuthSeq       uint32
	SentAuthSeq           uint32
}
