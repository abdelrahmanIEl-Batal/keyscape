package store

import "errors"

var (
	ErrInvalidCommand = errors.New("invalid_command")
	ErrKeyNotFound    = errors.New("key_not_found")
)
