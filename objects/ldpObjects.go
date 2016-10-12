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
//  _______  __       __________   ___      _______ ____    __    ____  __   ___________   ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

type LdpGlobal struct {
	baseObj
	Vrf             string `SNAPROUTE:"KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE:"true", DEFAULT:"default", DESCRIPTION:"VRF where LDP is globally enabled or disabled"`
	Enable          bool   `DESCRIPTION:"Global LDP state in this VRF", DEFAULT:"true"`
	RouterId        string `DESCRIPTION:"LDP Router ID", DEFAULT:"0.0.0.0"`
	GrEnable        bool   `DESCRIPTION:"LDP GR capability", DEFAULT:"false"`
	ReconnectTime   uint32 `DESCRIPTION:"LDP session reconnection timer (s)", MIN:"3", MAX:"3600", DEFAULT:"300"`
	RecoveryTime    uint32 `DESCRIPTION:"LSP recovery timer (s)", MIN:"3", MAX:"3600", DEFAULT:"300"`
	PeerLiveTime    uint32 `DESCRIPTION:"The neighbor keeplive timer (s)", MIN:"3", MAX:"3600", DEFAULT:"600"`
	BackOffInitTime uint32 `DESCRIPTION:"Init value of the exponential backoff timer (s)", MIN:"5", MAX:"2147483", DEFAULT:"15"`
	BackOffMaxTime  uint32 `DESCRIPTION:"Maximum value of the exponential backoff timer (s)", MIN:"5", MAX:"2147483", DEFAULT:"120"`
}

type LdpGlobalState struct {
	baseObj
	Vrf             string `SNAPROUTE:"KEY", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION:"VRF where LDP is globally enabled or disabled"`
	Enable          bool   `DESCRIPTION: "Global LDP state in this VRF"`
	RouterId        string `DESCRIPTION: "LDP Router ID"`
	GrEnable        bool   `DESCRIPTION: "LDP GR capability"`
	ReconnectTime   uint32 `DESCRIPTION: "LDP session reconnection timer (s)"`
	RecoveryTime    uint32 `DESCRIPTION: "LSP recovery timer (s)"`
	PeerLiveTime    uint32 `DESCRIPTION: "The neighbor keeplive timer (s)"`
	BackOffInitTime uint32 `DESCRIPTION: "Init value of the exponential backoff timer (s)"`
	BackOffMaxTime  uint32 `DESCRIPTION: "Maximum value of the exponential backoff timer (s)"`
}

type LdpIntf struct {
	baseObj
	IntfRef           string `SNAPROUTE:"KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION:"Interface reference"`
	Vrf               string `DESCRIPTION:"VRF where this interface exists", DEFAULT:"default"`
	HelloSendInterval uint16 `DESCRIPTION:"Hello packet sending interval (s)", MIN:"1", MAX:"65535", DEFAULT:"5"`
	HelloHoldTime     uint16 `DESCRIPTION:"Hello hold time (s)", MIN:"3", MAX:"65535", DEFAULT:"15"`
	KeepaliveSendTime uint16 `DESCRIPTION:"Keepalive packet sending timer (s)", MIN:"1", MAX:"65535", DEFAULT:"10"`
	KeepaliveHoldTime uint16 `DESCRIPTION:"Keepalive packet holding timer (s)", MIN:"30", MAX:"65535", DEFAULT:"30"`
	IgpSyncDelayTime  uint32 `DESCRIPTION:"Interval at which an interface waits to establish an LSP after an LDP session is set up (s), MIN:"0", MAX:"65535" DEFAULT:"0"`
}

type LdpIntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE:"KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Interface reference"`
	Vrf               string `DESCRIPTION:"VRF where this interface exists"`
	HelloSendInterval uint16 `DESCRIPTION:"Hello packet sending interval (s)"`
	HelloHoldTime     uint16 `DESCRIPTION:"Hello hold time (s)"`
	KeepaliveSendTime uint16 `DESCRIPTION:"Keepalive packet sending timer (s)"`
	KeepaliveHoldTime uint16 `DESCRIPTION:"Keepalive packet holding timer (s)"`
	IgpSyncDelayTime  uint32 `DESCRIPTION:"Interval at which an interface waits to establish an LSP after an LDP session is set up (s)`
}

type LdpPeer struct {
	baseObj
	IpAddr            string `SNAPROUTE:"KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION:"Remote peer IP address"`
	Vrf               string `DESCRIPTION:"Name of an LDP instance", DEFAULT:"default"`
	PeerName          string `DESCRIPTION:"Name of a remote neighbor", DEFAULT:""`
	HelloSendInterval uint16 `DESCRIPTION:"Hello packet sending interval (s)", MIN:"1", MAX:"65535", DEFAULT:"15"`
	HelloHoldTime     uint16 `DESCRIPTION:"Hello hold time (s)", MIN: "3", MAX:"65535", DEFAULT:"45"`
	KeepaliveSendTime uint16 `DESCRIPTION:"Keepalive packet sending timer (s)", MIN:"1", MAX:"65535", DEFAULT:"10"`
	KeepaliveHoldTime uint16 `DESCRIPTION:"Keepalive packet holding timer (s)", MIN:"30", MAX:"65535", DEFAULT:"30"`
	IgpSyncDelayTime  uint32 `DESCRIPTION:"Interval at which an interface waits to establish an LSP after an LDP session is set up (s), MIN:"0", MAX:"65535" DEFAULT:"0"`
}

type LdpPeerState struct {
	baseObj
	IpAddr            string `SNAPROUTE:"KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Remote peer IP address"`
	Vrf               string `DESCRIPTION:"Name of an LDP instance"`
	PeerName          string `DESCRIPTION:"Name of a remote neighbor"`
	HelloSendInterval uint16 `DESCRIPTION:"Hello packet sending interval (s)"`
	HelloHoldTime     uint16 `DESCRIPTION:"Hello hold time (s)"`
	KeepaliveSendTime uint16 `DESCRIPTION:"Keepalive packet sending timer (s)"`
	KeepaliveHoldTime uint16 `DESCRIPTION:"Keepalive packet holding timer (s)"`
	IgpSyncDelayTime  uint32 `DESCRIPTION:"Interval at which an interface waits to establish an LSP after an LDP session is set up (s)`
}
