package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj SubinterfaceVlanConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceVlanConfig " +
		"( " +
		"IndexKey TEXT, " +
		"Mtu INTEGER, " +
		"Name TEXT, " +
		"Type TEXT, " +
		"Unnumbered INTEGER, " +
		"Enabled INTEGER, " +
		"Description TEXT, " +
		"GlobalVlanId TEXT, " +
		"VlanId_VlanId INTEGER, " +
		"VlanId_VlanId_QinqId TEXT, " +
		"PRIMARY KEY(IndexKey) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj SubinterfaceVlanConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceVlanConfig (IndexKey, Mtu, Name, Type, Unnumbered, Enabled, Description, GlobalVlanId, VlanId_VlanId, VlanId_VlanId_QinqId) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.IndexKey, obj.Mtu, obj.Name, obj.Type, dbutils.ConvertBoolToInt(obj.Unnumbered), dbutils.ConvertBoolToInt(obj.Enabled), obj.Description, obj.GlobalVlanId, obj.VlanId_VlanId, obj.VlanId_VlanId_QinqId)
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

func (obj SubinterfaceVlanConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for SubinterfaceVlanConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from SubinterfaceVlanConfig where " + sqlKey
	fmt.Println("### DB Deleting SubinterfaceVlanConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj SubinterfaceVlanConfig) GetObjectFromDb(objSqlKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object SubinterfaceVlanConfig
	dbCmd := "select * from SubinterfaceVlanConfig where " + objSqlKey
	var tmp4 string
	var tmp5 string
	err := dbHdl.QueryRow(dbCmd).Scan(&object.IndexKey, &object.Mtu, &object.Name, &object.Type, &tmp4, &tmp5, &object.Description, &object.GlobalVlanId, &object.VlanId_VlanId, &object.VlanId_VlanId_QinqId)
	fmt.Println("### DB Get SubinterfaceVlanConfig\n", err)
	object.Unnumbered = dbutils.ConvertStrBoolIntToBool(tmp4)
	object.Enabled = dbutils.ConvertStrBoolIntToBool(tmp5)
	return object, err
}

func (obj SubinterfaceVlanConfig) GetKey() (string, error) {
	key := string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceVlanConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "IndexKey = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj SubinterfaceVlanConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]byte, error) {
	dbV4Route := dbObj.(SubinterfaceVlanConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]byte, objTyp.NumField())
	for i := 0; i < objTyp.NumField(); i++ {
		fieldTyp := objTyp.Field(i)
		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if _, ok := updateKeys[fieldTyp.Name]; ok {
			if objVal.Kind() == reflect.Int {
				if int(objVal.Int()) != int(dbObjVal.Int()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Int8 {
				if int8(objVal.Int()) != int8(dbObjVal.Int()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Int16 {
				if int16(objVal.Int()) != int16(dbObjVal.Int()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Int32 {
				if int32(objVal.Int()) != int32(dbObjVal.Int()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Int64 {
				if int64(objVal.Int()) != int64(dbObjVal.Int()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Uint {
				if uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Uint8 {
				if uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Uint16 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Uint32 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Uint64 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[i] = 1
				}
			} else if objVal.Kind() == reflect.Bool {
				if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
					attrIds[i] = 1
				}
			} else {
				if objVal.String() != dbObjVal.String() {
					attrIds[i] = 1
				}
			}
			if attrIds[i] == 1 {
				fmt.Println("attribute changed ", fieldTyp.Name)
			}
		}
	}
	return attrIds, nil
}

func (obj SubinterfaceVlanConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []byte) (ConfigObj, error) {
	var mergedSubinterfaceVlanConfig SubinterfaceVlanConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedSubinterfaceVlanConfig)
	for i := 1; i < objTyp.NumField(); i++ {
		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if attrSet[i] == 1 {
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
	}
	return mergedSubinterfaceVlanConfig, nil
}

func (obj SubinterfaceVlanConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbSubinterfaceVlanConfig := dbObj.(SubinterfaceVlanConfig)
	objKey, err := dbSubinterfaceVlanConfig.GetKey()
	objSqlKey, err := dbSubinterfaceVlanConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "SubinterfaceVlanConfig" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	for i := 0; i < objTyp.NumField(); i++ {
		if attrSet[i] == 1 {
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
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
