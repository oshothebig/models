package models

import ()

/*
 * Global Config for BFD
 */
type BfdGlobalConfig struct {
	BaseObj
	Bfd    string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1"`
	Enable bool
}

/*
 * Global State
 */
type BfdGlobalState struct {
	BaseObj
	Bfd                  string `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1"`
	Enable               bool
	NumInterfaces        uint32
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
	IfIndex                   int32 `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*"`
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
	IfIndex                   int32 `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*"`
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

/*
 * BFD Session config
 */
type BfdSessionConfig struct {
	BaseObj
	IpAddr    string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*"`
	PerLink   bool
	Owner     string
	Operation string
}

/*
 * BFD Session state
 */
type BfdSessionState struct {
	BaseObj
	SessionId             int32 `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*"`
	LocalIpAddr           string
	RemoteIpAddr          string
	IfIndex               int32
	PerLinkSession        bool
	LocalMacAddr          string
	RemoteMacAddr         string
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
