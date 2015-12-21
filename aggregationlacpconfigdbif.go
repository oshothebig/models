package models

import (
	"database/sql"
	"fmt"
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
	dbCmd := fmt.Sprintf("INSERT INTO AggregationLacpConfig (LagType, Description, Enabled, Mtu, MinLinks, Type, NameKey, Interval, LacpMode, SystemIdMac, SystemPriority) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
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

	dbCmd := "query from AggregationLacpConfig where " + sqlKey
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
	var arr []byte
	return arr, nil
}

func (obj AggregationLacpConfig) UpdateObjectInDb(dbV4Route ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	return nil
}
