package models

type NextHopInfo struct {
	NextHopIp          string   `DESCRIPTION: "next hop ip of the route"`
	OutgoingIntfType   string   `DESCRIPTION :"Interface type of the next hop interface"`
	OutgoingInterface  string   `DESCRIPTION :"Interface ID of the next hop interface"`
	Protocol           string   `DESCRIPTION :"Protocol type of the route"`
}
type IPv4Route struct {
	ConfigObj
	DestinationNw     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route"`
	NetworkMask       string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route"`
	NextHopIp         string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "next hop ip of the route"`
	Cost              uint32 `DESCRIPTION :"Cost of this route", DEFAULT:0`
	OutgoingIntfType  string `DESCRIPTION :"Interface type of the next hop interface"`
	OutgoingInterface string `DESCRIPTION :"Interface ID of the next hop interface"`
	Protocol           string   `DESCRIPTION :"Protocol type of the route"`
}
/*type IPv4RouteState struct {
	ConfigObj
	DestinationNw      string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route", USESTATEDB:"true"`
	NextHopIp          string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "next hop ip of the route"`
	OutgoingIntfType   string   `DESCRIPTION :"Interface type of the next hop interface"`
	OutgoingInterface  string   `DESCRIPTION :"Interface ID of the next hop interface"`
	Protocol           string   `DESCRIPTION :"Protocol type of the route"`
	PolicyList         []string `DESCRIPTION :"List of policies applied on this route"`
	IsNetworkReachable bool     `DESCRIPTION :"Indicates whether this network is reachable"`
	RouteCreatedTime   string   `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime   string   `DESCRIPTION :"Time when the route was last updated"`
}*/
/*type IPv4Route struct {
	ConfigObj
	DestinationNw     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route"`
	NetworkMask       string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route"`
	Protocol           string   `DESCRIPTION :"Protocol type of the route"`
	NextHop           []NextHopInfo
}*/
type IPv4RouteState struct {
	ConfigObj
	DestinationNw      string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route", USESTATEDB:"true"`
	Protocol           string   `DESCRIPTION :"Protocol type of the route"`
	IsNetworkReachable bool     `DESCRIPTION :"Indicates whether this network is reachable"`
	RouteCreatedTime   string   `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime   string   `DESCRIPTION :"Time when the route was last updated"`
	NextHopList        []NextHopInfo `DESCRIPTION: "List of next hops to reach this network"`
	PolicyList         []string `DESCRIPTION :"List of policies applied on this route"`
}
type IPv4EventState struct {
	ConfigObj
	Index     uint32 `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Event ID"`
	TimeStamp string `DESCRIPTION :"Time when the event occured"`
	EventInfo string `DESCRIPTION :"Detailed description of the event"`
}

type PolicyCondition struct {
	ConfigObj
	Name            string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "PolicyConditionName"`
	ConditionType   string `DESCRIPTION: "Specifies the match criterion this condition defines - eg: MatchProtocol/MatchDstIpPrefix"`
	MatchProtocol   string `DESCRIPTION: "Protocol to match on if the ConditionType is set to MatchProtocol"`
	IpPrefix        string `DESCRIPTION: "Used in conjunction with MaskLengthRange to specify the IP Prefix to match on when the ConditionType is MatchDstIpPrefix/MatchSrcIpPrefix."`
	MaskLengthRange string `DESCRIPTION: "Used in conjuction with IpPrefix to specify specify the IP Prefix to match on when the ConditionType is MatchDstIpPrefix/MatchSrcIpPrefix."`
}
type PolicyConditionState struct {
	ConfigObj
	Name           string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Condition name"`
	ConditionInfo  string
	PolicyStmtList []string `DESCRIPTION: "List of policy statements using this condition"`
}

type PolicyAction struct {
	ConfigObj
	Name                           string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "PolicyActionName"`
	ActionType                     string `DESCRIPTION: "Specifies the type of the action  - eg: RouteDisposition/NetworkStatementAdvertise/Redistribution/SetAdminDistance"`
	SetAdminDistanceValue          int32  `DESCRIPTION :"Specifies the value of the admin distance/protocol preference when the action type is SetAdminDistance"`
	Accept                         bool   `DESCRIPTION :"When set to true, along with action type RouteDisposition, indicates to accept the policy entity"`
	Reject                         bool   `DESCRIPTION :"When set to true, along with action type RouteDisposition, indicates to reject the policy entity"`
	RedistributeAction             string `DESCRIPTION :"Used in conjuction with RedistributeTargetProtocol for action type Redistribute, indicates to allow/block redistribution"`
	RedistributeTargetProtocol     string `DESCRIPTION :"Used in conjuction with RedistributeAction for action type Redistribute, indicates the target protocol for redistribution"`
	NetworkStatementTargetProtocol string `DESCRIPTION :"Used for action type NetworkStatementAdvertise, indicates the target protocol for Network Statement Advertisement"`
}

type PolicyActionState struct {
	ConfigObj
	Name           string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Action name"`
	ActionInfo     string
	PolicyStmtList []string `DESCRIPTION: "List of policy statements using this condition"`
}
type PolicyStmt struct {
	ConfigObj
	Name            string   `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Statement Name"`
	MatchConditions string   `DESCRIPTION :"Specifies whether to match all/any of the conditions of this policy statement"`
	Conditions      []string `DESCRIPTION :"List of conditions added to this policy statement"`
	Actions         []string `DESCRIPTION :"List of actions added to this policy statement"`
}
type PolicyStmtState struct {
	ConfigObj
	Name            string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyStmtState"`
	MatchConditions string   `DESCRIPTION :"Specifies whether to match all/any of the conditions of this policy statement"`
	Conditions      []string `DESCRIPTION :"List of conditions added to this policy statement"`
	Actions         []string `DESCRIPTION :"List of actions added to this policy statement"`
	PolicyList      []string `DESCRIPTION :"List of policies using this policy statement"`
}
type PolicyDefinitionStmtPrecedence struct {
	Precedence int32
	Statement  string
}
type PolicyDefinition struct {
	ConfigObj
	Name          string                           `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Name"`
	Precedence    int32                            `DESCRIPTION :"Priority of the policy w.r.t other policies configured"`
	MatchType     string                           `DESCRIPTION :"Specifies whether to match all/any of the statements within this policy"`
	StatementList []PolicyDefinitionStmtPrecedence `DESCRIPTION :"Specifies list of statements along with their precedence order."`
}
type PolicyDefinitionState struct {
	ConfigObj
	Name         string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyDefinitionState"`
	HitCounter   int32    `DESCRIPTION :"Number of times this policy has been applied"`
	IpPrefixList []string `DESCRIPTION :"List of networks/IP Prefixes this policy has been applied on to."`
}

type RouteDistanceState struct {
	ConfigObj
	Protocol string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "RouteDistanceState protocol"`
	Distance int32  `DESCRIPTION: "The current value of the admin distance of this protocol"`
}
