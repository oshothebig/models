package models

import (
	"encoding/json"
	"fmt"
)

func NewBfdIntfConfig() *BfdIntfConfig {
	newObj := &BfdIntfConfig{}
	return newObj
}

func (obj BfdIntfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### BfdIntfConfig called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}

func (obj BfdSessionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("### BfdSessionState called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
