package validators

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
)

func ValidateOfficeUser(officeUser models.OfficeUser) (err error) {
	if officeUser.OfficeID == emptyInt {
		return errs.ErrOfficeIDIsEmpty
	}

	if officeUser.WorkerID == emptyInt {
		return errs.ErrUserIDIsEmpty
	}

	return nil
}
