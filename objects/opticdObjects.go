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

type DWDMModuleState struct {
	baseObj
	ModuleId         uint8   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ModuleState      string  `DESCRIPTION: "Current MSA state of dwdm module"`
	ModuleVoltage    float64 `DESCRIPTION: "Module power supply voltage in Volts"`
	ModuleTemp       float64 `DESCRIPTION: "Module temperature in deg Celsius"`
	Populated        bool    `DESCRIPTION: "Is module populated"`
	VendorName       string  `DESCRIPTION: "Vendor name of dwdm module"`
	VendorPartNum    string  `DESCRIPTION: "Vendor assigned part number of dwdm module"`
	VendorSerialNum  string  `DESCRIPTION: "Vendor assigned serial number of dwdm module "`
	VendorDateCode   string  `DESCRIPTION: "Device manufacture data code of dwdm module"`
	ModuleHWVersion  string  `DESCRIPTION: "HW version of dwdm module"`
	ModuleFWAVersion string  `DESCRIPTION: "Firmware A version of dwdm module"`
	ModuleFWBVersion string  `DESCRIPTION: "Firmware B version of dwdm module"`
}

type DWDMModule struct {
	baseObj
	ModuleId   uint8 `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	AdminState uint8 `DESCRIPTION: Administrative state of this dwdm module (0 Disabled, 1 Enabled)", DEFAULT:0`
}
