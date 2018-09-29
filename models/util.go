package models

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

const gormInsertOption = "gorm:insert_option"

// genOnConflict build an ON CONFLICT clause from a model type.
func genOnConflict(t interface{}, conflictOn string, exclude ...string) string {
	rT := reflect.TypeOf(t)

	var updates []string
	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		if field.PkgPath != "" {
			continue // Field is unexported
		}
		dbName := gorm.ToDBName(field.Name)
		if tag := field.Tag.Get("gorm"); tag != "" {
			parts := strings.Split(tag, ",")
			for _, part := range parts {
				if strings.HasPrefix(part, "column:") {
					dbName = strings.TrimPrefix(tag, "column:")
				}
			}
		}
		if dbName == "" || dbName == "-" || dbName == conflictOn {
			continue
		}
		for _, excl := range exclude {
			if dbName == excl {
				continue
			}
		}
		updates = append(updates, dbName+"=EXCLUDED."+dbName)
	}

	return "ON CONFLICT (" + conflictOn + ") DO UPDATE SET " + strings.Join(updates, ", ")
}
