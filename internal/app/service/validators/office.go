package validators

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
)

func ValidateOffice(office models.Office) (err error) {
	if office.Title == "" {
		return errs.ErrInvalidTitle
	}

	if office.Description == "" {
		return errs.ErrInvalidDescription
	}

	if office.DirectorID == nil {
		return errs.ErrDirectorIDIsEmpty
	}

	return nil
}
