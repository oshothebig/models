package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj AggregationLacpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationLacpConfig " +
		"( " +
		"LagType INTEGER, " +
		"Description TEXT, " +
		"Enabled bool, " +
		"Mtu INTEGER, " +
		"MinLinks INTEGER, " +
		"Type TEXT, " +
		"NameKey TEXT, " +
		"Interval INTEGER, " +
		"LacpMode INTEGER, " +
		"SystemIdMac TEXT, " +
		"SystemPriority INTEGER, " +
		"PRIMARY KEY(NameKey) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj AggregationLacpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationLacpConfig (LagType, Description, Enabled, Mtu, MinLinks, Type, NameKey, Interval, LacpMode, SystemIdMac, SystemPriority) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.LagType, obj.Description, obj.Enabled, obj.Mtu, obj.MinLinks, obj.Type, obj.NameKey, obj.Interval, obj.LacpMode, obj.SystemIdMac, obj.SystemPriority)
	fmt.Println("**** Create Object called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}

	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}

	return objectId, err
}

func (obj AggregationLacpConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for AggregationLacpConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from AggregationLacpConfig where " + sqlKey
	fmt.Println("### DB Deleting AggregationLacpConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj AggregationLacpConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object AggregationLacpConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return object, err
	}

	dbCmd := "SELECT * from AggregationLacpConfig where " + sqlKey
	fmt.Println("### DB Get AggregationLacpConfig\n")
	err = dbHdl.QueryRow(dbCmd).Scan(&object.LagType, &object.Description, &object.Enabled, &object.Mtu, &object.MinLinks, &object.Type, &object.NameKey, &object.Interval, &object.LacpMode, &object.SystemIdMac, &object.SystemPriority)
	return object, err
}

func (obj AggregationLacpConfig) GetKey() (string, error) {
	key := string(obj.NameKey)
	return key, nil
}

func (obj AggregationLacpConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "NameKey = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj AggregationLacpConfig) CompareObjectsAndDiff(dbObj ConfigObj) ([]byte, error) {
	dbV4Route := dbObj.(AggregationLacpConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]byte, objTyp.NumField())
	for i := 0; i < objTyp.NumField(); i++ {
		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if objVal.Kind() == reflect.Int {
			if int(objVal.Int()) != 0 && int(objVal.Int()) != int(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int8 {
			if int8(objVal.Int()) != 0 && int8(objVal.Int()) != int8(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int16 {
			if int16(objVal.Int()) != 0 && int16(objVal.Int()) != int16(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int32 {
			if int32(objVal.Int()) != 0 && int32(objVal.Int()) != int32(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int64 {
			if int64(objVal.Int()) != 0 && int64(objVal.Int()) != int64(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Uint {
			if uint(objVal.Uint()) != 0 && uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Uint8 {
			if uint8(objVal.Uint()) != 0 && uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Uint16 {
			if uint16(objVal.Uint()) != 0 && uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Uint32 {
			if uint16(objVal.Uint()) != 0 && uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Uint64 {
			if uint16(objVal.Uint()) != 0 && uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Bool {
			if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
				attrIds[i] = 1
			}
		} else {
			if objVal.String() != "" && objVal.String() != dbObjVal.String() {
				attrIds[i] = 1
			}
		}
	}
	return attrIds, nil
}

func (obj AggregationLacpConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []byte) (ConfigObj, error) {
	var mergedAggregationLacpConfig AggregationLacpConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedAggregationLacpConfig)
	for i := 1; i < objTyp.NumField(); i++ {
		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if attrSet[i] == 1 {
			if dbObjField.Kind() == reflect.Int {
				mergedObjVal.Elem().Field(i).SetInt(objField.Int())
			} else if dbObjField.Kind() == reflect.Uint {
				mergedObjVal.Elem().Field(i).SetUint(objField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(objField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(objField.String())
			}
		} else {
			if dbObjField.Kind() == reflect.Int {
				mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
			} else if dbObjField.Kind() == reflect.Uint {
				mergedObjVal.Elem().Field(i).SetUint(dbObjField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(dbObjField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
			}
		}
	}
	return mergedAggregationLacpConfig, nil
}

func (obj AggregationLacpConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbAggregationLacpConfig := dbObj.(AggregationLacpConfig)
	objKey, err := dbAggregationLacpConfig.GetKey()
	objSqlKey, err := dbAggregationLacpConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "AggregationLacpConfig" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	for i := 0; i < objTyp.NumField(); i++ {
		if attrSet[i] == 1 {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
			} else if objVal.Kind() == reflect.Uint {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Uint()))
			} else if objVal.Kind() == reflect.Bool {
				fieldSqlStr = fmt.Sprintf(" %s = '%t' ", fieldTyp.Name, bool(fieldVal.Bool()))
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
