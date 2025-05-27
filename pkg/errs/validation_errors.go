package errs

import "errors"

// Validation Errors
var (
	ErrInvalidData                 = errors.New("ErrInvalidData")
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrPathParametrized            = errors.New("ErrPathParametrized")
	ErrInvalidAmount               = errors.New("ErrInvalidAmount")
	ErrInvalidPrice                = errors.New("ErrInvalidPrice")
	ErrInvalidID                   = errors.New("ErrInvalidID")
	ErrInvalidTitle                = errors.New("ErrInvalidTitle")
	ErrInvalidToken                = errors.New("ErrInvalidToken")
	ErrRefreshTokenExpired         = errors.New("ErrRefreshTokenExpired")
	ErrFilePathIsRequired          = errors.New("ErrFilePathIsRequired")
	ErrInvalidDescription          = errors.New("ErrInvalidDescription")
	ErrUsernameCanContainOnlyASCII = errors.New("ErrUsernameCanContainOnlyASCII")
	ErrInvalidBaseID               = errors.New("ErrInvalidBaseID")
)
