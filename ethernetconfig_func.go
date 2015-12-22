package models

import (
	"encoding/json"

	"fmt"
)

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
