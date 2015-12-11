package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	//"strconv"
)

func ExecuteSQLStmt(dbCmd string, dbHdl *sql.DB) (driver.Result, error) {
	var result driver.Result
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart db transaction for command", dbCmd)
		return result, err
	}
	result, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("### Failed to execute command ", dbCmd, err)
		return result, err
	}
	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction for command", dbCmd, err)
		return result, err
	}
	return result, err
}

func (obj IPV4Route) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS IPV4Routes " +
		"( DestinationNw varchar(255) PRIMARY KEY ," +
		"NetworkMask varchar(255) ," +
		"Cost integer ," +
		"NextHopIp varchar(255) ," +
		"OutgoingInterface varchar(255) ," +
		"Protocol varchar(255) )"

	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj IPV4Route) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf(`INSERT INTO IPV4Routes (DestinationNw, NetworkMask, Cost, NextHopIp, OutgoingInterface, Protocol) VALUES ('%v', '%v', %v, '%v', '%v', '%s') ;`,
		obj.DestinationNw, obj.NetworkMask, obj.Cost, obj.NextHopIp, obj.OutgoingInterface, obj.Protocol)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("### Failed to create IPV4Route ", err)
		return objectId, err
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to get ObjectID after executing ", dbCmd, err)
	}
	return objectId, err
}

func (obj IPV4Route) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	dbCmd := "delete from " + "IPV4Routes" + " where " + objKey
	fmt.Println("### DB Deleting IPV4Route")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj IPV4Route) GetKey() (string, error) {
	v4RouteKey := "DestinationNw = " + obj.DestinationNw + "NetworkMask = " + obj.NetworkMask
	return v4RouteKey, nil
}

func (obj BGPGlobalConfig) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (obj BGPGlobalConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	return int64(0), nil
}

func (obj BGPGlobalConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (obj BGPGlobalConfig) GetKey() (string, error) {
	s := ""
	return s, nil
}

func (obj BGPNeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (obj BGPNeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	return int64(0), nil
}

func (obj BGPNeighborConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (obj BGPNeighborConfig) GetKey() (string, error) {
	s := ""
	return s, nil
}

func (obj Vlan) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (obj Vlan) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	return int64(0), nil
}

func (obj Vlan) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (obj Vlan) GetKey() (string, error) {
	s := ""
	return s, nil
}

func (obj IPv4Intf) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (obj IPv4Intf) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	return int64(0), nil
}

func (obj IPv4Intf) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (obj IPv4Intf) GetKey() (string, error) {
	s := ""
	return s, nil
}

func (obj IPv4Neighbor) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (obj IPv4Neighbor) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	return int64(0), nil
}

func (obj IPv4Neighbor) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (obj IPv4Neighbor) GetKey() (string, error) {
	s := ""
	return s, nil
}

func (obj BGPGlobalState) GetKey() (string, error) {
	s := ""
	return s, nil
}

func (obj BGPNeighborState) GetKey() (string, error) {
	s := ""
	return s, nil
}
