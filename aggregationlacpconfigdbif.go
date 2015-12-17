package models

import (
	"database/sql"
	"fmt"
	"strings"
)

func (obj AggregationLacpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationLacpConfig " +
		"( " +
		" Description TEXT " +
		" MinLinks INTEGER " +
		" SystemPriority INTEGER " +
		" NameKey TEXT " +
		" Interval INTEGER " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" SystemIdMac TEXT " +
		" LagType INTEGER " +
		" Type TEXT " +
		" LacpMode INTEGER " +
		"PRIMARY KEY(NameKey) ) "

	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj AggregationLacpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationLacpConfig (Description, MinLinks, SystemPriority, NameKey, Interval, Enabled, Mtu, SystemIdMac, LagType, Type, LacpMode) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.MinLinks, obj.SystemPriority, obj.NameKey, obj.Interval, obj.Enabled, obj.Mtu, obj.SystemIdMac, obj.LagType, obj.Type, obj.LacpMode)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
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
	_, err = ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj AggregationLacpConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (AggregationLacpConfig, error) {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return AggregationLacpConfig{}, err
	}

	dbCmd := "query from AggregationLacpConfig where " + sqlKey
	fmt.Println("### DB Get AggregationLacpConfig\n")
	sqlobj, err2 := ExecuteSQLStmt(dbCmd, dbHdl)
	return sqlobj, err2
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
