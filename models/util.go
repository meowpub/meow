package models

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

const gormInsertOption = "gorm:insert_option"

// genOnConflict build an ON CONFLICT clause from a model type.
func genOnConflict(t interface{}, conflictOn string) string {
	rT := reflect.TypeOf(t)

	var updates []string
	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		if field.PkgPath != "" {
			continue // Field is unexported
		}
		dbName := gorm.ToDBName(field.Name)
		if dbName == conflictOn {
			continue
		}
		updates = append(updates, dbName+"=EXCLUDED."+dbName)
	}

	return "ON CONFLICT (" + conflictOn + ") DO UPDATE SET " + strings.Join(updates, ", ")
}
