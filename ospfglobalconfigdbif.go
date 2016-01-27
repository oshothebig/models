package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj OspfGlobalConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS OspfGlobalConfig " +
		"( " +
		"RouterIdKey TEXT, " +
		"AdminStat INTEGER, " +
		"ASBdrRtrStatus INTEGER, " +
		"TOSSupport INTEGER, " +
		"ExtLsdbLimit INTEGER, " +
		"MulticastExtensions INTEGER, " +
		"ExitOverflowInterval INTEGER, " +
		"DemandExtensions INTEGER, " +
		"RFC1583Compatibility INTEGER, " +
		"ReferenceBandwidth INTEGER, " +
		"RestartSupport INTEGER, " +
		"RestartInterval INTEGER, " +
		"RestartStrictLsaChecking INTEGER, " +
		"StubRouterAdvertisement INTEGER, " +
		"PRIMARY KEY(RouterIdKey) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj OspfGlobalConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO OspfGlobalConfig (RouterIdKey, AdminStat, ASBdrRtrStatus, TOSSupport, ExtLsdbLimit, MulticastExtensions, ExitOverflowInterval, DemandExtensions, RFC1583Compatibility, ReferenceBandwidth, RestartSupport, RestartInterval, RestartStrictLsaChecking, StubRouterAdvertisement) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.RouterIdKey, obj.AdminStat, dbutils.ConvertBoolToInt(obj.ASBdrRtrStatus), dbutils.ConvertBoolToInt(obj.TOSSupport), obj.ExtLsdbLimit, obj.MulticastExtensions, obj.ExitOverflowInterval, dbutils.ConvertBoolToInt(obj.DemandExtensions), dbutils.ConvertBoolToInt(obj.RFC1583Compatibility), obj.ReferenceBandwidth, obj.RestartSupport, obj.RestartInterval, dbutils.ConvertBoolToInt(obj.RestartStrictLsaChecking), obj.StubRouterAdvertisement)
	fmt.Println("**** Create Object called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	} else {
		objectId, err = result.LastInsertId()
		if err != nil {
			fmt.Println("### Failed to return last object id", err)
		}

	}
	return objectId, err
}

func (obj OspfGlobalConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for OspfGlobalConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from OspfGlobalConfig where " + sqlKey
	fmt.Println("### DB Deleting OspfGlobalConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj OspfGlobalConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object OspfGlobalConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from OspfGlobalConfig where " + sqlKey
	var tmp2 string
	var tmp3 string
	var tmp7 string
	var tmp8 string
	var tmp12 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.RouterIdKey, &object.AdminStat, &tmp2, &tmp3, &object.ExtLsdbLimit, &object.MulticastExtensions, &object.ExitOverflowInterval, &tmp7, &tmp8, &object.ReferenceBandwidth, &object.RestartSupport, &object.RestartInterval, &tmp12, &object.StubRouterAdvertisement)
	fmt.Println("### DB Get OspfGlobalConfig\n", err)
	object.ASBdrRtrStatus = dbutils.ConvertStrBoolIntToBool(tmp2)
	object.TOSSupport = dbutils.ConvertStrBoolIntToBool(tmp3)
	object.DemandExtensions = dbutils.ConvertStrBoolIntToBool(tmp7)
	object.RFC1583Compatibility = dbutils.ConvertStrBoolIntToBool(tmp8)
	object.RestartStrictLsaChecking = dbutils.ConvertStrBoolIntToBool(tmp12)
	return object, err
}

func (obj OspfGlobalConfig) GetKey() (string, error) {
	key := string(obj.RouterIdKey)
	return key, nil
}

func (obj OspfGlobalConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "RouterIdKey = " + "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *OspfGlobalConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []*OspfGlobalConfig, e error) {
	dbCmd := "select * from OspfGlobalConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'OspfGlobalConfig' with error OspfGlobalConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()

	var tmp2 string
	var tmp3 string
	var tmp7 string
	var tmp8 string
	var tmp12 string
	for rows.Next() {

		object := new(OspfGlobalConfig)
		if err = rows.Scan(&object.RouterIdKey, &object.AdminStat, &tmp2, &tmp3, &object.ExtLsdbLimit, &object.MulticastExtensions, &object.ExitOverflowInterval, &tmp7, &tmp8, &object.ReferenceBandwidth, &object.RestartSupport, &object.RestartInterval, &tmp12, &object.StubRouterAdvertisement); err != nil {

			fmt.Println("Db method Scan failed when interating over OspfGlobalConfig")
		}
		object.ASBdrRtrStatus = dbutils.ConvertStrBoolIntToBool(tmp2)
		object.TOSSupport = dbutils.ConvertStrBoolIntToBool(tmp3)
		object.DemandExtensions = dbutils.ConvertStrBoolIntToBool(tmp7)
		object.RFC1583Compatibility = dbutils.ConvertStrBoolIntToBool(tmp8)
		object.RestartStrictLsaChecking = dbutils.ConvertStrBoolIntToBool(tmp12)
		objList = append(objList, object)
	}
	return objList, nil
}
func (obj OspfGlobalConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(OspfGlobalConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]bool, objTyp.NumField())
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		fieldTyp := objTyp.Field(i)
		if fieldTyp.Anonymous {
			continue
		}

		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if _, ok := updateKeys[fieldTyp.Name]; ok {
			if objVal.Kind() == reflect.Int {
				if int(objVal.Int()) != int(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int8 {
				if int8(objVal.Int()) != int8(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int16 {
				if int16(objVal.Int()) != int16(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int32 {
				if int32(objVal.Int()) != int32(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Int64 {
				if int64(objVal.Int()) != int64(dbObjVal.Int()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint {
				if uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint8 {
				if uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint16 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint32 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Uint64 {
				if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
					attrIds[idx] = true
				}
			} else if objVal.Kind() == reflect.Bool {
				if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
					attrIds[idx] = true
				}
			} else {
				if objVal.String() != dbObjVal.String() {
					attrIds[idx] = true
				}
			}
			if attrIds[idx] {
				fmt.Println("attribute changed ", fieldTyp.Name)
			}
		}
		idx++

	}
	return attrIds[:idx], nil
}

func (obj OspfGlobalConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedOspfGlobalConfig OspfGlobalConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedOspfGlobalConfig)
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if attrSet[idx] {
			if dbObjField.Kind() == reflect.Int ||
				dbObjField.Kind() == reflect.Int8 ||
				dbObjField.Kind() == reflect.Int16 ||
				dbObjField.Kind() == reflect.Int32 ||
				dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(objField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint8 ||
				dbObjField.Kind() == reflect.Uint16 ||
				dbObjField.Kind() == reflect.Uint32 ||
				dbObjField.Kind() == reflect.Uint64 {
				mergedObjVal.Elem().Field(i).SetUint(objField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(objField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(objField.String())
			}
		} else {
			if dbObjField.Kind() == reflect.Int ||
				dbObjField.Kind() == reflect.Int8 ||
				dbObjField.Kind() == reflect.Int16 ||
				dbObjField.Kind() == reflect.Int32 ||
				dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint ||
				dbObjField.Kind() == reflect.Uint8 ||
				dbObjField.Kind() == reflect.Uint16 ||
				dbObjField.Kind() == reflect.Uint32 {
				mergedObjVal.Elem().Field(i).SetUint(dbObjField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
				mergedObjVal.Elem().Field(i).SetBool(dbObjField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
			}
		}
		idx++

	}
	return mergedOspfGlobalConfig, nil
}

func (obj OspfGlobalConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbOspfGlobalConfig := dbObj.(OspfGlobalConfig)
	objKey, err := dbOspfGlobalConfig.GetKey()
	objSqlKey, err := dbOspfGlobalConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "OspfGlobalConfig" + " set"
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	idx := 0
	for i := 0; i < objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		if attrSet[idx] {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int ||
				fieldVal.Kind() == reflect.Int8 ||
				fieldVal.Kind() == reflect.Int16 ||
				fieldVal.Kind() == reflect.Int32 ||
				fieldVal.Kind() == reflect.Int64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
			} else if fieldVal.Kind() == reflect.Uint ||
				fieldVal.Kind() == reflect.Uint8 ||
				fieldVal.Kind() == reflect.Uint16 ||
				fieldVal.Kind() == reflect.Uint32 ||
				fieldVal.Kind() == reflect.Uint64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Uint()))
			} else if objVal.Kind() == reflect.Bool {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, dbutils.ConvertBoolToInt(bool(fieldVal.Bool())))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
		idx++
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
