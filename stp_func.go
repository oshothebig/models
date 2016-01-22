package models

import (
	"encoding/json"

	"fmt"
)

func NewDot1dStpPortEntryConfig() *Dot1dStpPortEntryConfig {
	newObj := &Dot1dStpPortEntryConfig{}
	return newObj
}

func (obj Dot1dStpPortEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpPortEntryConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortPriority_Set(value int32) bool {
	if value < 0 || value > 255 {
		return false
	}

	d.Dot1dStpPortPriority = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortEnable_Set(value int32) bool {
	d.Dot1dStpPortEnable = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortPathCost_Set(value int32) bool {
	if value < 1 || value > 65535 {
		return false
	}

	d.Dot1dStpPortPathCost = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortPathCost32_Set(value int32) bool {
	if value < 1 || value > 200000000 {
		return false
	}

	d.Dot1dStpPortPathCost32 = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortProtocolMigration_Set(value int32) bool {
	d.Dot1dStpPortProtocolMigration = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortAdminPointToPoint_Set(value int32) bool {
	d.Dot1dStpPortAdminPointToPoint = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortAdminEdgePort_Set(value int32) bool {
	d.Dot1dStpPortAdminEdgePort = value
	return true
}
func (d *Dot1dStpPortEntryConfig) Dot1dStpPortAdminPathCost_Set(value int32) bool {
	if value < 0 || value > 200000000 {
		return false
	}

	d.Dot1dStpPortAdminPathCost = value
	return true
}
func NewDot1dStpPortEntryState() *Dot1dStpPortEntryState {
	newObj := &Dot1dStpPortEntryState{}
	return newObj
}

func (obj Dot1dStpPortEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpPortEntryState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func NewDot1dStpBridgeConfig() *Dot1dStpBridgeConfig {
	newObj := &Dot1dStpBridgeConfig{}
	return newObj
}

func (obj Dot1dStpBridgeConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpBridgeConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (d *Dot1dStpBridgeConfig) Dot1dStpBridgeMaxAge_Set(value int32) bool {
	if value < 600 || value > 4000 {
		return false
	}

	d.Dot1dStpBridgeMaxAge = value
	return true
}
func (d *Dot1dStpBridgeConfig) Dot1dStpBridgeHelloTime_Set(value int32) bool {
	if value < 100 || value > 1000 {
		return false
	}

	d.Dot1dStpBridgeHelloTime = value
	return true
}
func (d *Dot1dStpBridgeConfig) Dot1dStpBridgeForwardDelay_Set(value int32) bool {
	if value < 400 || value > 3000 {
		return false
	}

	d.Dot1dStpBridgeForwardDelay = value
	return true
}
func (d *Dot1dStpBridgeConfig) Dot1dStpBridgeForceVersion_Set(value int32) bool {
	d.Dot1dStpBridgeForceVersion = value
	return true
}
func (d *Dot1dStpBridgeConfig) Dot1dStpBridgeTxHoldCount_Set(value int32) bool {
	d.Dot1dStpBridgeTxHoldCount = value
	return true
}
func NewDot1dStpBridgeState() *Dot1dStpBridgeState {
	newObj := &Dot1dStpBridgeState{}
	return newObj
}

func (obj Dot1dStpBridgeState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpBridgeState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
