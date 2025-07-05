package validators

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
)

func ValidateKnowledgeBase(kb models.KnowledgeBase) (err error) {
	if kb.Title == "" {
		return errs.ErrInvalidTitle
	}

	return nil
}
