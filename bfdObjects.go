package models

import ()

/*
 * Global Config for BFD
 */
type BfdGlobalConfig struct {
	BaseObj
	Bfd    string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", DESCRIPTION: "VRF id where BFD is globally enabled or disabled"`
	Enable bool   `DESCRIPTION: "Global BFD state in this VRF", DEFAULT: "true"`
}

/*
 * Global State
 */
type BfdGlobalState struct {
	BaseObj
	Bfd                  string `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "VRF id for which global BFD state is requested"`
	Enable               bool   `DESCRIPTION: "Global BFD state in this VRF"`
	NumInterfaces        uint32 `DESCRIPTION: "Number of interfaces on which BFD is enabled"`
	NumTotalSessions     uint32 `DESCRIPTION: "Total number of BFD sessions"`
	NumUpSessions        uint32 `DESCRIPTION: "Number of BFD sessions in up state"`
	NumDownSessions      uint32 `DESCRIPTION: "Number of BFD sessions in down state"`
	NumAdminDownSessions uint32 `DESCRIPTION: "Number of BFD sessions in admin down state"`
}

/*
 * BFD Interface config
 */
type BfdIntfConfig struct {
	BaseObj
	IfIndex                   int32  `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Interface index on which BFD configuration will be applied"`
	LocalMultiplier           uint32 `DESCRIPTION: "Detection multiplier", DEFAULT: "3"`
	DesiredMinTxInterval      uint32 `DESCRIPTION: "Desired minimum tx interval in ms", DEFAULT: "1000"`
	RequiredMinRxInterval     uint32 `DESCRIPTION: "Required minimum rx interval in ms", DEFAULT: "1000"`
	RequiredMinEchoRxInterval uint32 `DESCRIPTION: "Required minimum echo rx interval in ms", DEFAULT: "0"`
	DemandEnabled             bool   `DESCRIPTION: "Enable or disable demand mode", DEFAULT: "false"`
	AuthenticationEnabled     bool   `DESCRIPTION: "Enable or disable authentication", DEFAULT: "false"`
	AuthType                  string `DESCRIPTION: "Authentication type", DEFAULT: "simple"`
	AuthKeyId                 uint32 `DESCRIPTION: "Authentication key id", DEFAULT: "1"`
	AuthData                  string `DESCRIPTION: "Authentication password", DEFAULT: "snaproute"`
}

/*
 * BFD Interface state
 */
type BfdIntfState struct {
	BaseObj
	IfIndex                   int32  `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Interface index for which BFD state is requested"`
	Enabled                   bool   `DESCRIPTION: "BFD state on this interface"`
	NumSessions               int32  `DESCRIPTION: "Number of sessions enabled"`
	LocalMultiplier           int32  `DESCRIPTION: "Detection multiplier"`
	DesiredMinTxInterval      string `DESCRIPTION: "Desired minimum tx interval"`
	RequiredMinRxInterval     string `DESCRIPTION: "Required minimum rx interval"`
	RequiredMinEchoRxInterval string `DESCRIPTION: "Required minimum echo rx interval"`
	DemandEnabled             bool   `DESCRIPTION: "Demand mode enabled"`
	AuthenticationEnabled     bool   `DESCRIPTION: "Authentication enabled"`
	AuthenticationType        string `DESCRIPTION: "Authentication type"`
	AuthenticationKeyId       int32  `DESCRIPTION: "Authentication key id"`
	AuthenticationData        string `DESCRIPTION: "Authentication password"`
}

/*
 * BFD Session config
 */
type BfdSessionConfig struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "BFD neighbor IP address"`
	PerLink bool   `DESCRIPTION: "Run BFD sessions on individual link of a LAG if the neighbor is reachable through LAG", DEFAULT: "false"`
	Owner   string `DESCRIPTION: "Module requesting BFD session configuration", DEFAULT: "user"`
}

/*
 * BFD Session state
 */
type BfdSessionState struct {
	BaseObj
	SessionId             int32  `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Session index"`
	LocalIpAddr           string `DESCRIPTION: "My IP address"`
	RemoteIpAddr          string `DESCRIPTION: "Neighbor IP address"`
	IfIndex               int32  `DESCRIPTION: "Interface index"`
	PerLinkSession        bool   `DESCRIPTION: "This is a perlink session on LAG"`
	LocalMacAddr          string `DESCRIPTION: "My MAC address"`
	RemoteMacAddr         string `DESCRIPTION: "Neighbor MAC address"`
	RegisteredProtocols   string `DESCRIPTION: "Registered owners"`
	SessionState          string `DESCRIPTION: "My state"`
	RemoteSessionState    string `DESCRIPTION: "Neighbor state"`
	LocalDiscriminator    uint32 `DESCRIPTION: "My discriminator"`
	RemoteDiscriminator   uint32 `DESCRIPTION: "Neighbor discriminator"`
	LocalDiagType         string `DESCRIPTION: "My diagnostic"`
	DesiredMinTxInterval  string `DESCRIPTION: "My desired minimum tx interval"`
	RequiredMinRxInterval string `DESCRIPTION: "My required minimum rx interval"`
	RemoteMinRxInterval   string `DESCRIPTION: "Neighbor minimum rx interval"`
	DetectionMultiplier   uint32 `DESCRIPTION: "My detection multiplier"`
	DemandMode            bool   `DESCRIPTION: "My demand mode"`
	RemoteDemandMode      bool   `DESCRIPTION: "Neighbor demand mode"`
	AuthSeqKnown          bool   `DESCRIPTION: "Authentication sequence known"`
	AuthType              string `DESCRIPTION: "My Authentication type"`
	ReceivedAuthSeq       uint32 `DESCRIPTION: "Received authentication sequence number"`
	SentAuthSeq           uint32 `DESCRIPTION: "Sent authentication sequence number"`
	NumTxPackets          uint32 `DESCRIPTION: "Number of control packets sent"`
	NumRxPackets          uint32 `DESCRIPTION: "Number of control packets received"`
}
