package model

import (
	"errors"
)

// Model errors
var (
	ErrInvalidArgs  = errors.New("Invalid Args")
	ErrKeyConflict  = errors.New("Key Conflict")
	ErrDataNotFound = errors.New("Record Not Found")
	ErrUserExists   = errors.New("User already exists")
	ErrUnknown      = errors.New("Unknown Error")
	ErrFailed       = errors.New("Failed")
)
