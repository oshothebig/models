package models

import (
	"database/sql"
	"fmt"
	"strings"
	"utils/dbutils"
	"reflect"
)

func (obj PolicyStmtConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS PolicyStmtConfig " +
		"( " +
		"Name TEXT, " +
		"MatchConditions TEXT, " +
		"Conditions TEXT, " +
		"Actions TEXT, " +
		"PRIMARY KEY(Name) " +
	")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj PolicyStmtConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO PolicyStmtConfig (Name, MatchConditions, Conditions, Actions) VALUES ('%v', '%v', '%v', '%v') ;",
		obj.Name, obj.MatchConditions, obj.Conditions, obj.Actions)
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

func (obj PolicyStmtConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for PolicyStmtConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from PolicyStmtConfig where " + sqlKey
	fmt.Println("### DB Deleting PolicyStmtConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj PolicyStmtConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object PolicyStmtConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from PolicyStmtConfig where " + sqlKey
	var tmp2 string
	var tmp3 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.Name, &object.MatchConditions, &tmp2, &tmp3, )
	fmt.Println("### DB Get PolicyStmtConfig\n", err)
convtmpConditions := strings.Split(tmp2, ",")
                        for _, x := range convtmpConditions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Conditions = append(object.Conditions, string(y))
                     }
convtmpActions := strings.Split(tmp3, ",")
                        for _, x := range convtmpActions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Actions = append(object.Actions, string(y))
                     }
	return object, err
}

func (obj PolicyStmtConfig) GetKey() (string, error) {
	key := string(obj.Name)
	return key, nil
}

func (obj PolicyStmtConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Name = "+ "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *PolicyStmtConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []*PolicyStmtConfig, e error) {
	dbCmd := "select * from PolicyStmtConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'PolicyStmtConfig' with error PolicyStmtConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()
    
	var tmp2 string
	var tmp3 string
	for rows.Next() {

             object := new(PolicyStmtConfig)
             if err = rows.Scan(&object.Name, &object.MatchConditions, &object.Conditions, &object.Actions, ); err != nil {

             fmt.Println("Db method Scan failed when interating over PolicyStmtConfig")
             }
convtmpConditions := strings.Split(tmp2, ",")
                        for _, x := range convtmpConditions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Conditions = append(object.Conditions, string(y))
                     }
convtmpActions := strings.Split(tmp3, ",")
                        for _, x := range convtmpActions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Actions = append(object.Actions, string(y))
                     }
	objList = append(objList, object)
    }
    return objList, nil
    }
    func (obj PolicyStmtConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(PolicyStmtConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]bool, objTyp.NumField())
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
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

    func (obj PolicyStmtConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedPolicyStmtConfig PolicyStmtConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedPolicyStmtConfig)
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if  attrSet[idx] {
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
	return mergedPolicyStmtConfig, nil
}

    func (obj PolicyStmtConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbPolicyStmtConfig := dbObj.(PolicyStmtConfig)
	objKey, err := dbPolicyStmtConfig.GetKey()
	objSqlKey, err := dbPolicyStmtConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "PolicyStmtConfig" + " set"
objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
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
