package models

import (
	"encoding/json"

	"fmt"
)

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
