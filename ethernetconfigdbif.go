package models

import (
	"database/sql"
	"utils/dbutils"
	"fmt"
	"strings"
)

func (obj EthernetConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetConfig " +
		"( " +
		" MacAddress TEXT " +
		" Description TEXT " +
		" AggregateId TEXT " +
		" NameKey TEXT " +
		" Enabled bool " +
		" Speed TEXT " +
		" Mtu INTEGER " +
		" DuplexMode INTEGER " +
		" EnableFlowControl bool " +
		" Auto bool " +
		" Type TEXT " +
		"PRIMARY KEY(NameKey) ) "

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj EthernetConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetConfig (MacAddress, Description, AggregateId, NameKey, Enabled, Speed, Mtu, DuplexMode, EnableFlowControl, Auto, Type) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.MacAddress, obj.Description, obj.AggregateId, obj.NameKey, obj.Enabled, obj.Speed, obj.Mtu, obj.DuplexMode, obj.EnableFlowControl, obj.Auto, obj.Type)
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
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return EthernetConfig{}, err
	}

	dbCmd := "query from EthernetConfig where " + sqlKey
	fmt.Println("### DB Get EthernetConfig\n")
	_, err2 := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	//return sqlobj, err2 // FIXME
   return EthernetConfig{}, err2
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
