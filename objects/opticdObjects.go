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
	ModuleId            uint8 `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ModuleReset         bool  `DESCRIPTION: "Reset state of this dwdm module (false (Reset deasserted), true (Reset asserted))", DEFAULT:false`
	IndependentLaneMode bool  `DESCRIPTION: "Network lane configuration for the DWDM Module. true-Independent lanes, false-Coupled lanes, DEFAULT:true"`
}

type DWDMModuleNwIntf struct {
	baseObj
	ModuleId      uint8  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId      uint8  `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module network interface identifier"`
	ModulationFmt string `DESCRIPTION: "Modulation format to use for this network interface", SELECTION: "QPSK"/"8QAM/"16QAM"`
	TxPower       int16  `DESCRIPTION: "Transmit output power for this network interface in dBm, MIN:0, MAX:4294967295"`
	WaveLength    uint16 `DESCRIPTION: "The ITU-T G.694.1 grid wavelength value to use for this network interface in nm", MIN:1530, MAX:1565`
}

type DWDMModuleNwIntfState struct {
	baseObj
	ModuleId          uint8  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId          uint8  `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module network interface identifier"`
	TxChanGridSpacing string `DESCRIPTION: "The channel grid spacing used for this network interface in GHz"`
}

/*
type DWDMModuleClntIntf struct {
	baseObj
	ModuleId   uint8 `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ClntIntfId uint8 `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module client interface identifier"`
}

type DWDMModuleClntIntfState struct {
	baseObj
	ModuleId   uint8 `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ClntIntfId uint8 `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module client interface identifier"`
}
*/
