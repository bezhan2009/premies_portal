package validators

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	utils2 "premiesPortal/pkg/utils"
)

func SignUpValidator(user models.User) (err error) {
	if user.Password == "" {
		return errs.ErrPasswordIsEmpty
	}

	if user.Email == "" {
		return errs.ErrEmailIsEmpty
	}

	if user.Username == "" {
		return errs.ErrUsernameIsEmpty
	}

	if user.Phone == "" || len(user.Phone) != 9 {
		return errs.ErrInvalidPhoneNumber
	}

	if user.RoleID == emptyInt {
		return errs.ErrRoleIsRequired
	}

	if user.RoleID >= 4 {
		return errs.ErrWrongRoleID
	}

	if !utils2.IsASCII(user.Username) {
		return errs.ErrUsernameCanContainOnlyASCII
	}

	return nil
}
