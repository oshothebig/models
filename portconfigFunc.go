package models
import (
	 "encoding/json"
	"fmt"
)
func NewPortConfig() *PortConfig {
	newObj := &PortConfig{}
	return newObj
}

func (obj PortConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### PortConfig called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
