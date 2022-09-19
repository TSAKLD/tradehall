package service

import "errors"

var (
	ErrNicknameTaken = errors.New("nickname already taken")
	ErrEmailTaken    = errors.New("email already taken")

	ErrNotFound   = errors.New("entity not found")
	ErrNotAllowed = errors.New("not enough permissions")
)
