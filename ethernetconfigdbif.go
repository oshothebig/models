package models

import (
	"database/sql"
	"fmt"
	"strings"
	"utils/dbutils"
)

func (obj EthernetConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetConfig " +
		"( " +
		"NameKey TEXT, " +
		"Enabled bool, " +
		"Description TEXT, " +
		"Mtu INTEGER, " +
		"Type TEXT, " +
		"MacAddress TEXT, " +
		"DuplexMode INTEGER, " +
		"Auto bool, " +
		"Speed TEXT, " +
		"EnableFlowControl bool, " +
		"AggregateId TEXT, " +
		"PRIMARY KEY(NameKey) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj EthernetConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetConfig (NameKey, Enabled, Description, Mtu, Type, MacAddress, DuplexMode, Auto, Speed, EnableFlowControl, AggregateId) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.NameKey, obj.Enabled, obj.Description, obj.Mtu, obj.Type, obj.MacAddress, obj.DuplexMode, obj.Auto, obj.Speed, obj.EnableFlowControl, obj.AggregateId)
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

func (obj EthernetConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for EthernetConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from EthernetConfig where " + sqlKey
	fmt.Println("### DB Deleting EthernetConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj EthernetConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (EthernetConfig, error) {
	var object EthernetConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return object, err
	}

	dbCmd := "query from EthernetConfig where " + sqlKey
	fmt.Println("### DB Get EthernetConfig\n")
	err = dbHdl.QueryRow(dbCmd).Scan(&object.NameKey, &object.Enabled, &object.Description, &object.Mtu, &object.Type, &object.MacAddress, &object.DuplexMode, &object.Auto, &object.Speed, &object.EnableFlowControl, &object.AggregateId)
	return object, err
}

func (obj EthernetConfig) GetKey() (string, error) {
	key := string(obj.NameKey)
	return key, nil
}

func (obj EthernetConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "NameKey = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}
