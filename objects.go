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
	"github.com/garyburd/redigo/redis"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	StoreObjectInDb(dbHdl redis.Conn) error
	DeleteObjectFromDb(dbHdl redis.Conn) error
	GetKey() string
	GetObjectFromDbByKey(objKey string, dbHdl redis.Conn) (ConfigObj, error)
	GetObjectFromDb(objKey string, dbHdl redis.Conn) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObjForPatchUpdate(dbObj ConfigObj, patchOpInfoSlice []PatchOpInfo) (ConfigObj, []bool, error) 
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl redis.Conn) error
	GetAllObjFromDb(dbHdl redis.Conn) ([]ConfigObj, error)
	GetBulkObjFromDb(startIndex int64, count int64, dbHdl redis.Conn) (error, int64, int64, bool, []ConfigObj)
}
type PatchOpInfo struct {
	Op string
	Path string
	Value string
}
type baseObj struct {
}

func (obj baseObj) UnmarshalObject(data []byte) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) StoreObjectInDb(dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) DeleteObjectFromDb(dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) GetKey() string {
	return ""
}
func (obj baseObj) GetObjectFromDbByKey(objKey string, dbHdl redis.Conn) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) GetObjectFromDb(objKey string, dbHdl redis.Conn) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	return nil, nil
}
func (obj baseObj) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) MergeDbAndConfigObjForPatchUpdate(dbObj ConfigObj, patchOpInfoSlice []PatchOpInfo) (ConfigObj, []bool, error) {
	return nil,nil, nil
}
func (obj baseObj) UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) GetAllObjFromDb(dbHdl redis.Conn) ([]ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) GetBulkObjFromDb(startIndex int64, count int64, dbHdl redis.Conn) (error, int64, int64, bool, []ConfigObj) {
	return nil, 0, 0, false, nil
}

type SystemStatusState struct {
	baseObj
	Name           string        `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "Name of the system"`
	Ready          bool          `DESCRIPTION: "System is ready to accept api calls"`
	Reason         string        `DESCRIPTION: "Reason if system not ready"`
	UpTime         string        `DESCRIPTION: "Uptime of this system"`
	NumCreateCalls string        `DESCRIPTION: "Number of create api calls made"`
	NumDeleteCalls string        `DESCRIPTION: "Number of delete api calls made"`
	NumUpdateCalls string        `DESCRIPTION: "Number of update api calls made"`
	NumGetCalls    string        `DESCRIPTION: "Number of get api calls made"`
	NumActionCalls string        `DESCRIPTION: "Number of action api calls made"`
	FlexDaemons    []DaemonState `DESCRIPTION: "Daemon states"`
}

func (obj SystemStatusState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var systemStatus SystemStatusState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &systemStatus); err != nil {
			fmt.Println("### Trouble in unmarshaling SystemStatus from Json", body)
		}
	}

	return systemStatus, err
}

func (obj SystemStatusState) GetKey() string {
	return ""
}

type RepoInfo struct {
	Name   string `DESCRIPTION: "Name of the git repo"`
	Sha1   string `DESCRIPTION: "Git commit Sha1"`
	Branch string `DESCRIPTION: "Branch name"`
	Time   string `DESCRIPTION: "Build time"`
}

type SystemSwVersionState struct {
	baseObj
	FlexswitchVersion string     `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "Flexswitch version"`
	Repos             []RepoInfo `DESCRIPTION: "Git repo details"`
}

func (obj SystemSwVersionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var systemSwVersion SystemSwVersionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &systemSwVersion); err != nil {
			fmt.Println("### Trouble in unmarshaling SystemSwVersion from Json", body)
		}
	}

	return systemSwVersion, err
}

func (obj SystemSwVersionState) GetKey() string {
	return ""
}
