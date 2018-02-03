package model

import (
	"errors"
)

// Model errors
var (
	ErrInvalidArgs  = errors.New("Invalid Args")
	ErrKeyConflict  = errors.New("Key Conflict")
	ErrDataNotFound = errors.New("Record Not Found")
	ErrUnknown      = errors.New("Unknown Error")
	ErrFailed       = errors.New("Failed")
)
