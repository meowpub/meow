package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var ErrNoTTL = errors.New("a TTL is required, but not given")

type NotFoundError string

func (err NotFoundError) Error() string {
	return string(err)
}

func IsNotFound(err error) bool {
	if _, ok := err.(NotFoundError); ok {
		return true
	}
	return gorm.IsRecordNotFoundError(err)
}
