package models
import (
	 "encoding/json"
	"fmt"
)
func NewDhcpRelayGlobalConfig() *DhcpRelayGlobalConfig {
	newObj := &DhcpRelayGlobalConfig{}
	return newObj
}

func (obj DhcpRelayGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### DhcpRelayGlobalConfig called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
