package models
import (
	 "encoding/json"
	"fmt"
)
func NewDhcpRelayIntfConfig() *DhcpRelayIntfConfig {
	newObj := &DhcpRelayIntfConfig{}
	return newObj
}

func (obj DhcpRelayIntfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### DhcpRelayIntfConfig called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
