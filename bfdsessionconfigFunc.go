package models
import (
	 "encoding/json"
	"fmt"
)
func NewBfdSessionConfig() *BfdSessionConfig {
	newObj := &BfdSessionConfig{}
	return newObj
}

func (obj BfdSessionConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### BfdSessionConfig called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
