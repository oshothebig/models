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
	ModuleId            uint8 `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER: "true", DESCRIPTION: "DWDM Module identifier"`
	AdminState          bool  `DESCRIPTION: "Reset state of this dwdm module (false (Reset deasserted), true (Reset asserted))", DEFAULT:true`
	IndependentLaneMode bool  `DESCRIPTION: "Network lane configuration for the DWDM Module. true-Independent lanes, false-Coupled lanes, DEFAULT:true"`
}

type DWDMModuleNwIntf struct {
	baseObj
	ModuleId                uint8   `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER: "true", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId                uint8   `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module network interface identifier"`
	ModulationFmt           string  `DESCRIPTION: "Modulation format to use for this network interface", SELECTION: "QPSK"/"8QAM/"16QAM", DEFAULT:"16QAM"`
	TxPower                 float64 `DESCRIPTION: "Transmit output power for this network interface in dBm, MIN:0, MAX:4294967295", DEFAULT:0`
	WaveLength              uint16  `DESCRIPTION: "The ITU-T G.694.1 grid wavelength value to use for this network interface in nm", MIN:1530, MAX:1565`
	FECMode                 string  `DESCRIPTION: "DWDM Module network interface FEC mode", SELECTION: "15%SDFEC"/"15%OvrHeadSDFEC/25%OvrHeadSDFEC"`
	DiffEncoding            bool    `DESCRIPTION: "Control to enable/disable DWDM Module network interface encoding type", DEFAULT: true`
	TxPulseShapeFltrType    string  `DESCRIPTION: "TX pulse shaping filter type", SELECTION: "RootRaisedCos"/"RaisedCos"/"Gaussian", DEFAULT:RootRaisedCos"`
	TxPulseShapeFltrRollOff float64 `DESCRIPTION: "TX pulse shape filter roll off factor, MIN:0.004, MAX:1.0", DEFAULT:0.301`
	AdminState              bool    `DESCRIPTION: "Administrative state of this network interface", DEFAULT: true`
}

type DWDMModuleNwIntfState struct {
	baseObj
	ModuleId                        uint8   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId                        uint8   `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module network interface identifier"`
	TxChanGridSpacing               string  `DESCRIPTION: "The channel grid spacing used for this network interface in GHz"`
	CurrentBER                      float64 `DESCRIPTION: "Current value of BER on the DWDM module network interface"`
	MinBEROverPMInterval            float64 `DESCRIPTION: "Minimum value of BER over the last PM interval for the DWDM module network interface"`
	AvgBEROverPMInterval            float64 `DESCRIPTION: "Average value of BER over the last PM interval for the DWDM module network interface"`
	MaxBEROverPMInterval            float64 `DESCRIPTION: "Maximum value of BER over the last PM interval for the DWDM module network interface"`
	CurrUncorrectableFECBlkCnt      float64 `DESCRIPTION: "Current value of uncorrectable FEC code block count"`
	UncorrectableFECBlkCntOverPMInt float64 `DESCRIPTION: "Average value of uncorrectable FEC code block count over the last PM interval"`
}

type DWDMModuleClntIntf struct {
	baseObj
	ModuleId                  uint8 `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER: "true", DESCRIPTION: "DWDM Module identifier"`
	ClntIntfId                uint8 `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module client interface identifier"`
	TXFECDecDisable           bool  `DESCRIPTION: "802.3bj FEC decoder enable/disable state for traffic from Host to DWDM Module", DEFAULT: false`
	RXFECDecDisable           bool  `DESCRIPTION: "802.3bj FEC decoder enable/disable state for traffic from DWDM module to Host", DEFAULT: false`
	HostTxEqLfCtle            uint8 `DESCRIPTION: "Host interface TX deserializer equalization. LELPZRC LF-CTLE LFPZ gain code.", MIN:0, MAX:8, DEFAULT:0`
	HostTxEqCtle              uint8 `DESCRIPTION: "Host interface TX deserializer equalization. LELRC CTLE LE gain code.", MIN:0, MAX:20, DEFAULT:18`
	HostTxEqDfe               uint8 `DESCRIPTION: "Host interface TX deserializer equalization. s-DFE, DFE tap coefficient", MIN:0, MAX:63, DEFAULT:0`
	HostRxSerializerTap0Gain  uint8 `DESCRIPTION: "Host RX Serializer tap 0 control, gain for equalization filter tap", DEFAULT:7, MIN:0, MAX:7`
	HostRxSerializerTap0Delay uint8 `DESCRIPTION: "Host RX Serializer tap 0 control, delay for equalization filter tap", DEFAULT:7, MIN:0, MAX:7`
	HostRxSerializerTap1Gain  uint8 `DESCRIPTION: "Host RX Serializer tap 1 control, gain for equalization filter tap", DEFAULT:7, MIN:0, MAX:7`
	HostRxSerializerTap2Gain  uint8 `DESCRIPTION: "Host RX Serializer tap 2 control, gain for equalization filter tap", DEFAULT:15, MIN:0, MAX:15`
	HostRxSerializerTap2Delay uint8 `DESCRIPTION: "Host RX Serializer tap 2 control, delay for equalization filter tap", DEFAULT:5, MIN:0, MAX:7`
	AdminState                bool  `DESCRIPTION: "Administrative state of this client interface", DEFAULT: true`
}

type DWDMModuleClntIntfState struct {
	baseObj
	ModuleId   uint8 `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ClntIntfId uint8 `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module client interface identifier"`
}

type DWDMModulePMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "PM Data Value"`
}

type DWDMModuleNwIntfPMState struct {
	baseObj
	ModuleId uint8  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId uint8  `SNAPROUTE: "KEY", DESCRIPTION: "DWDM Module network interface identifier"`
	Resource string `SNAPROUTE: "KEY", DESCRIPTION: "Opticd resource name for which PM Data is required"`
	Type     string `SNAPROUTE: "KEY", DESCRIPTION: "Min/Max/Avg"`
	Class    string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data     []DWDMModulePMData
}
