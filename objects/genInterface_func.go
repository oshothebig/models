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

import (
	"encoding/json"

	"fmt"
)

func NewInterfaceConfig() *InterfaceConfig {
	newObj := &InterfaceConfig{
		Enabled: true,
	}
	return newObj
}

func (obj InterfaceConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### InterfaceConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *InterfaceConfig) Type_Set(value string) bool {
	d.Type = value
	return true
}
func (d *InterfaceConfig) Mtu_Set(value uint16) bool {
	d.Mtu = value
	return true
}
func (d *InterfaceConfig) Name_Set(value string) bool {
	d.Name = value
	return true
}
func (d *InterfaceConfig) Description_Set(value string) bool {
	d.Description = value
	return true
}
func (d *InterfaceConfig) Enabled_Set(value bool) bool {
	d.Enabled = value
	return true
}
func NewInterfaceStateCounters() *InterfaceStateCounters {
	newObj := &InterfaceStateCounters{}
	return newObj
}

func (obj InterfaceStateCounters) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### InterfaceStateCounters called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewHoldTimeConfig() *HoldTimeConfig {
	newObj := &HoldTimeConfig{}
	return newObj
}

func (obj HoldTimeConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### HoldTimeConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *HoldTimeConfig) Up_Set(value uint32) bool {
	d.Up = value
	return true
}
func (d *HoldTimeConfig) Down_Set(value uint32) bool {
	d.Down = value
	return true
}
func NewHoldTimeState() *HoldTimeState {
	newObj := &HoldTimeState{}
	return newObj
}

func (obj HoldTimeState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### HoldTimeState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceConfig() *SubinterfaceConfig {
	newObj := &SubinterfaceConfig{
		Unnumbered: false,
		Enabled:    true,
	}
	return newObj
}

func (obj SubinterfaceConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceConfig) Index_Set(value uint32) bool {
	d.Index = value
	return true
}
func (d *SubinterfaceConfig) Unnumbered_Set(value bool) bool {
	d.Unnumbered = value
	return true
}
func (d *SubinterfaceConfig) Name_Set(value string) bool {
	d.Name = value
	return true
}
func (d *SubinterfaceConfig) Description_Set(value string) bool {
	d.Description = value
	return true
}
func (d *SubinterfaceConfig) Enabled_Set(value bool) bool {
	d.Enabled = value
	return true
}
func NewSubinterfaceStateCounters() *SubinterfaceStateCounters {
	newObj := &SubinterfaceStateCounters{}
	return newObj
}

func (obj SubinterfaceStateCounters) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceStateCounters called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceVlanConfig() *SubinterfaceVlanConfig {
	newObj := &SubinterfaceVlanConfig{}
	return newObj
}

func (obj SubinterfaceVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceVlanConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceVlanConfig) GlobalVlanId_Set(value []uint16) bool {
	d.GlobalVlanId = value
	return true
}
func (d *SubinterfaceVlanConfig) VlanId_VlanId_Set(value uint16) bool {
	d.VlanId_VlanId = value
	return true
}
func (d *SubinterfaceVlanConfig) VlanId_VlanId_QinqId_Set(value string) bool {
	d.VlanId_VlanId_QinqId = value
	return true
}
func NewSubinterfaceVlanState() *SubinterfaceVlanState {
	newObj := &SubinterfaceVlanState{}
	return newObj
}

func (obj SubinterfaceVlanState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceVlanState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv4AddressConfig() *SubinterfaceIpv4AddressConfig {
	newObj := &SubinterfaceIpv4AddressConfig{}
	return newObj
}

func (obj SubinterfaceIpv4AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4AddressConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv4AddressConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *SubinterfaceIpv4AddressConfig) PrefixLength_Set(value uint8) bool {
	if value < 0 || value > 32 {
		return false
	}

	d.PrefixLength = value
	return true
}
func NewSubinterfaceIpv4AddressState() *SubinterfaceIpv4AddressState {
	newObj := &SubinterfaceIpv4AddressState{}
	return newObj
}

func (obj SubinterfaceIpv4AddressState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4AddressState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv4AddressVrrpVrrpGroupConfig() *SubinterfaceIpv4AddressVrrpVrrpGroupConfig {
	newObj := &SubinterfaceIpv4AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4AddressVrrpVrrpGroupConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) VirtualRouterId_Set(value uint8) bool {
	if value < 1 || value > 255 {
		return false
	}

	d.VirtualRouterId = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) VirtualAddress_Set(value []string) bool {
	d.VirtualAddress = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) Priority_Set(value uint8) bool {
	if value < 1 || value > 254 {
		return false
	}

	d.Priority = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) Preempt_Set(value bool) bool {
	d.Preempt = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) PreemptDelay_Set(value uint16) bool {
	if value < 0 || value > 3600 {
		return false
	}

	d.PreemptDelay = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) AcceptMode_Set(value bool) bool {
	d.AcceptMode = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupConfig) AdvertisementInterval_Set(value uint16) bool {
	if value < 1 || value > 4095 {
		return false
	}

	d.AdvertisementInterval = value
	return true
}
func NewSubinterfaceIpv4AddressVrrpVrrpGroupState() *SubinterfaceIpv4AddressVrrpVrrpGroupState {
	newObj := &SubinterfaceIpv4AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4AddressVrrpVrrpGroupState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig() *SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig {
	newObj := &SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return newObj
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) TrackInterface_Set(value string) bool {
	d.TrackInterface = value
	return true
}
func (d *SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) PriorityDecrement_Set(value uint8) bool {
	if value < 0 || value > 254 {
		return false
	}

	d.PriorityDecrement = value
	return true
}
func NewSubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState() *SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState {
	newObj := &SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState{}
	return newObj
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv4NeighborConfig() *SubinterfaceIpv4NeighborConfig {
	newObj := &SubinterfaceIpv4NeighborConfig{}
	return newObj
}

func (obj SubinterfaceIpv4NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4NeighborConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv4NeighborConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *SubinterfaceIpv4NeighborConfig) LinkLayerAddress_Set(value string) bool {
	d.LinkLayerAddress = value
	return true
}
func NewSubinterfaceIpv4NeighborState() *SubinterfaceIpv4NeighborState {
	newObj := &SubinterfaceIpv4NeighborState{}
	return newObj
}

func (obj SubinterfaceIpv4NeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4NeighborState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv4Config() *SubinterfaceIpv4Config {
	newObj := &SubinterfaceIpv4Config{
		Enabled: true,
	}
	return newObj
}

func (obj SubinterfaceIpv4Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4Config called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv4Config) Enabled_Set(value bool) bool {
	d.Enabled = value
	return true
}
func (d *SubinterfaceIpv4Config) Mtu_Set(value uint16) bool {
	if value < 68 || value > ^uint16(0) {
		return false
	}

	d.Mtu = value
	return true
}
func NewSubinterfaceIpv4State() *SubinterfaceIpv4State {
	newObj := &SubinterfaceIpv4State{
		Enabled: true,
	}
	return newObj
}

func (obj SubinterfaceIpv4State) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv4State called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv6AddressConfig() *SubinterfaceIpv6AddressConfig {
	newObj := &SubinterfaceIpv6AddressConfig{}
	return newObj
}

func (obj SubinterfaceIpv6AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AddressConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv6AddressConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *SubinterfaceIpv6AddressConfig) PrefixLength_Set(value uint8) bool {
	if value < 0 || value > 128 {
		return false
	}

	d.PrefixLength = value
	return true
}
func NewSubinterfaceIpv6AddressState() *SubinterfaceIpv6AddressState {
	newObj := &SubinterfaceIpv6AddressState{}
	return newObj
}

func (obj SubinterfaceIpv6AddressState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AddressState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv6AddressVrrpVrrpGroupConfig() *SubinterfaceIpv6AddressVrrpVrrpGroupConfig {
	newObj := &SubinterfaceIpv6AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AddressVrrpVrrpGroupConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) VirtualRouterId_Set(value uint8) bool {
	if value < 1 || value > 255 {
		return false
	}

	d.VirtualRouterId = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) VirtualAddress_Set(value []string) bool {
	d.VirtualAddress = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) Priority_Set(value uint8) bool {
	if value < 1 || value > 254 {
		return false
	}

	d.Priority = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) Preempt_Set(value bool) bool {
	d.Preempt = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) PreemptDelay_Set(value uint16) bool {
	if value < 0 || value > 3600 {
		return false
	}

	d.PreemptDelay = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) AcceptMode_Set(value bool) bool {
	d.AcceptMode = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) AdvertisementInterval_Set(value uint16) bool {
	if value < 1 || value > 4095 {
		return false
	}

	d.AdvertisementInterval = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupConfig) VirtualLinkLocal_Set(value []string) bool {
	d.VirtualLinkLocal = value
	return true
}
func NewSubinterfaceIpv6AddressVrrpVrrpGroupState() *SubinterfaceIpv6AddressVrrpVrrpGroupState {
	newObj := &SubinterfaceIpv6AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AddressVrrpVrrpGroupState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig() *SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig {
	newObj := &SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return newObj
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) TrackInterface_Set(value string) bool {
	d.TrackInterface = value
	return true
}
func (d *SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) PriorityDecrement_Set(value uint8) bool {
	if value < 0 || value > 254 {
		return false
	}

	d.PriorityDecrement = value
	return true
}
func NewSubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState() *SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState {
	newObj := &SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState{}
	return newObj
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv6NeighborConfig() *SubinterfaceIpv6NeighborConfig {
	newObj := &SubinterfaceIpv6NeighborConfig{}
	return newObj
}

func (obj SubinterfaceIpv6NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6NeighborConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv6NeighborConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *SubinterfaceIpv6NeighborConfig) LinkLayerAddress_Set(value string) bool {
	d.LinkLayerAddress = value
	return true
}
func NewSubinterfaceIpv6NeighborState() *SubinterfaceIpv6NeighborState {
	newObj := &SubinterfaceIpv6NeighborState{}
	return newObj
}

func (obj SubinterfaceIpv6NeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6NeighborState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv6Config() *SubinterfaceIpv6Config {
	newObj := &SubinterfaceIpv6Config{
		Enabled: true,
	}
	return newObj
}

func (obj SubinterfaceIpv6Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6Config called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv6Config) Enabled_Set(value bool) bool {
	d.Enabled = value
	return true
}
func (d *SubinterfaceIpv6Config) Mtu_Set(value uint32) bool {
	if value < 1280 || value > ^uint32(0) {
		return false
	}

	d.Mtu = value
	return true
}
func (d *SubinterfaceIpv6Config) DupAddrDetectTransmits_Set(value uint32) bool {
	d.DupAddrDetectTransmits = value
	return true
}
func NewSubinterfaceIpv6State() *SubinterfaceIpv6State {
	newObj := &SubinterfaceIpv6State{
		Enabled: true,
	}
	return newObj
}

func (obj SubinterfaceIpv6State) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6State called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewSubinterfaceIpv6AutoconfConfig() *SubinterfaceIpv6AutoconfConfig {
	newObj := &SubinterfaceIpv6AutoconfConfig{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return newObj
}

func (obj SubinterfaceIpv6AutoconfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AutoconfConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *SubinterfaceIpv6AutoconfConfig) CreateGlobalAddresses_Set(value bool) bool {
	d.CreateGlobalAddresses = value
	return true
}
func (d *SubinterfaceIpv6AutoconfConfig) CreateTemporaryAddresses_Set(value bool) bool {
	d.CreateTemporaryAddresses = value
	return true
}
func (d *SubinterfaceIpv6AutoconfConfig) TemporaryValidLifetime_Set(value uint32) bool {
	d.TemporaryValidLifetime = value
	return true
}
func (d *SubinterfaceIpv6AutoconfConfig) TemporaryPreferredLifetime_Set(value uint32) bool {
	d.TemporaryPreferredLifetime = value
	return true
}
func NewSubinterfaceIpv6AutoconfState() *SubinterfaceIpv6AutoconfState {
	newObj := &SubinterfaceIpv6AutoconfState{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return newObj
}

func (obj SubinterfaceIpv6AutoconfState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### SubinterfaceIpv6AutoconfState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewEthernetConfig() *EthernetConfig {
	newObj := &EthernetConfig{
		EnableFlowControl: false,
	}
	return newObj
}

func (obj EthernetConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### EthernetConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *EthernetConfig) MacAddress_Set(value string) bool {
	d.MacAddress = value
	return true
}
func (d *EthernetConfig) DuplexMode_Set(value int32) bool {
	d.DuplexMode = value
	return true
}
func (d *EthernetConfig) Auto_Set(value bool) bool {
	d.Auto = value
	return true
}
func (d *EthernetConfig) Speed_Set(value string) bool {
	d.Speed = value
	return true
}
func (d *EthernetConfig) EnableFlowControl_Set(value bool) bool {
	d.EnableFlowControl = value
	return true
}
func (d *EthernetConfig) AggregateId_Set(value string) bool {
	d.AggregateId = value
	return true
}
func NewEthernetStateCounters() *EthernetStateCounters {
	newObj := &EthernetStateCounters{}
	return newObj
}

func (obj EthernetStateCounters) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### EthernetStateCounters called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewEthernetState() *EthernetState {
	newObj := &EthernetState{}
	return newObj
}

func (obj EthernetState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### EthernetState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewEthernetVlanConfig() *EthernetVlanConfig {
	newObj := &EthernetVlanConfig{}
	return newObj
}

func (obj EthernetVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### EthernetVlanConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *EthernetVlanConfig) InterfaceMode_Set(value int32) bool {
	d.InterfaceMode = value
	return true
}
func (d *EthernetVlanConfig) NativeVlan_VlanId_Set(value uint16) bool {
	d.NativeVlan_VlanId = value
	return true
}
func (d *EthernetVlanConfig) NativeVlan_VlanId_QinqId_Set(value string) bool {
	d.NativeVlan_VlanId_QinqId = value
	return true
}
func (d *EthernetVlanConfig) AccessVlan_VlanId_Set(value uint16) bool {
	d.AccessVlan_VlanId = value
	return true
}
func (d *EthernetVlanConfig) AccessVlan_VlanId_QinqId_Set(value string) bool {
	d.AccessVlan_VlanId_QinqId = value
	return true
}
func (d *EthernetVlanConfig) TrunkVlans_VlanId_Set(value uint16) bool {
	d.TrunkVlans_VlanId = value
	return true
}
func (d *EthernetVlanConfig) TrunkVlans_VlanId_VlanRange_Set(value string) bool {
	d.TrunkVlans_VlanId_VlanRange = value
	return true
}
func (d *EthernetVlanConfig) TrunkVlans_VlanId_VlanRange_QinqId_Set(value string) bool {
	d.TrunkVlans_VlanId_VlanRange_QinqId = value
	return true
}
func (d *EthernetVlanConfig) TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange_Set(value string) bool {
	d.TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange = value
	return true
}
func NewEthernetVlanState() *EthernetVlanState {
	newObj := &EthernetVlanState{}
	return newObj
}

func (obj EthernetVlanState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### EthernetVlanState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewAggregationConfig() *AggregationConfig {
	newObj := &AggregationConfig{}
	return newObj
}

func (obj AggregationConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *AggregationConfig) LagType_Set(value int32) bool {
	d.LagType = value
	return true
}
func (d *AggregationConfig) MinLinks_Set(value uint16) bool {
	d.MinLinks = value
	return true
}
func NewAggregationState() *AggregationState {
	newObj := &AggregationState{}
	return newObj
}

func (obj AggregationState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewAggregationLacpConfig() *AggregationLacpConfig {
	newObj := &AggregationLacpConfig{}
	return newObj
}

func (obj AggregationLacpConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationLacpConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *AggregationLacpConfig) Interval_Set(value int32) bool {
	d.Interval = value
	return true
}
func (d *AggregationLacpConfig) LacpMode_Set(value int32) bool {
	d.LacpMode = value
	return true
}
func (d *AggregationLacpConfig) SystemIdMac_Set(value string) bool {
	d.SystemIdMac = value
	return true
}
func (d *AggregationLacpConfig) SystemPriority_Set(value uint16) bool {
	d.SystemPriority = value
	return true
}
func (d *AggregationLacpConfig) LagHash_Set(value int32) bool {
	d.LagHash = value
	return true
}
func NewAggregationLacpState() *AggregationLacpState {
	newObj := &AggregationLacpState{}
	return newObj
}

func (obj AggregationLacpState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationLacpState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewAggregationLacpMemberStateCounters() *AggregationLacpMemberStateCounters {
	newObj := &AggregationLacpMemberStateCounters{}
	return newObj
}

func (obj AggregationLacpMemberStateCounters) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationLacpMemberStateCounters called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewAggregationVlanConfig() *AggregationVlanConfig {
	newObj := &AggregationVlanConfig{}
	return newObj
}

func (obj AggregationVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationVlanConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *AggregationVlanConfig) InterfaceMode_Set(value int32) bool {
	d.InterfaceMode = value
	return true
}
func (d *AggregationVlanConfig) NativeVlan_VlanId_Set(value uint16) bool {
	d.NativeVlan_VlanId = value
	return true
}
func (d *AggregationVlanConfig) NativeVlan_VlanId_QinqId_Set(value string) bool {
	d.NativeVlan_VlanId_QinqId = value
	return true
}
func (d *AggregationVlanConfig) AccessVlan_VlanId_Set(value uint16) bool {
	d.AccessVlan_VlanId = value
	return true
}
func (d *AggregationVlanConfig) AccessVlan_VlanId_QinqId_Set(value string) bool {
	d.AccessVlan_VlanId_QinqId = value
	return true
}
func (d *AggregationVlanConfig) TrunkVlans_VlanId_Set(value uint16) bool {
	d.TrunkVlans_VlanId = value
	return true
}
func (d *AggregationVlanConfig) TrunkVlans_VlanId_VlanRange_Set(value string) bool {
	d.TrunkVlans_VlanId_VlanRange = value
	return true
}
func (d *AggregationVlanConfig) TrunkVlans_VlanId_VlanRange_QinqId_Set(value string) bool {
	d.TrunkVlans_VlanId_VlanRange_QinqId = value
	return true
}
func (d *AggregationVlanConfig) TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange_Set(value string) bool {
	d.TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange = value
	return true
}
func NewAggregationVlanState() *AggregationVlanState {
	newObj := &AggregationVlanState{}
	return newObj
}

func (obj AggregationVlanState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### AggregationVlanState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanConfig() *RoutedVlanConfig {
	newObj := &RoutedVlanConfig{}
	return newObj
}

func (obj RoutedVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanConfig) Vlan_Uint16_Set(value uint16) bool {
	d.Vlan_Uint16 = value
	return true
}
func (d *RoutedVlanConfig) Vlan_Uint16_String_Set(value string) bool {
	d.Vlan_Uint16_String = value
	return true
}
func NewRoutedVlanState() *RoutedVlanState {
	newObj := &RoutedVlanState{}
	return newObj
}

func (obj RoutedVlanState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv4AddressConfig() *RoutedVlanIpv4AddressConfig {
	newObj := &RoutedVlanIpv4AddressConfig{}
	return newObj
}

func (obj RoutedVlanIpv4AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4AddressConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv4AddressConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *RoutedVlanIpv4AddressConfig) PrefixLength_Set(value uint8) bool {
	if value < 0 || value > 32 {
		return false
	}

	d.PrefixLength = value
	return true
}
func NewRoutedVlanIpv4AddressState() *RoutedVlanIpv4AddressState {
	newObj := &RoutedVlanIpv4AddressState{}
	return newObj
}

func (obj RoutedVlanIpv4AddressState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4AddressState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv4AddressVrrpVrrpGroupConfig() *RoutedVlanIpv4AddressVrrpVrrpGroupConfig {
	newObj := &RoutedVlanIpv4AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4AddressVrrpVrrpGroupConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) VirtualRouterId_Set(value uint8) bool {
	if value < 1 || value > 255 {
		return false
	}

	d.VirtualRouterId = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) VirtualAddress_Set(value []string) bool {
	d.VirtualAddress = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) Priority_Set(value uint8) bool {
	if value < 1 || value > 254 {
		return false
	}

	d.Priority = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) Preempt_Set(value bool) bool {
	d.Preempt = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) PreemptDelay_Set(value uint16) bool {
	if value < 0 || value > 3600 {
		return false
	}

	d.PreemptDelay = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) AcceptMode_Set(value bool) bool {
	d.AcceptMode = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupConfig) AdvertisementInterval_Set(value uint16) bool {
	if value < 1 || value > 4095 {
		return false
	}

	d.AdvertisementInterval = value
	return true
}
func NewRoutedVlanIpv4AddressVrrpVrrpGroupState() *RoutedVlanIpv4AddressVrrpVrrpGroupState {
	newObj := &RoutedVlanIpv4AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4AddressVrrpVrrpGroupState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig() *RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig {
	newObj := &RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return newObj
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) TrackInterface_Set(value string) bool {
	d.TrackInterface = value
	return true
}
func (d *RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) PriorityDecrement_Set(value uint8) bool {
	if value < 0 || value > 254 {
		return false
	}

	d.PriorityDecrement = value
	return true
}
func NewRoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState() *RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState {
	newObj := &RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState{}
	return newObj
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv4NeighborConfig() *RoutedVlanIpv4NeighborConfig {
	newObj := &RoutedVlanIpv4NeighborConfig{}
	return newObj
}

func (obj RoutedVlanIpv4NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4NeighborConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv4NeighborConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *RoutedVlanIpv4NeighborConfig) LinkLayerAddress_Set(value string) bool {
	d.LinkLayerAddress = value
	return true
}
func NewRoutedVlanIpv4NeighborState() *RoutedVlanIpv4NeighborState {
	newObj := &RoutedVlanIpv4NeighborState{}
	return newObj
}

func (obj RoutedVlanIpv4NeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4NeighborState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv4Config() *RoutedVlanIpv4Config {
	newObj := &RoutedVlanIpv4Config{
		Enabled: true,
	}
	return newObj
}

func (obj RoutedVlanIpv4Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4Config called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv4Config) Enabled_Set(value bool) bool {
	d.Enabled = value
	return true
}
func (d *RoutedVlanIpv4Config) Mtu_Set(value uint16) bool {
	if value < 68 || value > ^uint16(0) {
		return false
	}

	d.Mtu = value
	return true
}
func NewRoutedVlanIpv4State() *RoutedVlanIpv4State {
	newObj := &RoutedVlanIpv4State{
		Enabled: true,
	}
	return newObj
}

func (obj RoutedVlanIpv4State) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv4State called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv6AddressConfig() *RoutedVlanIpv6AddressConfig {
	newObj := &RoutedVlanIpv6AddressConfig{}
	return newObj
}

func (obj RoutedVlanIpv6AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AddressConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv6AddressConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *RoutedVlanIpv6AddressConfig) PrefixLength_Set(value uint8) bool {
	if value < 0 || value > 128 {
		return false
	}

	d.PrefixLength = value
	return true
}
func NewRoutedVlanIpv6AddressState() *RoutedVlanIpv6AddressState {
	newObj := &RoutedVlanIpv6AddressState{}
	return newObj
}

func (obj RoutedVlanIpv6AddressState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AddressState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv6AddressVrrpVrrpGroupConfig() *RoutedVlanIpv6AddressVrrpVrrpGroupConfig {
	newObj := &RoutedVlanIpv6AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AddressVrrpVrrpGroupConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) VirtualRouterId_Set(value uint8) bool {
	if value < 1 || value > 255 {
		return false
	}

	d.VirtualRouterId = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) VirtualAddress_Set(value []string) bool {
	d.VirtualAddress = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) Priority_Set(value uint8) bool {
	if value < 1 || value > 254 {
		return false
	}

	d.Priority = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) Preempt_Set(value bool) bool {
	d.Preempt = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) PreemptDelay_Set(value uint16) bool {
	if value < 0 || value > 3600 {
		return false
	}

	d.PreemptDelay = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) AcceptMode_Set(value bool) bool {
	d.AcceptMode = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) AdvertisementInterval_Set(value uint16) bool {
	if value < 1 || value > 4095 {
		return false
	}

	d.AdvertisementInterval = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupConfig) VirtualLinkLocal_Set(value []string) bool {
	d.VirtualLinkLocal = value
	return true
}
func NewRoutedVlanIpv6AddressVrrpVrrpGroupState() *RoutedVlanIpv6AddressVrrpVrrpGroupState {
	newObj := &RoutedVlanIpv6AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return newObj
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AddressVrrpVrrpGroupState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig() *RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig {
	newObj := &RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return newObj
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) TrackInterface_Set(value string) bool {
	d.TrackInterface = value
	return true
}
func (d *RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) PriorityDecrement_Set(value uint8) bool {
	if value < 0 || value > 254 {
		return false
	}

	d.PriorityDecrement = value
	return true
}
func NewRoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState() *RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState {
	newObj := &RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState{}
	return newObj
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv6NeighborConfig() *RoutedVlanIpv6NeighborConfig {
	newObj := &RoutedVlanIpv6NeighborConfig{}
	return newObj
}

func (obj RoutedVlanIpv6NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6NeighborConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv6NeighborConfig) Ip_Set(value string) bool {
	d.Ip = value
	return true
}
func (d *RoutedVlanIpv6NeighborConfig) LinkLayerAddress_Set(value string) bool {
	d.LinkLayerAddress = value
	return true
}
func NewRoutedVlanIpv6NeighborState() *RoutedVlanIpv6NeighborState {
	newObj := &RoutedVlanIpv6NeighborState{}
	return newObj
}

func (obj RoutedVlanIpv6NeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6NeighborState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv6Config() *RoutedVlanIpv6Config {
	newObj := &RoutedVlanIpv6Config{
		Enabled: true,
	}
	return newObj
}

func (obj RoutedVlanIpv6Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6Config called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv6Config) Enabled_Set(value bool) bool {
	d.Enabled = value
	return true
}
func (d *RoutedVlanIpv6Config) Mtu_Set(value uint32) bool {
	if value < 1280 || value > ^uint32(0) {
		return false
	}

	d.Mtu = value
	return true
}
func (d *RoutedVlanIpv6Config) DupAddrDetectTransmits_Set(value uint32) bool {
	d.DupAddrDetectTransmits = value
	return true
}
func NewRoutedVlanIpv6State() *RoutedVlanIpv6State {
	newObj := &RoutedVlanIpv6State{
		Enabled: true,
	}
	return newObj
}

func (obj RoutedVlanIpv6State) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6State called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewRoutedVlanIpv6AutoconfConfig() *RoutedVlanIpv6AutoconfConfig {
	newObj := &RoutedVlanIpv6AutoconfConfig{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return newObj
}

func (obj RoutedVlanIpv6AutoconfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AutoconfConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *RoutedVlanIpv6AutoconfConfig) CreateGlobalAddresses_Set(value bool) bool {
	d.CreateGlobalAddresses = value
	return true
}
func (d *RoutedVlanIpv6AutoconfConfig) CreateTemporaryAddresses_Set(value bool) bool {
	d.CreateTemporaryAddresses = value
	return true
}
func (d *RoutedVlanIpv6AutoconfConfig) TemporaryValidLifetime_Set(value uint32) bool {
	d.TemporaryValidLifetime = value
	return true
}
func (d *RoutedVlanIpv6AutoconfConfig) TemporaryPreferredLifetime_Set(value uint32) bool {
	d.TemporaryPreferredLifetime = value
	return true
}
func NewRoutedVlanIpv6AutoconfState() *RoutedVlanIpv6AutoconfState {
	newObj := &RoutedVlanIpv6AutoconfState{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return newObj
}

func (obj RoutedVlanIpv6AutoconfState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### RoutedVlanIpv6AutoconfState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
