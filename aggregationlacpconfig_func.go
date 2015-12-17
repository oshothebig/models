package models

import (
	"encoding/json"

	"fmt"
)

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
