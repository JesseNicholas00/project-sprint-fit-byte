package activity

import "errors"

var (
	ErrEmailNotFound = errors.New("userRepository: no such email found")
	ErrEmailExists   = errors.New("userRepository: email already exists")
)
