package models

import (
	"github.com/jinzhu/gorm"
)

func IsNotFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
