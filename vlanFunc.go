package models
import (
	 "encoding/json"
	"fmt"
)
func NewVlan() *Vlan {
	newObj := &Vlan{}
	return newObj
}

func (obj Vlan) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### Vlan called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }
