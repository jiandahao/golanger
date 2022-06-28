package model

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInvalidArgument = errors.New("invalid arguments")

// IsRecordNotFound returns true if err represents a RecordNotFound error.
func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// IsInvalidArgumentError returns true if err represents a ErrInvalidArgument error.
func IsInvalidArgumentError(err error) bool {
	return errors.Is(err, ErrInvalidArgument)
}
