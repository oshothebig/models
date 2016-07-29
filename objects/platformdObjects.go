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
	ObjName      string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "ObjName", DEFAULT: "System"`
	ProductName  string `DESCRIPTION: "Product Number"`
	SerialNum    string `DESCRIPTION: "Serial Number"`
	Manufacturer string `DESCRIPTION: "Manufacturer"`
	Vendor       string `DESCRIPTION: "Vendor"`
	Release      string `DESCRIPTION: "Relese version"`
	PlatformName string `DESCRIPTION: "Platform Number"`
	ONIEVersion  string `DESCRIPTION: "ONIE version"`
}

type Sfp struct {
	baseObj
	SfpId      int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "SFP id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Admin PORT UP/DOWN(TX OFF)"`
}

type SfpState struct {
	baseObj
	SfpId   int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "SFP id", DEFAULT:0`
	Status  string `DESCRIPTION: "SFP status PRESENT/MISSING"`
	SfpLOS  string `DESCRIPTION: "SFP status RX LOS"`
	SfpType string `DESCRIPTION: "SFP type Copper/Optical"`
	EEPROM  string `DESCRIPTION: "SFP eeprom"`
}

type ThermalState struct {
	baseObj
	ThermalId                 int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Thermal sensor id", DEFAULT:0`
	Location                  string `DESCRIPTION: "Thermal sensor location CPU/PSU/Motherboard"`
	Temperature               string `DESCRIPTION: "Temperature current"`
	LowerWatermarkTemperature string `DESCRIPTION: "Temperature warning"`
	UpperWatermarkTemperature string `DESCRIPTION: "Temperature error"`
	ShutdownTemperature       string `DESCRIPTION: "Temperature panic"`
}

type Psu struct {
	baseObj
	PsuId      int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "PSU id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Admin UP/DOWN PSU"`
}

type PsuState struct {
	baseObj
	PsuId      int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "PSU id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Admin UP/DOWN PSU"`
	ModelNum   string `DESCRIPTION: "Model Number"`
	SerialNum  string `DESCRIPTION: "Serial Number"`
	Vin        int32  `DESCRIPTION: "Voltage in"`
	Vout       int32  `DESCRIPTION: "Voltage out"`
	Iin        int32  `DESCRIPTION: "Current in"`
	Iout       int32  `DESCRIPTION: "Current out"`
	Pin        int32  `DESCRIPTION: "Power in"`
	Pout       int32  `DESCRIPTION: "power out"`
	Fan        string `DESCRIPTION: "Fan PRESENT/MISSING"`
	FanId      int32  `DESCRIPTION: "Fan Info"`
	LedId      int32  `DESCRIPTION: "LED Info"`
}

type Led struct {
	baseObj
	LedId       int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "LED id", DEFAULT:0`
	LedAdmin    string `DESCRIPTION: "LED ON/OFF"`
	LedSetColor string `DESCRIPTION: "LED set color"`
}

type LedState struct {
	baseObj
	LedId       int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "LED id", DEFAULT:0`
	LedIdentify string `DESCRIPTION: "LED represents FAN/PSU/RESET etc"`
	LedState    string `DESCRIPTION: "LED State ON/OFF"`
	LedColor    string `DESCRIPTION: "LED Color"`
}

type Fan struct {
	baseObj
	FanId      int32 `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Fan unit id", DEFAULT:0`
	AdminState int32 `DESCRIPTION: "Fan admin ON/OFF"`
	AdminSpeed int32 `DESCRIPTION: "Fan set speed in rpm"`
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
	LedId         int32  `DESCRIPTION: "LED Info"`
}
