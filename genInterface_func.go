package models

import (
	"encoding/json"

	"fmt"
)

func NewInterfaceConfig() *InterfaceConfig {
	new := &InterfaceConfig{
		Enabled: true,
	}
	return new
}

func (obj InterfaceConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj InterfaceConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### InterfaceConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &InterfaceStateCounters{}
	return new
}

func NewHoldTimeConfig() *HoldTimeConfig {
	new := &HoldTimeConfig{}
	return new
}

func (obj HoldTimeConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj HoldTimeConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### HoldTimeConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &HoldTimeState{}
	return new
}

func NewSubinterfaceConfig() *SubinterfaceConfig {
	new := &SubinterfaceConfig{
		Unnumbered: false,
		Enabled:    true,
	}
	return new
}

func (obj SubinterfaceConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceStateCounters{}
	return new
}

func NewSubinterfaceVlanConfig() *SubinterfaceVlanConfig {
	new := &SubinterfaceVlanConfig{}
	return new
}

func (obj SubinterfaceVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceVlanConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceVlanConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceVlanState{}
	return new
}

func NewSubinterfaceIpv4AddressConfig() *SubinterfaceIpv4AddressConfig {
	new := &SubinterfaceIpv4AddressConfig{}
	return new
}

func (obj SubinterfaceIpv4AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv4AddressConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv4AddressConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv4AddressState{}
	return new
}

func NewSubinterfaceIpv4AddressVrrpVrrpGroupConfig() *SubinterfaceIpv4AddressVrrpVrrpGroupConfig {
	new := &SubinterfaceIpv4AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv4AddressVrrpVrrpGroupConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv4AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func NewSubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig() *SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig {
	new := &SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return new
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState{}
	return new
}

func NewSubinterfaceIpv4NeighborConfig() *SubinterfaceIpv4NeighborConfig {
	new := &SubinterfaceIpv4NeighborConfig{}
	return new
}

func (obj SubinterfaceIpv4NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv4NeighborConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv4NeighborConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv4NeighborState{}
	return new
}

func NewSubinterfaceIpv4Config() *SubinterfaceIpv4Config {
	new := &SubinterfaceIpv4Config{
		Enabled: true,
	}
	return new
}

func (obj SubinterfaceIpv4Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv4Config
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv4Config create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv4State{
		Enabled: true,
	}
	return new
}

func NewSubinterfaceIpv6AddressConfig() *SubinterfaceIpv6AddressConfig {
	new := &SubinterfaceIpv6AddressConfig{}
	return new
}

func (obj SubinterfaceIpv6AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv6AddressConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv6AddressConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv6AddressState{}
	return new
}

func NewSubinterfaceIpv6AddressVrrpVrrpGroupConfig() *SubinterfaceIpv6AddressVrrpVrrpGroupConfig {
	new := &SubinterfaceIpv6AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv6AddressVrrpVrrpGroupConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv6AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func NewSubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig() *SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig {
	new := &SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return new
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState{}
	return new
}

func NewSubinterfaceIpv6NeighborConfig() *SubinterfaceIpv6NeighborConfig {
	new := &SubinterfaceIpv6NeighborConfig{}
	return new
}

func (obj SubinterfaceIpv6NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv6NeighborConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv6NeighborConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv6NeighborState{}
	return new
}

func NewSubinterfaceIpv6Config() *SubinterfaceIpv6Config {
	new := &SubinterfaceIpv6Config{
		Enabled: true,
	}
	return new
}

func (obj SubinterfaceIpv6Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv6Config
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv6Config create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv6State{
		Enabled: true,
	}
	return new
}

func NewSubinterfaceIpv6AutoconfConfig() *SubinterfaceIpv6AutoconfConfig {
	new := &SubinterfaceIpv6AutoconfConfig{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return new
}

func (obj SubinterfaceIpv6AutoconfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj SubinterfaceIpv6AutoconfConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### SubinterfaceIpv6AutoconfConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &SubinterfaceIpv6AutoconfState{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return new
}

func NewEthernetConfig() *EthernetConfig {
	new := &EthernetConfig{
		EnableFlowControl: false,
	}
	return new
}

func (obj EthernetConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj EthernetConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### EthernetConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &EthernetStateCounters{}
	return new
}

func NewEthernetState() *EthernetState {
	new := &EthernetState{}
	return new
}

func NewEthernetVlanConfig() *EthernetVlanConfig {
	new := &EthernetVlanConfig{}
	return new
}

func (obj EthernetVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj EthernetVlanConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### EthernetVlanConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &EthernetVlanState{}
	return new
}

func NewAggregationConfig() *AggregationConfig {
	new := &AggregationConfig{}
	return new
}

func (obj AggregationConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj AggregationConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### AggregationConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &AggregationState{}
	return new
}

func NewAggregationLacpConfig() *AggregationLacpConfig {
	new := &AggregationLacpConfig{}
	return new
}

func (obj AggregationLacpConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj AggregationLacpConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### AggregationLacpConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
func NewAggregationLacpState() *AggregationLacpState {
	new := &AggregationLacpState{}
	return new
}

func NewAggregationLacpMemberStateCounters() *AggregationLacpMemberStateCounters {
	new := &AggregationLacpMemberStateCounters{}
	return new
}

func NewAggregationVlanConfig() *AggregationVlanConfig {
	new := &AggregationVlanConfig{}
	return new
}

func (obj AggregationVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj AggregationVlanConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### AggregationVlanConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &AggregationVlanState{}
	return new
}

func NewRoutedVlanConfig() *RoutedVlanConfig {
	new := &RoutedVlanConfig{}
	return new
}

func (obj RoutedVlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanState{}
	return new
}

func NewRoutedVlanIpv4AddressConfig() *RoutedVlanIpv4AddressConfig {
	new := &RoutedVlanIpv4AddressConfig{}
	return new
}

func (obj RoutedVlanIpv4AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv4AddressConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv4AddressConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv4AddressState{}
	return new
}

func NewRoutedVlanIpv4AddressVrrpVrrpGroupConfig() *RoutedVlanIpv4AddressVrrpVrrpGroupConfig {
	new := &RoutedVlanIpv4AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv4AddressVrrpVrrpGroupConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv4AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func NewRoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig() *RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig {
	new := &RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return new
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState{}
	return new
}

func NewRoutedVlanIpv4NeighborConfig() *RoutedVlanIpv4NeighborConfig {
	new := &RoutedVlanIpv4NeighborConfig{}
	return new
}

func (obj RoutedVlanIpv4NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv4NeighborConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv4NeighborConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv4NeighborState{}
	return new
}

func NewRoutedVlanIpv4Config() *RoutedVlanIpv4Config {
	new := &RoutedVlanIpv4Config{
		Enabled: true,
	}
	return new
}

func (obj RoutedVlanIpv4Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv4Config
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv4Config create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv4State{
		Enabled: true,
	}
	return new
}

func NewRoutedVlanIpv6AddressConfig() *RoutedVlanIpv6AddressConfig {
	new := &RoutedVlanIpv6AddressConfig{}
	return new
}

func (obj RoutedVlanIpv6AddressConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv6AddressConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv6AddressConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv6AddressState{}
	return new
}

func NewRoutedVlanIpv6AddressVrrpVrrpGroupConfig() *RoutedVlanIpv6AddressVrrpVrrpGroupConfig {
	new := &RoutedVlanIpv6AddressVrrpVrrpGroupConfig{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv6AddressVrrpVrrpGroupConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv6AddressVrrpVrrpGroupState{
		Preempt:    true,
		AcceptMode: false,
	}
	return new
}

func NewRoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig() *RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig {
	new := &RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig{}
	return new
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState{}
	return new
}

func NewRoutedVlanIpv6NeighborConfig() *RoutedVlanIpv6NeighborConfig {
	new := &RoutedVlanIpv6NeighborConfig{}
	return new
}

func (obj RoutedVlanIpv6NeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv6NeighborConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv6NeighborConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv6NeighborState{}
	return new
}

func NewRoutedVlanIpv6Config() *RoutedVlanIpv6Config {
	new := &RoutedVlanIpv6Config{
		Enabled: true,
	}
	return new
}

func (obj RoutedVlanIpv6Config) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv6Config
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv6Config create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv6State{
		Enabled: true,
	}
	return new
}

func NewRoutedVlanIpv6AutoconfConfig() *RoutedVlanIpv6AutoconfConfig {
	new := &RoutedVlanIpv6AutoconfConfig{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return new
}

func (obj RoutedVlanIpv6AutoconfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj RoutedVlanIpv6AutoconfConfig
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### RoutedVlanIpv6AutoconfConfig create called, unmarshal failed", Obj)
	}
	return Obj, err
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
	new := &RoutedVlanIpv6AutoconfState{
		CreateGlobalAddresses:    true,
		CreateTemporaryAddresses: false,
	}
	return new
}
