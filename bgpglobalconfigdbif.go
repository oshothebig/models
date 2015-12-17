package models

import (
	"database/sql"
	"fmt"
	"strings"
)

func (obj BGPGlobalConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BGPGlobalConfig " +
		"( " +
		" RouterId TEXT " +
		" AS INTEGER " +
		"PRIMARY KEY(RouterId) ) "

	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPGlobalConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BGPGlobalConfig (RouterId, AS) VALUES (%v, %v);",
		obj.RouterId, obj.AS)
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

func (obj BGPGlobalConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKey(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BGPGlobalConfig with key", objKey, "failed with error", err)
		return nil, err
	}

	dbCmd := "delete from BGPGlobalConfig where " + sqlKey
	fmt.Println("### DB Deleting BGPGlobalConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPGlobalConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (BGPGlobalConfig, error) {
	sqlKey, err := obj.GetSqlKey(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)
		return nil, err
	}

	dbCmd := "query from BGPGlobalConfig where " + sqlKey
	fmt.Println("### DB Get BGPGlobalConfig\n")
	obj, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return obj, err
}

func (obj BGPGlobalConfig) GetKey() (string, error) {
	key := string(obj.RouterId)
	return key, nil
}

func (obj BGPGlobalConfig) GetSqlKey(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "RouterId = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}
