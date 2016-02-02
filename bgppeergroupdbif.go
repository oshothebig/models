package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj BGPPeerGroup) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BGPPeerGroup " +
		"( " +
		"PeerAS INTEGER, " +
		"LocalAS INTEGER, " +
		"AuthPassword TEXT, " +
		"Description TEXT, " +
		"Name TEXT, " +
		"RouteReflectorClusterId INTEGER, " +
		"RouteReflectorClient INTEGER, " +
		"MultiHopEnable INTEGER, " +
		"MultiHopTTL INTEGER, " +
		"ConnectRetryTime INTEGER, " +
		"HoldTime INTEGER, " +
		"KeepaliveTime INTEGER, " +
		"PRIMARY KEY(Name) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPPeerGroup) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BGPPeerGroup (PeerAS, LocalAS, AuthPassword, Description, Name, RouteReflectorClusterId, RouteReflectorClient, MultiHopEnable, MultiHopTTL, ConnectRetryTime, HoldTime, KeepaliveTime) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.PeerAS, obj.LocalAS, obj.AuthPassword, obj.Description, obj.Name, obj.RouteReflectorClusterId, dbutils.ConvertBoolToInt(obj.RouteReflectorClient), dbutils.ConvertBoolToInt(obj.MultiHopEnable), obj.MultiHopTTL, obj.ConnectRetryTime, obj.HoldTime, obj.KeepaliveTime)
	fmt.Println("**** Create Object called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	} else {
		objectId, err = result.LastInsertId()
		if err != nil {
			fmt.Println("### Failed to return last object id", err)
		}

	}
	return objectId, err
}

func (obj BGPPeerGroup) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BGPPeerGroup with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from BGPPeerGroup where " + sqlKey
	fmt.Println("### DB Deleting BGPPeerGroup\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPPeerGroup) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object BGPPeerGroup
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from BGPPeerGroup where " + sqlKey
	var tmp6 string
	var tmp7 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.PeerAS, &object.LocalAS, &object.AuthPassword, &object.Description, &object.Name, &object.RouteReflectorClusterId, &tmp6, &tmp7, &object.MultiHopTTL, &object.ConnectRetryTime, &object.HoldTime, &object.KeepaliveTime)
	fmt.Println("### DB Get BGPPeerGroup\n", err)
	object.RouteReflectorClient = dbutils.ConvertStrBoolIntToBool(tmp6)
	object.MultiHopEnable = dbutils.ConvertStrBoolIntToBool(tmp7)
	return object, err
}

func (obj BGPPeerGroup) GetKey() (string, error) {
	key := string(obj.Name)
	return key, nil
}

func (obj BGPPeerGroup) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Name = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *BGPPeerGroup) GetAllObjFromDb(dbHdl *sql.DB) (objList []*BGPPeerGroup, e error) {
	dbCmd := "select * from BGPPeerGroup"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'BGPPeerGroup' with error BGPPeerGroup", dbCmd, err))
		return objList, err
	}

	defer rows.Close()

	var tmp6 string
	var tmp7 string
	for rows.Next() {

		object := new(BGPPeerGroup)
		if err = rows.Scan(&object.PeerAS, &object.LocalAS, &object.AuthPassword, &object.Description, &object.Name, &object.RouteReflectorClusterId, &tmp6, &tmp7, &object.MultiHopTTL, &object.ConnectRetryTime, &object.HoldTime, &object.KeepaliveTime); err != nil {

			fmt.Println("Db method Scan failed when interating over BGPPeerGroup")
		}
		object.RouteReflectorClient = dbutils.ConvertStrBoolIntToBool(tmp6)
		object.MultiHopEnable = dbutils.ConvertStrBoolIntToBool(tmp7)
		objList = append(objList, object)
	}
	return objList, nil
}
func (obj BGPPeerGroup) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(BGPPeerGroup)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]bool, objTyp.NumField())
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		fieldTyp := objTyp.Field(i)
		if fieldTyp.Anonymous {
			continue
		}

		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if _, ok := updateKeys[fieldTyp.Name]; ok {
			if objVal.Kind() == reflect.Int {
				if int(objVal.Int()) != int(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int8 {
				if int8(objVal.Int()) != int8(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int16 {
				if int16(objVal.Int()) != int16(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int32 {
				if int32(objVal.Int()) != int32(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int64 {
				if int64(objVal.Int()) != int64(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint {
				if uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint8 {
				if uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint16 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint32 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint64 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Bool {
				if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
					attrIds[idx] = true
				}
			} else {
				if objVal.String() != dbObjVal.String() {
					attrIds[idx] = true
				}
			}
			if attrIds[idx] {
				fmt.Println("attribute changed ", fieldTyp.Name)
			}
		}
		idx++

	}
	return attrIds[:idx], nil
}

func (obj BGPPeerGroup) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedBGPPeerGroup BGPPeerGroup
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedBGPPeerGroup)
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if attrSet[idx] {
			if dbObjField.Kind() == reflect.Int ||
				dbObjField.Kind() == reflect.Int8 ||
				dbObjField.Kind() == reflect.Int16 ||
				dbObjField.Kind() == reflect.Int32 ||
				dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(objField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint8 ||
				dbObjField.Kind() == reflect.Uint16 ||
				dbObjField.Kind() == reflect.Uint32 ||
				dbObjField.Kind() == reflect.Uint64 {
				mergedObjVal.Elem().Field(i).SetUint(objField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(objField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(objField.String())
			}
		} else {
			if dbObjField.Kind() == reflect.Int ||
				dbObjField.Kind() == reflect.Int8 ||
				dbObjField.Kind() == reflect.Int16 ||
				dbObjField.Kind() == reflect.Int32 ||
				dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint8 ||
				dbObjField.Kind() == reflect.Uint16 ||
				dbObjField.Kind() == reflect.Uint32 {
				mergedObjVal.Elem().Field(i).SetUint(dbObjField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(dbObjField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
			}
		}
		idx++

	}
	return mergedBGPPeerGroup, nil
}

func (obj BGPPeerGroup) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbBGPPeerGroup := dbObj.(BGPPeerGroup)
	objKey, err := dbBGPPeerGroup.GetKey()
	objSqlKey, err := dbBGPPeerGroup.GetSqlKeyStr(objKey)
	dbCmd := "update " + "BGPPeerGroup" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		if attrSet[idx] {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int ||
				fieldVal.Kind() == reflect.Int8 ||
				fieldVal.Kind() == reflect.Int16 ||
				fieldVal.Kind() == reflect.Int32 ||
				fieldVal.Kind() == reflect.Int64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
			} else if fieldVal.Kind() == reflect.Uint ||
				fieldVal.Kind() == reflect.Uint8 ||
				fieldVal.Kind() == reflect.Uint16 ||
				fieldVal.Kind() == reflect.Uint32 ||
				fieldVal.Kind() == reflect.Uint64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Uint()))
			} else if objVal.Kind() == reflect.Bool {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, dbutils.ConvertBoolToInt(bool(fieldVal.Bool())))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
		idx++
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
