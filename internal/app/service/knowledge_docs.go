package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func GetKnowledgeDocByID(knowledgeDocsID uint) (knowledgeDocs models.KnowledgeDocs, err error) {
	knowledgeDocs, err = repository.GetKnowledgeDocByID(knowledgeDocsID)
	if err != nil {
		return models.KnowledgeDocs{}, err
	}

	return knowledgeDocs, nil
}

func GetKnowledgeDocsByKnowledgeID(knowledgeID uint) (knowledgeDocs models.KnowledgeDocs, err error) {
	knowledgeDocs, err = repository.GetKnowledgeDocsByKnowledgeID(knowledgeID)
	if err != nil {
		return models.KnowledgeDocs{}, err
	}

	return knowledgeDocs, nil
}

func CreateKnowledgeDocs(knowledgeDocs models.KnowledgeDocs) (err error) {
	err = repository.CreateKnowledgeDocs(knowledgeDocs)
	if err != nil {
		return err
	}

	return nil
}

func UpdateKnowledgeDocs(knowledgeDocs models.KnowledgeDocs) (err error) {
	err = repository.UpdateKnowledgeDocs(knowledgeDocs)
	if err != nil {
		return err
	}

	return nil
}

func DeleteKnowledgeDocs(knowledgeDocID uint) (err error) {
	err = repository.DeleteKnowledgeDocs(knowledgeDocID)
	if err != nil {
		return err
	}

	return nil
}
