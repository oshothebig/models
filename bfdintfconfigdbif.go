package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"utils/dbutils"
)

func (obj BfdIntfConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BfdIntfConfig " +
		"( " +
		"Interface INTEGER, " +
		"LocalMultiplier INTEGER, " +
		"DesiredMinTxInterval INTEGER, " +
		"RequiredMinRxInterval INTEGER, " +
		"RequiredMinEchoRxInterval INTEGER, " +
		"DemandEnabled INTEGER, " +
		"AuthenticationEnabled INTEGER, " +
		"AuthType INTEGER, " +
		"AuthKeyId INTEGER, " +
		"AuthData TEXT, " +
		"PRIMARY KEY(Interface) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BfdIntfConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BfdIntfConfig (Interface, LocalMultiplier, DesiredMinTxInterval, RequiredMinRxInterval, RequiredMinEchoRxInterval, DemandEnabled, AuthenticationEnabled, AuthType, AuthKeyId, AuthData) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.Interface, obj.LocalMultiplier, obj.DesiredMinTxInterval, obj.RequiredMinRxInterval, obj.RequiredMinEchoRxInterval, dbutils.ConvertBoolToInt(obj.DemandEnabled), dbutils.ConvertBoolToInt(obj.AuthenticationEnabled), obj.AuthType, obj.AuthKeyId, obj.AuthData)
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

func (obj BfdIntfConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BfdIntfConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from BfdIntfConfig where " + sqlKey
	fmt.Println("### DB Deleting BfdIntfConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BfdIntfConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object BfdIntfConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from BfdIntfConfig where " + sqlKey
	var tmp5 string
	var tmp6 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.Interface, &object.LocalMultiplier, &object.DesiredMinTxInterval, &object.RequiredMinRxInterval, &object.RequiredMinEchoRxInterval, &tmp5, &tmp6, &object.AuthType, &object.AuthKeyId, &object.AuthData)
	fmt.Println("### DB Get BfdIntfConfig\n", err)
	object.DemandEnabled = dbutils.ConvertStrBoolIntToBool(tmp5)
	object.AuthenticationEnabled = dbutils.ConvertStrBoolIntToBool(tmp6)
	return object, err
}

func (obj BfdIntfConfig) GetKey() (string, error) {
	key := string(strconv.Itoa(int(obj.Interface)))
	return key, nil
}

func (obj BfdIntfConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Interface = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *BfdIntfConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []*BfdIntfConfig, e error) {
	dbCmd := "select * from BfdIntfConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'BfdIntfConfig' with error BfdIntfConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()

	var tmp5 string
	var tmp6 string
	for rows.Next() {

		object := new(BfdIntfConfig)
		if err = rows.Scan(&object.Interface, &object.LocalMultiplier, &object.DesiredMinTxInterval, &object.RequiredMinRxInterval, &object.RequiredMinEchoRxInterval, &tmp5, &tmp6, &object.AuthType, &object.AuthKeyId, &object.AuthData); err != nil {

			fmt.Println("Db method Scan failed when interating over BfdIntfConfig")
		}
		object.DemandEnabled = dbutils.ConvertStrBoolIntToBool(tmp5)
		object.AuthenticationEnabled = dbutils.ConvertStrBoolIntToBool(tmp6)
		objList = append(objList, object)
	}
	return objList, nil
}
func (obj BfdIntfConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(BfdIntfConfig)
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

func (obj BfdIntfConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedBfdIntfConfig BfdIntfConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedBfdIntfConfig)
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
	return mergedBfdIntfConfig, nil
}

func (obj BfdIntfConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbBfdIntfConfig := dbObj.(BfdIntfConfig)
	objKey, err := dbBfdIntfConfig.GetKey()
	objSqlKey, err := dbBfdIntfConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "BfdIntfConfig" + " set"
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
			} else if fieldVal.Kind() == reflect.Bool {
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
