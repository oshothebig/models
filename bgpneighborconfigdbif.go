package models

import (
	"database/sql"
	"fmt"
	"strings"
)

func (obj BGPNeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BGPNeighborConfig " +
		"( " +
		" NeighborAddress TEXT " +
		" Description TEXT " +
		" PeerAS INTEGER " +
		" LocalAS INTEGER " +
		" AuthPassword TEXT " +
		"PRIMARY KEY(NeighborAddress) ) "

	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPNeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BGPNeighborConfig (NeighborAddress, Description, PeerAS, LocalAS, AuthPassword) VALUES (%v, %v, %v, %v, %v);",
		obj.NeighborAddress, obj.Description, obj.PeerAS, obj.LocalAS, obj.AuthPassword)
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

func (obj BGPNeighborConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKey(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BGPNeighborConfig with key", objKey, "failed with error", err)
		return nil, err
	}

	dbCmd := "delete from BGPNeighborConfig where " + sqlKey
	fmt.Println("### DB Deleting BGPNeighborConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPNeighborConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (BGPNeighborConfig, error) {
	sqlKey, err := obj.GetSqlKey(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return nil, err
	}

	dbCmd := "query from BGPNeighborConfig where " + sqlKey
	fmt.Println("### DB Get BGPNeighborConfig\n")
	obj, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return obj, err
}

func (obj BGPNeighborConfig) GetKey() (string, error) {
	key := string(obj.NeighborAddress)
	return key, nil
}

func (obj BGPNeighborConfig) GetSqlKey(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "NeighborAddress = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}
