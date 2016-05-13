package models

type SourcePolicyList struct {
	Sources string `DESCRIPTION: "Source Protocol(s) which BGP is interested in. Multiple sources can be specified as comma separated strings when the same policy needs to be applied"`
	Policy  string `DESCRIPTION: "Policy that needs to be applied for redistribution of the specified sources into BGP"`
}
type BGPGlobal struct {
	baseObj
	ASNum               uint32             `DESCRIPTION: "Local AS for BGP global config"`
	RouterId            string             `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "Router id for BGP global config"`
	UseMultiplePaths    bool               `DESCRIPTION: "Enable/disable ECMP for BGP", DEFAULT: "false"`
	EBGPMaxPaths        uint32             `DESCRIPTION: "Max ECMP paths from External BGP neighbors", DEFAULT: "0"`
	EBGPAllowMultipleAS bool               `DESCRIPTION: "Enable/diable ECMP paths from multiple ASes", DEFAULT: "false"`
	IBGPMaxPaths        uint32             `DESCRIPTION: "Max ECMP paths from Internal BGP neighbors", DEFAULT: "0"`
	Redistribution      []SourcePolicyList `DESCRIPTION: "Provide redistribution policies for BGP from different sources"`
}

type BGPGlobalState struct {
	baseObj
	AS                  uint32 `DESCRIPTION: "Local AS for BGP global config"`
	RouterId            string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "Router id for BGP global config"`
	UseMultiplePaths    bool   `DESCRIPTION: "Enable/disable ECMP for BGP"`
	EBGPMaxPaths        uint32 `DESCRIPTION: "Max ECMP paths from External BGP neighbors"`
	EBGPAllowMultipleAS bool   `DESCRIPTION: "Enable/diable ECMP paths from multiple ASes"`
	IBGPMaxPaths        uint32 `DESCRIPTION: "Max ECMP paths from Internal BGP neighbors"`
	TotalPaths          uint32 `DESCRIPTION: "Total number of paths received from neighbors"`
	TotalPrefixes       uint32 `DESCRIPTION: "Total number of destinations received from neighbors"`
}

const (
	PeerTypeInternal int8 = iota
	PeerTypeExternal
)

type BGPCounters struct {
	Update       uint64 `ACCESS:"", MULTIPLICITY:"*", DESCRIPTION: "Number of update messages"`
	Notification uint64 `ACCESS:"", MULTIPLICITY:"*", DESCRIPTION: "Number of notification messages"`
}

type BGPMessages struct {
	Sent     BGPCounters `ACCESS:"", MULTIPLICITY:"*", DESCRIPTION: "Tx counters of the BGP neighbor"`
	Received BGPCounters `ACCESS:"", MULTIPLICITY:"*", DESCRIPTION: "Rx counters of the BGP neighbor"`
}

type BGPQueues struct {
	Input  uint32 `ACCESS:"", MULTIPLICITY:"*", DESCRIPTION: "Input queue length of the BGP neighbor"`
	Output uint32 `ACCESS:"", MULTIPLICITY:"*", DESCRIPTION: "Output queue length of the BGP neighbor"`
}

type BGPNeighbor struct {
	baseObj
	PeerAS                  uint32 `DESCRIPTION: "Peer AS of the BGP neighbor"`
	LocalAS                 uint32 `DESCRIPTION: "Local AS of the BGP neighbor", DEFAULT: "0"`
	AuthPassword            string `DESCRIPTION: "Password to connect to the BGP neighbor", DEFAULT: ""`
	Description             string `DESCRIPTION: "Description of the BGP neighbor", DEFAULT: ""`
	NeighborAddress         string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Address of the BGP neighbor"`
	IfIndex                 int32  `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Interface of the BGP neighbor"`
	RouteReflectorClusterId uint32 `DESCRIPTION: "Cluster Id of the internal BGP neighbor route reflector client", DEFAULT: "0"`
	RouteReflectorClient    bool   `DESCRIPTION: "Set/Clear BGP neighbor as a route reflector client", DEFAULT: "false"`
	MultiHopEnable          bool   `DESCRIPTION: "Enable/Disable multi hop for BGP neighbor", DEFAULT: "false"`
	MultiHopTTL             uint8  `DESCRIPTION: "TTL for multi hop BGP neighbor", DEFAULT: "0"`
	ConnectRetryTime        uint32 `DESCRIPTION: "Connect retry time to connect to BGP neighbor after disconnect", DEFAULT: "60"`
	HoldTime                uint32 `DESCRIPTION: "Hold time for the BGP neighbor", DEFAULT: "180"`
	KeepaliveTime           uint32 `DESCRIPTION: "Keep alive time for the BGP neighbor", DEFAULT: "60"`
	AddPathsRx              bool   `DESCRIPTION: "Receive additional paths from BGP neighbor", DEFAULT: "false"`
	AddPathsMaxTx           uint8  `DESCRIPTION: "Max number of additional paths that can be transmitted to BGP neighbor", DEFAULT: "0"`
	PeerGroup               string `DESCRIPTION: "Peer group of the BGP neighbor", DEFAULT: ""`
	BfdEnable               bool   `DESCRIPTION: "Enable/Disable BFD for the BGP neighbor", DEFAULT: "false"`
	BfdSessionParam         string `DESCRIPTION: "Bfd session param name to be applied", DEFAULT: "default"`
	MaxPrefixes             uint32 `DESCRIPTION: "Maximum number of prefixes that can be received from the BGP neighbor", DEFAULT: "0"`
	MaxPrefixesThresholdPct uint8  `DESCRIPTION: "The percentage of maximum prefixes before we start logging", DEFAULT: "80"`
	MaxPrefixesDisconnect   bool   `DESCRIPTION: "Disconnect the BGP peer session when we receive the max prefixes from the neighbor", DEFAULT: "false"`
	MaxPrefixesRestartTimer uint8  `DESCRIPTION: "Time in seconds to wait before we start BGP peer session when we receive max prefixes", DEFAULT: "0"`
}

type BGPNeighborState struct {
	baseObj
	PeerAS                  uint32      `DESCRIPTION: "Peer AS of the BGP neighbor"`
	LocalAS                 uint32      `DESCRIPTION: "Local AS of the BGP neighbor"`
	PeerType                int8        `DESCRIPTION: "Type of the peer (internal/external)"`
	AuthPassword            string      `DESCRIPTION: "Password to connect to the BGP neighbor"`
	Description             string      `DESCRIPTION: "Description of the BGP neighbor"`
	NeighborAddress         string      `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Address of the BGP neighbor"`
	IfIndex                 int32       `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface of the BGP neighbor"`
	SessionState            uint32      `DESCRIPTION: "Session state of the BGP neighbor"`
	Messages                BGPMessages `DESCRIPTION: "Rx/Tx counter for BGP update and notification packets"`
	Queues                  BGPQueues   `DESCRIPTION: "Input/Output size of BGP packet queues"`
	RouteReflectorClusterId uint32      `DESCRIPTION: "Cluster Id of the internal BGP neighbor route reflector client"`
	RouteReflectorClient    bool        `DESCRIPTION: "Set/Clear BGP neighbor as a route reflector client"`
	MultiHopEnable          bool        `DESCRIPTION: "Enable/Disable multi hop for BGP neighbor"`
	MultiHopTTL             uint8       `DESCRIPTION: "TTL for multi hop BGP neighbor"`
	ConnectRetryTime        uint32      `DESCRIPTION: "Connect retry time to connect to BGP neighbor after disconnect"`
	HoldTime                uint32      `DESCRIPTION: "Hold time for the BGP neighbor"`
	KeepaliveTime           uint32      `DESCRIPTION: "Keep alive time for the BGP neighbor"`
	PeerGroup               string      `DESCRIPTION: "Peer group of the BGP neighbor"`
	BfdNeighborState        string      `DESCRIPTION: "BFD state of the BGP neighbor"`
	AddPathsRx              bool        `DESCRIPTION: "Receive additional paths from BGP neighbor"`
	AddPathsMaxTx           uint8       `DESCRIPTION: "Max number of additional paths that can be transmitted to BGP neighbor"`
	MaxPrefixes             uint32      `DESCRIPTION: "Maximum number of prefixes that can be received from the BGP neighbor"`
	MaxPrefixesThresholdPct uint8       `DESCRIPTION: "The percentage of maximum prefixes before we start logging"`
	MaxPrefixesDisconnect   bool        `DESCRIPTION: "Disconnect the BGP peer session when we receive the max prefixes from the neighbor"`
	MaxPrefixesRestartTimer uint8       `DESCRIPTION: "Time to wait before we start BGP peer session when we receive max prefixes"`
	TotalPrefixes           uint32      `DESCRIPTION: "Total number of prefixes received from the BGP neighbor"`
}

type BGPPeerGroup struct {
	baseObj
	PeerAS                  uint32 `DESCRIPTION: "Peer AS of the BGP neighbor"`
	LocalAS                 uint32 `DESCRIPTION: "Local AS of the BGP neighbor", DEFAULT: "0"`
	AuthPassword            string `DESCRIPTION: "Password to connect to the BGP neighbor", DEFAULT: ""`
	Description             string `DESCRIPTION: "Description of the BGP neighbor", DEFAULT: ""`
	Name                    string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP peer group"`
	RouteReflectorClusterId uint32 `DESCRIPTION: "Cluster Id of the internal BGP neighbor route reflector client", DEFAULT: "0"`
	RouteReflectorClient    bool   `DESCRIPTION: "Set/Clear BGP neighbor as a route reflector client", DEFAULT: "false"`
	MultiHopEnable          bool   `DESCRIPTION: "Enable/Disable multi hop for BGP neighbor", DEFAULT: "false"`
	MultiHopTTL             uint8  `DESCRIPTION: "TTL for multi hop BGP neighbor", DEFAULT: "0"`
	ConnectRetryTime        uint32 `DESCRIPTION: "Connect retry time to connect to BGP neighbor after disconnect", DEFAULT: "60"`
	HoldTime                uint32 `DESCRIPTION: "Hold time for the BGP neighbor", DEFAULT: "180"`
	KeepaliveTime           uint32 `DESCRIPTION: "Keep alive time for the BGP neighbor", DEFAULT: "60"`
	AddPathsRx              bool   `DESCRIPTION: "Receive additional paths from BGP neighbor", DEFAULT: "false"`
	AddPathsMaxTx           uint8  `DESCRIPTION: "Max number of additional paths that can be transmitted to BGP neighbor", DEFAULT: "0"`
	MaxPrefixes             uint32 `DESCRIPTION: "Maximum number of prefixes that can be received from the BGP neighbor", DEFAULT: "0"`
	MaxPrefixesThresholdPct uint8  `DESCRIPTION: "The percentage of maximum prefixes before we start logging", DEFAULT: "0"`
	MaxPrefixesDisconnect   bool   `DESCRIPTION: "Disconnect the BGP peer session when we receive the max prefixes from the neighbor", DEFAULT: "false"`
	MaxPrefixesRestartTimer uint8  `DESCRIPTION: "Time to wait before we start BGP peer session when we receive max prefixes", DEFAULT: "0"`
}

type BGPRouteState struct {
	baseObj
	Network         string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Network address of the destination"`
	CIDRLen         uint16   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "CIDR length of the destination"`
	NextHop         string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Next hop address for the destination"`
	Metric          uint32   `DESCRIPTION: "MED of the path to the destination"`
	LocalPref       uint32   `DESCRIPTION: "Local preference of the path to the destination"`
	Path            []string `DESCRIPTION: "AS path to the destination"`
	UpdatedTime     string   `DESCRIPTION: "Last time the destination was updated"`
	UpdatedDuration string   `DESCRIPTION: "Time since the destination was last updated"`
	PathId          uint32   `DESCRIPTION: "Path id of the path"`
}

type BGPPolicyCondition struct {
	baseObj
	Name            string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy condition"`
	ConditionType   string `DESCRIPTION: "Type of the BGP policy condition. "`
	IpPrefix        string `DESCRIPTION: "IP adddress to match in CIDR format"`
	MaskLengthRange string `DESCRIPTION: "IP address mask lenght range to match"`
}

type BGPPolicyConditionState struct {
	baseObj
	Name           string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy condition"`
	ConditionInfo  string   `DESCRIPTION: "Description of the BGP policy condition"`
	PolicyStmtList []string `DESCRIPTION: "Policy statements that use the BGP policy condition"`
}

type BGPPolicyAction struct {
	baseObj
	Name            string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy action"`
	ActionType      string `DESCRIPTION: "Type of the BGP policy action"`
	GenerateASSet   bool   `DESCRIPTION: "Enable/Disable generating AS set for BGP aggregate action"`
	SendSummaryOnly bool   `DESCRIPTION: "Enable/Disable sending summary only for BGP aggregate action"`
}

type BGPPolicyActionState struct {
	baseObj
	Name           string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy action"`
	ActionInfo     string   `DESCRIPTION: "Description of the BGP policy action"`
	PolicyStmtList []string `DESCRIPTION: "Policy statements that use the BGP policy action"`
}

type BGPPolicyStmt struct {
	baseObj
	Name            string   `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy statement"`
	MatchConditions string   `DESCRIPTION: "Match conditions all/any"`
	Conditions      []string `DESCRIPTION: "List of conditions"`
	Actions         []string `DESCRIPTION: "List of actions"`
}

type BGPPolicyStmtState struct {
	baseObj
	Name            string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy statement"`
	MatchConditions string   `DESCRIPTION: "Match conditions ,SELECTION: All/Any"`
	Conditions      []string `DESCRIPTION: "List of conditions"`
	Actions         []string `DESCRIPTION: "List of actions"`
}

type BGPPolicyDefinitionStmtPrecedence struct {
	Precedence int32  `ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Precedence of the BGP policy statement"`
	Statement  string `ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy statement"`
}

type BGPPolicyDefinition struct {
	baseObj
	Name          string                              `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy definition"`
	Precedence    int32                               `DESCRIPTION: "Precedence of the policy definition"`
	MatchType     string                              `DESCRIPTION: "Match type for policy definition  ,SELECTION: All/Any"`
	StatementList []BGPPolicyDefinitionStmtPrecedence `DESCRIPTION: "Precedence of statements in the policy"`
}

type BGPPolicyDefinitionState struct {
	baseObj
	Name         string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Name of the BGP policy definition"`
	HitCounter   int32    `DESCRIPTION: "Number of matches for this policy"`
	IpPrefixList []string `DESCRIPTION: "IP addresses that matched the policy"`
}
