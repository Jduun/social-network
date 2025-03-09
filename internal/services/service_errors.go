package services

import "errors"

var (
	ErrUserNotFound       = errors.New("User not found")
	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrPasswordHashing    = errors.New("Password hashing error")
	ErrCannotCreateUser   = errors.New("Cannot create user")
	ErrCannotFindUser     = errors.New("Cannot find user")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrCannotSignToken    = errors.New("Cannot sign token")
	ErrInvalidToken       = errors.New("Invalid token")
	ErrCannotLoginUser    = errors.New("Cannot login user")
)
