package core_errors

import "errors"

var (
	ErrNotFound      = errors.New("resource not found")
	ErrValidation    = errors.New("validation error")
	ErrConflict      = errors.New("resource conflict")
	ErrUnexpectedNil = errors.New("unexpected nil dependency")
)
