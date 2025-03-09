package repositories

import "errors"

var (
	ErrNotFound      = errors.New("Not found")
	ErrAlreadyExists = errors.New("Already exists")
)
