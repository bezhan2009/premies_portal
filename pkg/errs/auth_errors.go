package errs

import "errors"

// Authentication Errors
var (
	ErrInvalidCredentials          = errors.New("invalid credentials")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrYouHaveNotRegisteredYet     = errors.New("ErrYouHaveNotRegisteredYet")
	ErrPasswordIsEmpty             = errors.New("ErrPasswordIsEmpty")
	ErrUsernameIsEmpty             = errors.New("ErrUsernameIsEmpty")
	ErrInvalidPhoneNumber          = errors.New("ErrInvalidPhoneNumber")
	ErrEmailIsEmpty                = errors.New("ErrEmailIsEmpty")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUnauthorized                = errors.New("ErrUnauthorized")
	ErrEmailIsRequired             = errors.New("email is required")
	ErrUsernameIsRequired          = errors.New("username is required")
	ErrFirstNameIsRequired         = errors.New("first name is required")
	ErrLastNameIsRequired          = errors.New("last name is required")
	ErrPasswordIsRequired          = errors.New("password is required")
	ErrRoleIsRequired              = errors.New("role is required")
	ErrWrongRoleID                 = errors.New("wrong role")
)
