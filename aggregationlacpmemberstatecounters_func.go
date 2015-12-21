package models

import (
	"encoding/json"

	"fmt"
)

func NewAggregationLacpMemberStateCounters() *AggregationLacpMemberStateCounters {
	new := &AggregationLacpMemberStateCounters{}
	return new
}

func (obj AggregationLacpMemberStateCounters) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj AggregationLacpMemberStateCounters
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### AggregationLacpMemberStateCounters create called, unmarshal failed", Obj, err)
	}
	return Obj, err
}
