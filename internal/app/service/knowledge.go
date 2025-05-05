package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func GetKnowledgeByBaseID(knowledgeBaseID uint) (knowledge []models.Knowledge, err error) {
	knowledge, err = repository.GetKnowledgeByBaseID(knowledgeBaseID)
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}

func GetKnowledgeByID(knowledgeID uint) (knowledge models.Knowledge, err error) {
	knowledge, err = repository.GetKnowledgeByID(knowledgeID)
	if err != nil {
		return models.Knowledge{}, err
	}

	return knowledge, nil

}

func CreateKnowledgeTable(knowledge models.Knowledge) (err error) {
	err = repository.CreateKnowledgeTable(knowledge)
	if err != nil {
		return err
	}

	return nil
}

func UpdateKnowledge(knowledge models.Knowledge) (err error) {
	err = repository.UpdateKnowledge(knowledge)
	if err != nil {
		return err
	}

	return nil
}

func DeleteKnowledge(knowledgeID uint) (err error) {
	err = repository.DeleteKnowledge(knowledgeID)
	if err != nil {
		return err
	}

	return nil
}
