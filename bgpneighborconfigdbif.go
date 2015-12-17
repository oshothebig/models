package models

import (
	"database/sql"
	"fmt"
	"strings"
	"utils/dbutils"
)

func (obj BGPNeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BGPNeighborConfig " +
		"( " +
		"PeerAS INTEGER, " +
		"LocalAS INTEGER, " +
		"AuthPassword TEXT, " +
		"Description TEXT, " +
		"NeighborAddress TEXT, " +
		"PRIMARY KEY(NeighborAddress) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPNeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BGPNeighborConfig (PeerAS, LocalAS, AuthPassword, Description, NeighborAddress) VALUES (%v, %v, %v, %v, %v);",
		obj.PeerAS, obj.LocalAS, obj.AuthPassword, obj.Description, obj.NeighborAddress)
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

func (obj BGPNeighborConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BGPNeighborConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from BGPNeighborConfig where " + sqlKey
	fmt.Println("### DB Deleting BGPNeighborConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPNeighborConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (BGPNeighborConfig, error) {
	var object BGPNeighborConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return object, err
	}

	dbCmd := "query from BGPNeighborConfig where " + sqlKey
	fmt.Println("### DB Get BGPNeighborConfig\n")
	err = dbHdl.QueryRow(dbCmd).Scan(&object.PeerAS, &object.LocalAS, &object.AuthPassword, &object.Description, &object.NeighborAddress)
	return object, err
}

func (obj BGPNeighborConfig) GetKey() (string, error) {
	key := string(obj.NeighborAddress)
	return key, nil
}

func (obj BGPNeighborConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "NeighborAddress = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}
