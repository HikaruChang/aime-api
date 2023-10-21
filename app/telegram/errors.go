package telegram

import "errors"

var (
	ErrAuthDateMissing  = errors.New("auth_date is missing")
	ErrSignMissing      = errors.New("sign is missing")
	ErrSignInvalid      = errors.New("sign is invalid")
	ErrUnexpectedFormat = errors.New("data has unexpected format")
	ErrExpired          = errors.New("data is expired")
)
