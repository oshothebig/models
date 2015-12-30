package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj PortIntfConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS PortIntfConfig " +
		"( " +
		"PortNum INTEGER, " +
		"Name TEXT, " +
		"Description TEXT, " +
		"Type TEXT, " +
		"AdminState TEXT, " +
		"OperState TEXT, " +
		"MacAddr TEXT, " +
		"Speed INTEGER, " +
		"Duplex TEXT, " +
		"Autoneg TEXT, " +
		"MediaType TEXT, " +
		"Mtu INTEGER, " +
		"PRIMARY KEY(PortNum) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj PortIntfConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO PortIntfConfig (PortNum, Name, Description, Type, AdminState, OperState, MacAddr, Speed, Duplex, Autoneg, MediaType, Mtu) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.PortNum, obj.Name, obj.Description, obj.Type, obj.AdminState, obj.OperState, obj.MacAddr, obj.Speed, obj.Duplex, obj.Autoneg, obj.MediaType, obj.Mtu)
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

func (obj PortIntfConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for PortIntfConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from PortIntfConfig where " + sqlKey
	fmt.Println("### DB Deleting PortIntfConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj PortIntfConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object PortIntfConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return object, err
	}

	dbCmd := "select * from PortIntfConfig where " + sqlKey
	fmt.Println("### DB Get PortIntfConfig\n")
	err = dbHdl.QueryRow(dbCmd).Scan(&object.PortNum, &object.Name, &object.Description, &object.Type, &object.AdminState, &object.OperState, &object.MacAddr, &object.Speed, &object.Duplex, &object.Autoneg, &object.MediaType, &object.Mtu)
	return object, err
}

func (obj PortIntfConfig) GetKey() (string, error) {
	key := string(obj.PortNum)
	return key, nil
}

func (obj PortIntfConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "PortNum = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj PortIntfConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]byte, error) {
	dbStruct := dbObj.(PortIntfConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbStruct)
	attrIds := make([]byte, objTyp.NumField())
	for i := 0; i < objTyp.NumField(); i++ {
		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if objVal.Kind() == reflect.Int {
			if (int(objVal.Int()) != 0) && (int(objVal.Int()) != int(dbObjVal.Int())) {
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
func (obj PortIntfConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []byte) (ConfigObj, error) {
	var mergedStruct PortIntfConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedStruct)
	for i := 1; i < objTyp.NumField(); i++ {
		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if attrSet[i] == 1 {
			if dbObjField.Kind() == reflect.Int {
				mergedObjVal.Elem().Field(i).SetInt(objField.Int())
			} else {
				mergedObjVal.Elem().Field(i).SetString(objField.String())
			}
		} else {
			if dbObjField.Kind() == reflect.Int {
				mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
			} else {
				mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
			}
		}
	}
	return mergedStruct, nil
}
func (obj PortIntfConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbStruct := dbObj.(PortIntfConfig)
	objKey, err := dbStruct.GetKey()
	objSqlKey, err := dbStruct.GetSqlKeyStr(objKey)
	dbCmd := "update " + "PortIntfConfig" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	for i := 0; i < objTyp.NumField(); i++ {
		if attrSet[i] == 1 {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int {
				fieldSqlStr = fmt.Sprintf(" %s = %d ", fieldTyp.Name, int(fieldVal.Int()))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = %s ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
