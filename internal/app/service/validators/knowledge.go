package validators

import (
	"errors"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
	"reflect"
)

func ValidateKnowledge(knowledge models.Knowledge) (err error) {
	if knowledge.KnowledgeBaseID == emptyInt {
		return errs.ErrInvalidBaseID
	}

	if knowledge.Title == "" {
		return errs.ErrInvalidTitle
	}

	return nil
}

func KnowledgeValidateData(knowledge models.Knowledge) (err error) {
	if err = ValidateKnowledge(knowledge); err != nil {
		return err
	}

	if knowledge.ID != emptyInt {
		_, err = repository.GetKnowledgeByID(knowledge.ID)
		if err != nil {
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrRecordNotFound
			}

			return err
		}
	}

	knowledgeFromDB, err := repository.GetKnowledgeByTitleAndBaseID(knowledge.Title, knowledge.KnowledgeBaseID)
	if err == nil {
		if knowledge.ID == emptyInt {
			return errs.ErrKnowledgeAlreadyExists
		}

		if reflect.DeepEqual(knowledgeFromDB.Tags, knowledge.Tags) {
			return errs.ErrKnowledgeAlreadyExists
		}
	}

	_, err = repository.GetKnowledgeBaseByID(knowledge.KnowledgeBaseID)
	if err != nil {
		return errs.ErrKnowledgeBaseNotFound
	}

	return nil
}
