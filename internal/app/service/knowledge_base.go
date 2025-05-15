package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
)

func GetAllKnowledgeBases() (knowledgeBases []models.KnowledgeBase, err error) {
	knowledgeBases, err = repository.GetKnowledgeBasesWithDetails()
	if err != nil {
		return nil, err
	}

	return knowledgeBases, nil
}

func GetKnowledgeBaseByID(kbid uint) (knowledgeBases models.KnowledgeBase, err error) {
	knowledgeBases, err = repository.GetKnowledgeBaseByID(kbid)
	if err != nil {
		return models.KnowledgeBase{}, err
	}

	return knowledgeBases, nil
}

func CreateKnowledgeBase(kb models.KnowledgeBase) (err error) {
	if err = validators.ValidateKnowledgeBase(kb); err != nil {
		return err
	}

	_, err = repository.GetKnowledgeBaseByTitle(kb.Title)
	if err == nil {
		return errs.ErrKnowledgeBaseUniquenessFailed
	}

	err = repository.CreateKnowledgeBase(kb)
	if err != nil {
		return err
	}

	return nil
}

func UpdateKnowledgeBase(kb models.KnowledgeBase) (err error) {
	if err = validators.ValidateKnowledgeBase(kb); err != nil {
		return err
	}

	_, err = repository.GetKnowledgeBaseByTitle(kb.Title)
	if err == nil {
		return errs.ErrKnowledgeBaseUniquenessFailed
	}

	err = repository.UpdateKnowledgeBase(kb)
	if err != nil {
		return err
	}

	return nil
}

func DeleteKnowledgeBase(kbID uint) (err error) {
	err = repository.DeleteKnowledgeBase(kbID)
	if err != nil {
		return err
	}

	return nil
}
