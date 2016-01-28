package models
import (
	 "encoding/json"
	"fmt"
)
func NewBfdGlobalConfig() *BfdGlobalConfig {
	newObj := &BfdGlobalConfig{}
	return newObj
}

func (obj BfdGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### BfdGlobalConfig called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
