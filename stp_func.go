package models

import (
	"encoding/json"

	"fmt"
)

func (obj Dot1dStpPortEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpPortEntryConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
func (obj Dot1dStpPortEntryStateCountersFsmStates) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpPortEntryStateCountersFsmStates called, unmarshal failed", obj, err)
		}
	}
	return obj, err
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
func (obj Dot1dStpBridgeState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### Dot1dStpBridgeState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
