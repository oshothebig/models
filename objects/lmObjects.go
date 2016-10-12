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

package objects

type MplsAction struct {
	ActionType string `DESCRIPTION: "MPLS action type", STRLEN:"10" SELECTION: POP/SWAP/SWAP_PUSH/PUSH`
}

type NextHopLfe struct {
	baseObj
	NhLFEIdx   int32        `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Label index"`
	Vrf        string       `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	IntfRef    string       `DESCRIPTION :"Next hop mpls interface index"`
	LabelMap   int32        `DESCRIPTION :"lable map for ILM entry into ILM table"`
	ActionList []MplsAction `DESCRIPTION :"label action list"`
}

type NextHopLfeState struct {
	baseObj
	NhLFEIdx   int32        `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Label index"`
	Vrf        string       `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	IntfRef    string       `DESCRIPTION :"Next hop mpls interface index"`
	LabelMap   int32        `DESCRIPTION :"lable map for ILM entry into ILM table"`
	ActionList []MplsAction `DESCRIPTION :"label action list"`
}

type FtnEntry struct {
	baseObj
	IpAddr   string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "IP address prefix IP/Net mask in CIDR format", STRLEN:"18"`
	Vrf      string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	NhLFEIdx NextHopLfe `DESCRIPTION :"NH LFE for this FTN Entry"`
}

type FtnEntryState struct {
	baseObj
	IpAddr       string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "IP address prefix IP/Net mask in CIDR format", STRLEN:"18"`
	Vrf          string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	NhLFEIdx     NextHopLfe `DESCRIPTION :"NH LFE for this FTN Entry"`
	RibQualified string     `DESCRIPTION :"Installed into RIB YES/NO", STRLEN:"3" SELECTION: Yes/No`
}

type MplsLabel struct {
	baseObj
	LabelIdx int32      `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Label index"`
	Vrf      string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	Protocol string     `DESCRIPTION :"Protocol that learned or created the label", OPTIONAL, DEFAULT:"STATIC"`
	NhLFEIdx NextHopLfe `DESCRIPTION :"NH LFE for this label"`
}

type MplsLabelState struct {
	baseObj
	LabelIdx      int32      `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Label index"`
	Vrf           string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	Protocol      string     `DESCRIPTION :"Protocol that learned or created the label", OPTIONAL, DEFAULT:"STATIC"`
	NhLFEIdx      NextHopLfe `DESCRIPTION :"NH LFE for this label"`
	LFibQualified string     `DESCRIPTION :"Installed into LFIB YES/NO", STRLEN:"3" SELECTION: Yes/No`
}

type MplsIntf struct {
	IntfRef string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this IPv4 object is configured"`
	Vrf     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
}

type MplsGlobal struct {
	baseObj
	Vrf        string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	MplsIntf   []MplsIntf `DESCRIPTION :"Interfaces enabled with MPLS"`
	AdminState string     `DESCRIPTION :"Admin state of Router MPLS", OPTIONAL, DEFAULT:"UP"`
}

type MplsGlobalState struct {
	baseObj
	Vrf        string     `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "Routing and Forwarding context", DEFAULT:"default"`
	MplsIntf   []MplsIntf `DESCRIPTION :"Interfaces enabled with MPLS"`
	AdminState string     `DESCRIPTION :"Admin state of Router MPLS", OPTIONAL, DEFAULT:"UP"`
}

type MplsLsp struct {
	baseObj
	Vrf        int32  `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Routing and Forwarding context"`
	IngressIP  string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "Ingress IP of the LSP"`
	EgressIP   string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "Egress IP of the LSP"`
	AdminState string `DESCRIPTION :"Admin state of LSP", OPTIONAL, DEFAULT:"UP"`
}

type MplsLspState struct {
	baseObj
	Vrf       int32  `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Routing and Forwarding context"`
	IngressIP string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "Ingress IP of the LSP"`
	EgressIP  string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "Egress IP of the LSP"`
	OperState string `DESCRIPTION: "Operational state of LSP"`
}

type LabelAlloc struct {
	baseObj
	Vrf      int32  `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Routing and Forwarding context"`
	IpAddr   string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", ACCELERATED: "true", DESCRIPTION: "IP address prefix IP/Net mask in CIDR format", STRLEN:"18"`
	Protocol string `DESCRIPTION :"Protocol that is requesting the label"`
}

//
//type LabelBindReq struct {
//	baseObj
//	IpAddress string
//	IpMask    string
//	VrfId     int32
//	Protocol  string
//}
