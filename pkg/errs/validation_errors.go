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
	ErrDirectorIDIsEmpty           = errors.New("ErrDirectorIDIsEmpty")
	ErrInvalidToken                = errors.New("ErrInvalidToken")
	ErrRefreshTokenExpired         = errors.New("ErrRefreshTokenExpired")
	ErrFilePathIsRequired          = errors.New("ErrFilePathIsRequired")
	ErrOfficeIDIsEmpty             = errors.New("ErrOfficeIDIsEmpty")
	ErrUserIDIsEmpty               = errors.New("ErrUserIDIsEmpty")
	ErrInvalidDescription          = errors.New("ErrInvalidDescription")
	ErrUsernameCanContainOnlyASCII = errors.New("ErrUsernameCanContainOnlyASCII")
	ErrInvalidBaseID               = errors.New("ErrInvalidBaseID")
	ErrYouAreNotWorker             = errors.New("ErrYouAreNotWorker")
	ErrYouAreWorker                = errors.New("ErrYouAreWorker")
	ErrInvalidAfterID              = errors.New("ErrInvalidAfterID")
	ErrInvalidMonth                = errors.New("ErrInvalidMonth")
	ErrInvalidYear                 = errors.New("ErrInvalidYear")
)
