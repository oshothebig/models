// BGPRoutingPolicy.go
package models

import (
	"encoding/json"
	"fmt"
)

/*type BGPPolicyPrefix struct {
	IpPrefix string
	MaskLengthRange string
}
*/
type BGPPolicyPrefixSet struct {
	BaseObj
	PrefixSetName string `SNAPROUTE: "KEY"`
	IpPrefixList  []PolicyPrefix
}

func (obj BGPPolicyPrefixSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyPrefixSet BGPPolicyPrefixSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyPrefixSet); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPPolicyPrefixSet from Json", body)
		}
	}
	return policyPrefixSet, err
}

/*
type BGPPolicyDstIpMatchPrefixSetCondition struct {
	//yang_name: prefix-set class: leaf
	PrefixSet string
	//yang_name: match-set-options class: leaf
    Prefix PolicyPrefix
}
*/

type BGPPolicyAggregateAction struct {
	GenerateASSet   bool
	SendSummaryOnly bool
}
