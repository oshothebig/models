package models

import (
	"encoding/json"
	"fmt"
)
type PolicyPrefix struct {
	IpPrefix string
	/*
	   Defines a range for the masklength, or 'exact' if
	   the prefix has an exact length.

	   Example: 10.3.192.0/21 through 10.3.192.0/24 would be
	   expressed as prefix: 10.3.192.0/21,
	   masklength-range: 21..24.

	   Example: 10.3.192.0/21 would be expressed as
	   prefix: 10.3.192.0/21,
	   masklength-range: exact
	*/
	MaskLengthRange string
}

type PolicyPrefixSet struct {
	ConfigObj
	PrefixSetName string `SNAPROUTE: "KEY"`
	/*
	   List of prefix expressions that are part of the set
	*/
	IpPrefixList [ ]PolicyPrefix
}

func (obj PolicyPrefixSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyPrefixSet PolicyPrefixSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyPrefixSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyPrefixSet from Json", body)
		}
	}
	return policyPrefixSet, err
}

type PolicyDstIpMatchPrefixSetCondition struct {
	//yang_name: prefix-set class: leaf
	PrefixSet string
	//yang_name: match-set-options class: leaf
	Prefix PolicyPrefix
}



