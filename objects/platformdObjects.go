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

/*
 * This DS will be used while Created/Deleting Platform Config
 */
type PlatformSystemState struct {
	baseObj
	ObjName   string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "ObjName", DEFAULT: "System"`
	SerialNum string `DESCRIPTION: "Serial Number"`
}

type Fan struct {
	baseObj
	FanId          int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Fan unit id", DEFAULT:0`
	AdminSpeed     int32  `DESCRIPTION: "Fan admin speed in rpm"`
	AdminDirection string `DESCRIPTION: "Air flow caused because of fan rotation", SELECTION: B2F/F2B, DEFAULT: B2F"`
}

type FanState struct {
	baseObj
	FanId         int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fan unit id", DEFAULT:0`
	OperMode      string `DESCRIPTION: "Operational state of Fan", SELECTION: ON/OFF`
	OperSpeed     int32  `DESCRIPTION: "Fan operational speed in rpm"`
	OperDirection string `DESCRIPTION: "Air flow caused because of fan rotation", SELECTION: B2F/F2B"`
	Status        string `DESCRIPTION: "Fan status PRESENT/MISSING/FAILED/NORMAL"`
	Model         string `DESCRIPTION: "Model of Fan"`
	SerialNum     string `DESCRIPTION: "Serial Number"`
}
