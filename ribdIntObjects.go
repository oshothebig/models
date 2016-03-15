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
	BaseObj
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

type PolicyConditionState struct {
	BaseObj
	Name           string `SNAPROUTE: "KEY"`
	ConditionInfo  string
	PolicyStmtList []string
}

func (obj PolicyConditionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyConditionState PolicyConditionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyConditionState); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyConditionState from Json", body)
		}
	}
	return policyConditionState, err
}

type PolicyActionState struct {
	BaseObj
	Name           string `SNAPROUTE: "KEY"`
	ActionInfo     string
	PolicyStmtList []string
}

func (obj PolicyActionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyActionState PolicyActionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyActionState); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyActionState from Json", body)
		}
	}
	return policyActionState, err
}

type PolicyStmtConfig struct {
	BaseObj
	Name string                   `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Statement Name"`
	MatchConditions string        `DESCRIPTION "Specifies whether to match all/any of the conditions of this policy statement"`
	Conditions      []string      `DESCRIPTION "List of conditions added to this policy statement"`
	Actions         []string      `DESCRIPTION "List of actions added to this policy statement"`
}


func (obj PolicyStmtConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyStmt PolicyStmtConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyStmtConfig from Json", body)
		}
	}
	return policyStmt, err
}

type PolicyStmtState struct {
	BaseObj
	Name string                    `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyStmtState"`
	MatchConditions string         `DESCRIPTION "Specifies whether to match all/any of the conditions of this policy statement"`
	Conditions      []string       `DESCRIPTION "List of conditions added to this policy statement"`
	Actions         []string       `DESCRIPTION "List of actions added to this policy statement"`
	PolicyList      []string       `DESCRIPTION "List of policies using this policy statement"`
}

func (obj PolicyStmtState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyStmt PolicyStmtState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyStmtState from Json", body)
		}
	}
	return policyStmt, err
}

type PolicyDefinitionStmtPrecedence struct {
	Precedence int
	Statement  string
}


type PolicyDefinitionConfig struct {
	BaseObj
	Name       string               `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Name"`
	Precedence int32                  `DESCRIPTION "Priority of the policy w.r.t other policies configured"`
	MatchType  string               `DESCRIPTION "Specifies whether to match all/any of the statements within this policy"`
	StatementList [] PolicyDefinitionStmtPrecedence `DESCRIPTION "Specifies list of statements along with their precedence order."`
}


func (obj PolicyDefinitionConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinition PolicyDefinitionConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinition); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionConfig from Json", body)
		}
	}
	return policyDefinition, err
}

type PolicyDefinitionState struct {
	BaseObj
	Name         string              `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyDefinitionState"`
	HitCounter   int                 `DESCRIPTION "Number of times this policy has been applied"`
	IpPrefixList []string            `DESCRIPTION "List of networks/IP Prefixes this policy has been applied on to."`
}

func (obj PolicyDefinitionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmt PolicyDefinitionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionState from Json", body)
		}
	}
	return policyDefinitionStmt, err
}
