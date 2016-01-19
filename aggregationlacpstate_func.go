package models

import (
	"encoding/json"

	"fmt"
)

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
