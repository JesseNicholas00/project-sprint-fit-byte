package user

import "errors"

var (
	ErrEmailAlreadyRegistered = errors.New("userService: email already registered")
	ErrUserNotFound           = errors.New("userService: no such user found")
)
