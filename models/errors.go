package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var ErrNoTTL = errors.New("a TTL is required, but not given")

type notFoundError string

func (err notFoundError) Error() string {
	return string(err)
}

func IsNotFound(err error) bool {
	if _, ok := err.(notFoundError); ok {
		return true
	}
	return gorm.IsRecordNotFoundError(err)
}
