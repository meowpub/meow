package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var ErrNoTTL = errors.New("a TTL is required, but not given")

type NotFoundError string

func (err NotFoundError) Error() string {
	return string(err)
}

func IsNotFound(err error) bool {
	err = errors.Cause(err)

	if _, ok := err.(NotFoundError); ok {
		return true
	} else if gorm.IsRecordNotFoundError(err) {
		return true
	} else {
		return false
	}
}
