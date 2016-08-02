package actions

import (
	"encoding/json"
	"fmt"
)

func (obj ApplyConfig) UnmarshalAction(body []byte) (ActionObj, error) {
	var err error
	fmt.Println("applyconfig unmarshalAction body:", body)
	if len(body) > 0 {
		fmt.Println("len(body)>0, so unmarshal")
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("###  called, unmarshal failed", obj, err)
		}
	}
	fmt.Println("ApplyConfig obj:", obj, " err:", err)
	return obj, err
}

func (obj ResetConfig) UnmarshalAction(body []byte) (ActionObj, error) {
	var err error

	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("###  called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}

func (obj SaveConfig) UnmarshalAction(body []byte) (ActionObj, error) {
	var err error

	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("###  called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}

func (obj Daemon) UnmarshalAction(body []byte) (ActionObj, error) {
	var err error

	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("###  called, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
