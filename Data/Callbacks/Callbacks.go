package Callbacks

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

var GlobalByName string

func SetCallbacks(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("SetCreatedBy", setCreatedBy)
	db.Callback().Update().Before("gorm:update").Register("SetUpdatedBy", setUpdatedBy)
	db.Callback().Delete().Before("gorm:delete").Register("SetDeletedBy", setDeletedBy)
}

func setCreatedBy(db *gorm.DB) {
	setByName(db, "Created")
}

func setUpdatedBy(db *gorm.DB) {
	setByName(db, "Updated")
}

func setDeletedBy(db *gorm.DB) {
	setByNameDeleted(db, "Deleted")
}

func checkByName() {
	if GlobalByName == "" {
		GlobalByName = "System"
	}
}

func setByName(db *gorm.DB, CrudType string) {
	checkByName()

	if db.Statement.Schema == nil {
		return
	}

	switch db.Statement.ReflectValue.Kind() {
	case reflect.Struct:
		db.Statement.SetColumn(CrudType+"By", GlobalByName)
	}
}

func setByNameDeleted(db *gorm.DB, CrudType string) {
	checkByName()

	tableName := db.Statement.Schema.Table
	newStatement := fmt.Sprintf("UPDATE " + tableName + " SET DeletedBy = ? WHERE Id = ?")

	dest, err := getIdFromDest(db.Statement.Dest)
	if err != nil {
		fmt.Printf("soft delete failed: " + err.Error())
		return
	}

	err = db.Exec(newStatement, GlobalByName, dest).Error
	if err != nil {
		fmt.Printf("soft delete failed: " + err.Error())
		return
	}
}

func getIdFromDest(obj interface{}) (int32, error) {
	b, _ := json.Marshal(obj)
	var subset struct {
		Id int
	}
	json.Unmarshal(b, &subset)

	return int32(subset.Id), nil
}
