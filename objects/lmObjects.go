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
	ActionType string `DESCRIPTION: "MPLS action type. POP, SWAP and SWAP_PUSH"`
}

type FtnEntry struct {
	IpAddress string `DESCRIPTION :"IP address for the FTN map"`
	IpMask    string `DESCRIPTION :"IP address mask for the FTN map"`
	VrfId     int32  `DESCRIPTION :"VRF ID"`
}

type NextHopLfe struct {
	baseObj
	NhLFEIdx   int32        `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Label index"`
	Vrf        int32        `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Routing and Forwarding context"`
	IntfIdx    int32        `DESCRIPTION :"Next hop mpls interface index"`
	LabelMap   int32        `DESCRIPTION :"lable map for ILM entry into ILM table"`
	FtnMap     FtnEntry     `DESCRIPTION :"FEC(IP) Map for this NH LFE"`
	ActionList []MplsAction `DESCRIPTION :"label action list"`
}

type Label struct {
	baseObj
	LabelIdx int32      `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Label index"`
	Vrf      int32      `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Routing and Forwarding context"`
	Protocol string     `DESCRIPTION :"Protocol that learned or created the label", OPTIONAL, DEFAULT:"STATIC"`
	NhLFEIdx NextHopLfe `DESCRIPTION :"NH LFE for this label"`
}

type LabelAlloc struct {
	IpAddress string `DESCRIPTION :"IP address"`
	IpMask    string `DESCRIPTION :"IP mask"`
	VrfId     int32  `DESCRIPTION :"VRF ID"`
	Protocol  string `DESCRIPTION :"Protocol that is requesting the label"`
}

type MplsIntf struct {
	baseObj
	IntfIdx    int32  `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: ""Interface index of the interface enabled with MPLS"`
	Vrf        int32  `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "Routing and Forwarding context"`
	AdminState string `DESCRIPTION: "Current Admin state of MPLS for this interface"`
}

type MplsGlobal struct {
	ConfigObj
	MplsIntf   []MplsIntf
	AdminState string
}

type MplsLsp struct {
	ConfigObj
	Vrf       int32
	IngressIP string
	EgressIP  string
}

//
//type LabelBindReq struct {
//	baseObj
//	IpAddress string
//	IpMask    string
//	VrfId     int32
//	Protocol  string
//}
