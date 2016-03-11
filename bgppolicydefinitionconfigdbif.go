package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"utils/dbutils"
)

func (obj BGPPolicyDefinitionConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS BGPPolicyDefinitionConfig " +
		"( " +
		"Name TEXT, " +
		"Precedence INTEGER, " +
		"MatchType TEXT, " +
		"Export INTEGER, " +
		"Import INTEGER, " +
		"Global INTEGER, " +
		"PRIMARY KEY(Name) " +
		")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)

	dbCmd = "CREATE TABLE IF NOT EXISTS BGPPolicyDefStmtPrecedence " +
		"( " +
		"Policy TEXT, " +
		"Statement TEXT, " +
		"Precedence INTEGER, " +
		"FOREIGN KEY(Policy) REFERENCES BGPPolicyDefinitionConfig(Name) ON DELETE CASCADE, " +
		"FOREIGN KEY(Statement) REFERENCES BGPPolicyStmtConfig(Name) ON DELETE CASCADE, " +
		"PRIMARY KEY(Policy, Statement) " +
		")"

	_, err1 := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err == nil {
		err = err1
	}

	return err
}

func (obj BGPPolicyDefinitionConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO BGPPolicyDefinitionConfig (Name, Precedence, MatchType, Export, Import, Global) VALUES ('%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.Name, obj.Precedence, obj.MatchType, dbutils.ConvertBoolToInt(obj.Export), dbutils.ConvertBoolToInt(obj.Import), dbutils.ConvertBoolToInt(obj.Global))
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

	for idx := 0; idx < len(obj.StatementList); idx++ {
		dbCmd = fmt.Sprintf("INSERT INTO BGPPolicyDefStmtPrecedence (Policy, Statement, Precedence) VALUES ('%v', '%v', '%v') ;",
			obj.Name, obj.StatementList[idx].Statement, obj.StatementList[idx].Precedence)
		fmt.Println("**** Insert BGPPolicyDefStmtPrecedence called with ", obj)

		result, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
		if err != nil {
			fmt.Println("**** Failed to execute statement", dbCmd, "on BGPPolicyDefStmtPrecedence", err)
		}
	}

	return objectId, err
}

func (obj BGPPolicyDefinitionConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for BGPPolicyDefinitionConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from BGPPolicyDefinitionConfig where " + sqlKey
	fmt.Println("### DB Deleting BGPPolicyDefinitionConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj BGPPolicyDefinitionConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object BGPPolicyDefinitionConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from BGPPolicyDefinitionConfig where " + sqlKey
	var tmp3 string
	var tmp4 string
	var tmp5 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.Name, &object.Precedence, &object.MatchType, &tmp3, &tmp4, &tmp5)
	fmt.Println("### DB Get BGPPolicyDefinitionConfig\n", err)
	object.Export = dbutils.ConvertStrBoolIntToBool(tmp3)
	object.Import = dbutils.ConvertStrBoolIntToBool(tmp4)
	object.Global = dbutils.ConvertStrBoolIntToBool(tmp5)
	object.StatementList = make([]PolicyDefinitionStmtPrecedence, 0)

	if err == nil {
		dbCmd = "select * from BGPPolicyDefStmtPrecedence where POLICY=\"" + object.Name + "\""
		rows, err1 := dbHdl.Query(dbCmd)
		if err1 == nil {
			defer rows.Close()
			for rows.Next() {
				stmtPrecedenceObj := PolicyDefinitionStmtPrecedence{}
				if err = rows.Scan(&tmp3, &stmtPrecedenceObj.Statement, &stmtPrecedenceObj.Precedence); err == nil {
					object.StatementList = append(object.StatementList, stmtPrecedenceObj)
				}
			}
		} else {
			err = err1
		}
	}

	return object, err
}

func (obj BGPPolicyDefinitionConfig) GetKey() (string, error) {
	keyName := "BGPPolicyDefinitionConfig"
	keyName = strings.TrimSuffix(keyName, "Config")
	keyName = strings.TrimSuffix(keyName, "State")
	key := keyName + "#" + string(obj.Name)
	return key, nil
}

func (obj BGPPolicyDefinitionConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Name = " + "\"" + keys[1] + "\""
	return sqlKey, nil
}

func (obj BGPPolicyDefinitionConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []ConfigObj, err error) {
	var object BGPPolicyDefinitionConfig
	dbCmd := "select * from BGPPolicyDefinitionConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'BGPPolicyDefinitionConfig' with error BGPPolicyDefinitionConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()

	var tmp3 string
	var tmp4 string
	var tmp5 string
	for rows.Next() {

		if err = rows.Scan(&object.Name, &object.Precedence, &object.MatchType, &tmp3, &tmp4, &tmp5); err != nil {

			fmt.Println("Db method Scan failed when interating over BGPPolicyDefinitionConfig")
		}
		object.Export = dbutils.ConvertStrBoolIntToBool(tmp3)
		object.Import = dbutils.ConvertStrBoolIntToBool(tmp4)
		object.Global = dbutils.ConvertStrBoolIntToBool(tmp5)

		stmtPrecDBCmd := "select * from BGPPolicyDefStmtPrecedence where POLICY=\"" + object.Name + "\""
		stmtPrecRows, err := dbHdl.Query(stmtPrecDBCmd)
		if err == nil {
			for stmtPrecRows.Next() {
				stmtPrecedenceObj := PolicyDefinitionStmtPrecedence{}
				if err = stmtPrecRows.Scan(&tmp3, &stmtPrecedenceObj.Statement, &stmtPrecedenceObj.Precedence); err == nil {
					object.StatementList = append(object.StatementList, stmtPrecedenceObj)
				}
			}
			stmtPrecRows.Close()
		}

		objList = append(objList, object)
	}
	return objList, nil
}
func (obj BGPPolicyDefinitionConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(BGPPolicyDefinitionConfig)
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

func (obj BGPPolicyDefinitionConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedBGPPolicyDefinitionConfig BGPPolicyDefinitionConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedBGPPolicyDefinitionConfig)
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
	return mergedBGPPolicyDefinitionConfig, nil
}

func (obj BGPPolicyDefinitionConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbBGPPolicyDefinitionConfig := dbObj.(BGPPolicyDefinitionConfig)
	objKey, err := dbBGPPolicyDefinitionConfig.GetKey()
	objSqlKey, err := dbBGPPolicyDefinitionConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "BGPPolicyDefinitionConfig" + " set"
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
			} else if fieldVal.Kind() == reflect.Bool {
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
