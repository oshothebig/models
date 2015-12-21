package models

import (
	"encoding/json"

	"fmt"
)

func NewAggregationLacpState() *AggregationLacpState {
	new := &AggregationLacpState{}
	return new
}

func (obj AggregationLacpState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var Obj AggregationLacpState
	var err error
	if err = json.Unmarshal(body, &Obj); err != nil {
		fmt.Println("### AggregationLacpState create called, unmarshal failed", Obj, err)
	}
	return Obj, err
}
