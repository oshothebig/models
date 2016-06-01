//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __  
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  | 
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  | 
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   | 
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  | 
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__| 
//                                                                                                           

package models

type NextHopInfo struct {
	NextHopIp     string `DESCRIPTION: "next hop ip of the route, DEFAULT:"0.0.0.0""`
	NextHopIntRef string `DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this next hop is configured", OPTIONAL`
	Weight        int32  `DESCRIPTION : "Weight of the next hop",DEFAULT:0, RANGE:0-31, OPTIONAL `
}

/*
type IPv4Route struct {
	baseObj
	DestinationNw     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route, can be specified either in CIDR notation or as a IP address."`
	NetworkMask       string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route, required to be non-zero in case DestinationNw is a non-CIDR address."`
	NextHopIp         string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "next hop ip of the route"`
	Cost              uint32 `DESCRIPTION :"Cost of this route", DEFAULT:0`
	OutgoingIntfType  string `DESCRIPTION :"Interface type of the next hop interface"`
	OutgoingInterface string `DESCRIPTION :"Interface ID of the next hop interface"`
	Protocol          string `DESCRIPTION :"Protocol type of the route", DEFAULT:"STATIC"`
	Weight            int32  `DESCRIPTION : "Weight of the next hop", DEFAULT:0, RANGE:0-31`
}

/*type IPv4Route struct {
	baseObj
	DestinationNw     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route"`
	NetworkMask       string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route"`
	Protocol           string   `DESCRIPTION :"Protocol type of the route"`
	NextHop           []NextHopInfo
}*/
type IPv4Route struct {
	baseObj
	DestinationNw string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route"`
	NetworkMask   string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route"`
	Protocol      string `DESCRIPTION :"Protocol type of the route", OPTIONAL, DEFAULT:"STATIC"`
	Cost          uint32 `DESCRIPTION :"Cost of this route", OPTIONAL, DEFAULT:0`
	NullRoute     bool   `DESCRIPTION : "Specify if this is a null route", OPTIONAL, DEFAULT:false`
	NextHop       []NextHopInfo
}
type IPv4RouteState struct {
	baseObj
	DestinationNw      string        `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route", USESTATEDB:"true"`
	Protocol           string        `DESCRIPTION :"Protocol type of the route"`
	IsNetworkReachable bool          `DESCRIPTION :"Indicates whether this network is reachable"`
	RouteCreatedTime   string        `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime   string        `DESCRIPTION :"Time when the route was last updated"`
	NextHopList        []NextHopInfo `DESCRIPTION: "List of next hops to reach this network"`
	PolicyList         []string      `DESCRIPTION :"List of policies applied on this route"`
}
type IPv4EventState struct {
	baseObj
	Index     uint32 `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Event ID"`
	TimeStamp string `DESCRIPTION :"Time when the event occured"`
	EventInfo string `DESCRIPTION :"Detailed description of the event"`
}

type PolicyCondition struct {
	baseObj
	Name            string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "PolicyConditionName"`
	ConditionType   string `DESCRIPTION: "Specifies the match criterion this condition defines", SELECTION: "MatchProtocol"/"MatchDstIpPrefix"/"MatchSrcIpPrefix"`
	Protocol        string `DESCRIPTION: "Protocol to match on if the ConditionType is set to MatchProtocol",SELECTION:"CONNECTED"/"STATIC"/"OSPF"/"BGP"`
	IpPrefix        string `DESCRIPTION: "Used in conjunction with MaskLengthRange to specify the IP Prefix to match on when the ConditionType is MatchDstIpPrefix/MatchSrcIpPrefix."`
	MaskLengthRange string `DESCRIPTION: "Used in conjuction with IpPrefix to specify specify the IP Prefix to match on when the ConditionType is MatchDstIpPrefix/MatchSrcIpPrefix."`
}
type PolicyConditionState struct {
	baseObj
	Name           string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Condition name"`
	ConditionInfo  string
	PolicyStmtList []string `DESCRIPTION: "List of policy statements using this condition"`
}

type PolicyStmt struct {
	baseObj
	Name            string   `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Statement Name"`
	MatchConditions string   `DESCRIPTION :"Specifies whether to match all/any of the conditions of this policy statement",SELECTION:"any"/"all",DEFAULT:"all"`
	Conditions      []string `DESCRIPTION :"List of conditions added to this policy statement"`
	Action          string   `DESCRIPTION :"Action for this policy statement", SELECTION:"permit"/"deny",DEFAULT: "deny"`
}
type PolicyStmtState struct {
	baseObj
	Name            string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyStmtState"`
	MatchConditions string   `DESCRIPTION :"Specifies whether to match all/any of the conditions of this policy statement"`
	Conditions      []string `DESCRIPTION :"List of conditions added to this policy statement"`
	Action          string   `DESCRIPTION :"Action corresponding to this policy statement"`
	PolicyList      []string `DESCRIPTION :"List of policies using this policy statement"`
}
type PolicyDefinitionStmtPriority struct {
	Priority  int32
	Statement string
}
type PolicyDefinition struct {
	baseObj
	Name          string                         `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Name"`
	Priority      int32                          `DESCRIPTION :"Priority of the policy w.r.t other policies configured",RANGE:0-255`
	MatchType     string                         `DESCRIPTION :"Specifies whether to match all/any of the statements within this policy",SELECTION:"all"/"any",DEFAULT:"all"`
	PolicyType    string                         `DESCRIPTION : Specifies the intended protocol application for the policy", SELECTION: "BGP"/"OSPF"/"ALL", DEFAULT:"ALL"`
	StatementList []PolicyDefinitionStmtPriority `DESCRIPTION :"Specifies list of statements along with their precedence order."`
}
type PolicyDefinitionState struct {
	baseObj
	Name         string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyDefinitionState"`
	IpPrefixList []string `DESCRIPTION :"List of networks/IP Prefixes this policy has been applied on to."`
}

type RouteDistanceState struct {
	baseObj
	Protocol string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "RouteDistanceState protocol"`
	Distance int32  `DESCRIPTION: "The current value of the admin distance of this protocol"`
}
