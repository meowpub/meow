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
	seenExcludes := make(map[string]struct{})

	var updates []string
fieldLoop:
	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		if field.PkgPath != "" {
			continue // Field is unexported
		}
		dbName := gorm.ToDBName(field.Name)
		if tag := field.Tag.Get("gorm"); tag != "" {
			if tag == "-" {
				continue
			}
			parts := strings.Split(tag, ",")
			for _, part := range parts {
				if strings.HasPrefix(part, "column:") {
					dbName = strings.TrimPrefix(tag, "column:")
				}
			}
		}
		if dbName == "" || dbName == conflictOn {
			continue
		}
		for _, excl := range exclude {
			if dbName == excl {
				seenExcludes[dbName] = struct{}{}
				continue fieldLoop
			}
		}
		updates = append(updates, dbName+"=EXCLUDED."+dbName)
	}

	for _, excl := range exclude {
		if _, ok := seenExcludes[excl]; !ok {
			panic("genOnConflict exclusion on nonexistent column: " + excl)
		}
	}

	return "ON CONFLICT (" + conflictOn + ") DO UPDATE SET " + strings.Join(updates, ", ")
}
