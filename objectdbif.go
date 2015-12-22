package models

import (
	"database/sql"
	"fmt"
	//"strconv"
	"utils/dbutils"
	"strings"
	"reflect"
)

func (obj IPV4Route) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS IPV4Routes " +
		"( DestinationNw varchar(255) PRIMARY KEY ," +
		"NetworkMask varchar(255) ," +
		"Cost integer ," +
		"NextHopIp varchar(255) ," +
		"OutgoingIntfType varchar(255) ," +
		"OutgoingInterface varchar(255) ," +
		"Protocol varchar(255) )"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj IPV4Route) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf(`INSERT INTO IPV4Routes (DestinationNw, NetworkMask, Cost, NextHopIp, OutgoingIntfType, OutgoingInterface, Protocol) VALUES ('%v', '%v', %v, '%v', '%v', '%v', '%s') ;`,
		obj.DestinationNw, obj.NetworkMask, obj.Cost, obj.NextHopIp, obj.OutgoingIntfType, obj.OutgoingInterface, obj.Protocol)
	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
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

func (obj IPV4Route) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var v4Route IPV4Route
	dbCmd := "select * from IPV4Routes where " + objKey
	err := dbHdl.QueryRow(dbCmd).Scan(&v4Route.DestinationNw, &v4Route.NetworkMask, &v4Route.Cost, &v4Route.NextHopIp, &v4Route.OutgoingIntfType, &v4Route.OutgoingInterface, &v4Route.Protocol)
	if v4Route.DestinationNw != "" && v4Route.NetworkMask != "" {
		return v4Route, err
	} else {
		return v4Route, nil
	}
}

func (obj IPV4Route) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	dbCmd := "delete from " + "IPV4Routes" + " where " + objKey
	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj IPV4Route) GetKey() (string, error) {
	v4RouteKey := obj.DestinationNw + "#" + obj.NetworkMask
	return v4RouteKey, nil
}

func (obj IPV4Route) GetSqlKeyStr(objKey string) (string, error) {
	str := strings.Split(objKey, "#")
	destNw, netMask := str[0], str[1]
	v4RouteSqlKey := "DestinationNw = " + "\"" + destNw + "\"" + " and NetworkMask = " + "\"" + netMask + "\""
	return v4RouteSqlKey, nil
}

func (obj IPV4Route) CompareObjectsAndDiff(dbObj ConfigObj) ([]byte, error) {
	dbV4Route := dbObj.(IPV4Route)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]byte, objTyp.NumField())
	for i:=0; i<objTyp.NumField(); i++ {
		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if objVal.Kind() == reflect.Int {
			if int(objVal.Int()) != 0 && int(objVal.Int()) != int(dbObjVal.Int()) {
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

func (obj IPV4Route) UpdateObjectInDb(dbObj ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbV4Route := dbObj.(IPV4Route)
	objKey, err := dbV4Route.GetKey()
	objSqlKey, err := dbV4Route.GetSqlKeyStr(objKey)
	dbCmd := "update " + "IPV4Routes" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	for i:=0; i<objTyp.NumField(); i++ {
		if attrSet[i] == 1 {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj Vlan) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var vlan Vlan
	return vlan, nil
}

func (obj IPv4Intf) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var v4Intf IPv4Intf
	return v4Intf, nil
}

func (obj IPv4Neighbor) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var v4Neighbor IPv4Neighbor
	return v4Neighbor, nil
}

func (obj BGPGlobalState) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var bgpGlobalState BGPGlobalState
	return bgpGlobalState, nil
}

func (obj BGPNeighborState) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var bgpNeighborState BGPNeighborState
	return bgpNeighborState, nil
}

func (obj ArpTimeout) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
        var arpTimeout ArpTimeout
        return arpTimeout, nil
}
