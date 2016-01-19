package models

import (
	"database/sql"
)

type BaseObj struct{}

func (b BaseObj) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (b BaseObj) StoreObjectInDb(dbHdl *sql.DB) (objId int64, err error) {
	return objId, err
}

func (b BaseObj) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (b BaseObj) GetKey() (string, error) {
	s := ""
	return s, nil
}
func (b BaseObj) GetSqlKeyStr(objKey string) (string, error) {
	s := ""
	return s, nil
}

func (b BaseObj) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	return nil, nil
}

func (b BaseObj) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	var arr []bool
	return arr, nil
}

func (b BaseObj) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	return nil, nil
}

func (b BaseObj) UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	return nil
}

