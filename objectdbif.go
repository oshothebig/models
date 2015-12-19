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
	dbObjTyp := reflect.TypeOf(dbV4Route)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := []byte{0, 0, 0, 0, 0, 0, 0}
	//attrs := make(map[string]reflect.Value)
	for i:=1; i<objTyp.NumField(); i++ {
		objTyp := objTyp.Field(i)
		objVal := objVal.Field(i)
		dbObjTyp := dbObjTyp.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if objVal.Kind() == reflect.Int {
			if int(objVal.Int()) != 0 && int(objVal.Int()) != int(dbObjVal.Int()) {
				attrIds[i] = 1
				fmt.Println("\n", objTyp.Name, objTyp.Type, int(objVal.Int()))
				fmt.Println("\n", dbObjTyp.Name, dbObjTyp.Type, int(dbObjVal.Int()))
			}
		} else {
			if objVal.String() != "" && objVal.String() != dbObjVal.String() {
				attrIds[i] = 1
				fmt.Println("\n", objTyp.Name, objTyp.Type, objVal.String())
				fmt.Println("\n", dbObjTyp.Name, dbObjTyp.Type, dbObjVal.String())
			}
		}
	}

/*
	if obj.DestinationNw != "" && obj.DestinationNw != dbV4Route.DestinationNw {
		attrIds[0] = 1
	}
	if obj.NetworkMask != "" && obj.NetworkMask != dbV4Route.NetworkMask {
		attrIds[1] = 1
	}
	if obj.Cost != 0 && obj.Cost != dbV4Route.Cost {
		attrIds[2] = 1
	}
	if obj.NextHopIp != "" && obj.NextHopIp != dbV4Route.NextHopIp {
		attrIds[3] = 1
	}
	if obj.OutgoingIntfType != "" && obj.OutgoingIntfType != dbV4Route.OutgoingIntfType {
		attrIds[4] = 1
	}
	if obj.OutgoingInterface != "" && obj.OutgoingInterface != dbV4Route.OutgoingInterface {
		attrIds[5] = 1
	}
	if obj.Protocol != "" && obj.Protocol != dbV4Route.Protocol {
		attrIds[6] = 1
	}
*/
	return attrIds, nil
}

/*
func (obj IPV4Route) UpdateObjectInDb(dbHdl *sql.DB) error {
	dbCmd := "update " + "IPV4Routes" + " set" + " name" + " value" .... + " DestinationNw = " + "value" + " NetworkMask = " + "value"
}
*/

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
