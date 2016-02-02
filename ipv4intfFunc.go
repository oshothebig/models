package models
import (
	 "encoding/json"
	"fmt"
)
func NewIPv4Intf() *IPv4Intf {
	newObj := &IPv4Intf{}
	return newObj
}

func (obj IPv4Intf) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### IPv4Intf called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
