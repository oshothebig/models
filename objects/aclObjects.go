//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

type Acl struct {
	baseObj
	AclName  string    `SNAPROUTE: "KEY", ACCESS:"w",MULTIPLICITY: "*", DESCRIPTION: "Acl name to be used to refer to this ACL"`
	RuleList []AclRule `DESCRIPTION: "List of rules to be applied to this ACL"`
}

type AclRule struct {
	baseObj
	RuleName   string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"w", DESCRIPTION: "Acl rule name"`
	SourceMac  string `SNAPROUTE: "KEY",  DESCRIPTION: "Source MAC address."`
	DestMac    string `DESCRIPTION: "Destination MAC address"`
	SourceIp   string `DESCRIPTION: "Source IP address"`
	DestIp     string `DESCRIPTION: "Destination IP address"`
	SourceMask string `DESCRIPTION: "Network mask for source IP"`
	DestMask   string `DESCRIPTION: "Network mark for dest IP"`
	Action     string `DESCRIPTION: "Type of action (Allow/Deny)", DEFAULT:"Allow", STRLEN:"16"`
	Proto      string `DESCRIPTION: "Protocol type"`
}

type AclState struct {
	baseObj
	AclName  string    `SNAPROUTE: "KEY", ACCESS:"r",MULTIPLICITY: "*", DESCRIPTION: "Acl name to be used to refer to this ACL"`
	RuleList []AclRule `DESCRIPTION: "List of rules to be applied to this ACL"`
}

type AclRuleState struct {
	baseObj
	RuleName   string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION: "Acl rule name"`
	SourceMac  string `SNAPROUTE: "KEY",  DESCRIPTION: "Source MAC address."`
	DestMac    string `DESCRIPTION: "Destination MAC address"`
	SourceIp   string `DESCRIPTION: "Source IP address"`
	DestIp     string `DESCRIPTION: "Destination IP address"`
	SourceMask string `DESCRIPTION: "Network mask for source IP"`
	DestMask   string `DESCRIPTION: "Network mark for dest IP"`
	Action     string `DESCRIPTION: "Type of action (Allow/Deny)", DEFAULT:"Allow", STRLEN:"16"`
	Proto      string `DESCRIPTION: "Protocol type"`
}
