package models

import (
	"encoding/json"
	"fmt"
)

/*
 * Global DataStructure for BFD
 */
type DhcpRelayGlobalConfig struct {
	BaseObj
	Bfd    string `SNAPROUTE: "KEY"`
	Enable bool
}

func (obj BfdGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BfdGlobalConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BfdGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

/*
 * BFD Interface config
 */
type BfdIntfConfig struct {
	BaseObj
	Interface     int32 `SNAPROUTE: "KEY"`
	Mode          string
	TxMinInterval uint32
	RxMinInterval uint32
	Multiplier    uint32
	EchoFunction  bool
}

func (obj bfdIntfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BfdIntfConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BfdIntfConfig from Json", body)
		}
	}
	return gConf, err
}
