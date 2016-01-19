package models

import (
	"encoding/json"

	"fmt"
)

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
