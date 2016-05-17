//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __  
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  | 
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  | 
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   | 
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  | 
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__| 
//                                                                                                           

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
	baseObj
	PrefixSetName string `SNAPROUTE: "KEY"`
	/*
	   List of prefix expressions that are part of the set
	*/
	IpPrefixList []PolicyPrefix
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
