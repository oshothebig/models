package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj Dot1dStpPortEntryConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS Dot1dStpPortEntryConfig " +
		"( " +
		"Dot1dStpPortKey INTEGER, " +
		"Dot1dBrgIfIndex INTEGER, " +
		"Dot1dStpPortPriority INTEGER, " +
		"Dot1dStpPortEnable INTEGER, " +
		"Dot1dStpPortPathCost INTEGER, " +
		"Dot1dStpPortPathCost32 INTEGER, " +
		"Dot1dStpPortProtocolMigration INTEGER, " +
		"Dot1dStpPortAdminPointToPoint INTEGER, " +
		"Dot1dStpPortAdminEdgePort INTEGER, " +
		"Dot1dStpPortAdminPathCost INTEGER, " +
		"PRIMARY KEY(Dot1dStpPortKey) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj Dot1dStpPortEntryConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO Dot1dStpPortEntryConfig (Dot1dStpPortKey, Dot1dBrgIfIndex, Dot1dStpPortPriority, Dot1dStpPortEnable, Dot1dStpPortPathCost, Dot1dStpPortPathCost32, Dot1dStpPortProtocolMigration, Dot1dStpPortAdminPointToPoint, Dot1dStpPortAdminEdgePort, Dot1dStpPortAdminPathCost) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.Dot1dStpPortKey, obj.Dot1dBrgIfIndex, obj.Dot1dStpPortPriority, obj.Dot1dStpPortEnable, obj.Dot1dStpPortPathCost, obj.Dot1dStpPortPathCost32, obj.Dot1dStpPortProtocolMigration, obj.Dot1dStpPortAdminPointToPoint, obj.Dot1dStpPortAdminEdgePort, obj.Dot1dStpPortAdminPathCost)
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

func (obj Dot1dStpPortEntryConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for Dot1dStpPortEntryConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from Dot1dStpPortEntryConfig where " + sqlKey
	fmt.Println("### DB Deleting Dot1dStpPortEntryConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj Dot1dStpPortEntryConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object Dot1dStpPortEntryConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from Dot1dStpPortEntryConfig where " + sqlKey
	err = dbHdl.QueryRow(dbCmd).Scan(&object.Dot1dStpPortKey, &object.Dot1dBrgIfIndex, &object.Dot1dStpPortPriority, &object.Dot1dStpPortEnable, &object.Dot1dStpPortPathCost, &object.Dot1dStpPortPathCost32, &object.Dot1dStpPortProtocolMigration, &object.Dot1dStpPortAdminPointToPoint, &object.Dot1dStpPortAdminEdgePort, &object.Dot1dStpPortAdminPathCost)
	fmt.Println("### DB Get Dot1dStpPortEntryConfig\n", err)
	return object, err
}

func (obj Dot1dStpPortEntryConfig) GetKey() (string, error) {
	key := string(obj.Dot1dStpPortKey)
	return key, nil
}

func (obj Dot1dStpPortEntryConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Dot1dStpPortKey = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *Dot1dStpPortEntryConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []*Dot1dStpPortEntryConfig, e error) {
	dbCmd := "select * from Dot1dStpPortEntryConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'Dot1dStpPortEntryConfig' with error Dot1dStpPortEntryConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()

	for rows.Next() {

		object := new(Dot1dStpPortEntryConfig)
		if err = rows.Scan(&object.Dot1dStpPortKey, &object.Dot1dBrgIfIndex, &object.Dot1dStpPortPriority, &object.Dot1dStpPortEnable, &object.Dot1dStpPortPathCost, &object.Dot1dStpPortPathCost32, &object.Dot1dStpPortProtocolMigration, &object.Dot1dStpPortAdminPointToPoint, &object.Dot1dStpPortAdminEdgePort, &object.Dot1dStpPortAdminPathCost); err != nil {

			fmt.Println("Db method Scan failed when interating over Dot1dStpPortEntryConfig")
		}
		objList = append(objList, object)
	}
	return objList, nil
}
func (obj Dot1dStpPortEntryConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(Dot1dStpPortEntryConfig)
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

func (obj Dot1dStpPortEntryConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedDot1dStpPortEntryConfig Dot1dStpPortEntryConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedDot1dStpPortEntryConfig)
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
	return mergedDot1dStpPortEntryConfig, nil
}

func (obj Dot1dStpPortEntryConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbDot1dStpPortEntryConfig := dbObj.(Dot1dStpPortEntryConfig)
	objKey, err := dbDot1dStpPortEntryConfig.GetKey()
	objSqlKey, err := dbDot1dStpPortEntryConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "Dot1dStpPortEntryConfig" + " set"
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
