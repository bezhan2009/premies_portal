package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/repository"
)

func GetAllKnowledgeBases() (knowledgeBases []models.KnowledgeBase, err error) {
	knowledgeBases, err = repository.GetKnowledgeBasesWithDetails()
	if err != nil {
		return nil, err
	}

	return knowledgeBases, nil
}

func CreateKnowledgeBase(kb models.KnowledgeBase) (err error) {
	if err = validators.ValidateKnowledgeBase(kb); err != nil {
		return err
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
