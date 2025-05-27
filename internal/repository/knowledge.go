package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetKnowledgeByBaseID(knowledgeBaseID uint) (knowledge []models.Knowledge, err error) {
	if err = db.GetDBConn().
		Preload("KnowledgeDocs").
		Where("knowledge_base_id = ?", knowledgeBaseID).Find(&knowledge).Error; err != nil {
		logger.Error.Printf("[repository.GetKnowledgeByBaseID] Error finding knowledge by base id: %v", err)

		return nil, TranslateGormError(err)
	}

	return knowledge, nil
}

func GetKnowledgeByID(knowledgeID uint) (knowledge models.Knowledge, err error) {
	if err = db.GetDBConn().
		Preload("KnowledgeDocs").
		Where("id = ?", knowledgeID).First(&knowledge).Error; err != nil {
		logger.Error.Printf("[repository.GetKnowledgeByID] Error finding knowledge by id: %v", err)

		return models.Knowledge{}, TranslateGormError(err)
	}

	return knowledge, nil
}

func GetKnowledgeByTitleAndBaseID(title string, baseID uint) (knowledge models.Knowledge, err error) {
	if err = db.GetDBConn().Model(&models.Knowledge{}).Where("title = ? AND knowledge_base_id = ?", title, baseID).First(&knowledge).Error; err != nil {
		logger.Error.Printf("[repository.GetKnowledgeByTitle] Error finding knowledge: %v", err)

		return models.Knowledge{}, TranslateGormError(err)
	}

	return knowledge, nil
}

func CreateKnowledgeTable(knowledge models.Knowledge) (err error) {
	if err = db.GetDBConn().Model(&models.Knowledge{}).Create(&knowledge).Error; err != nil {
		logger.Error.Printf("[repository.CreateKnowledgeTable] Error while creating knowledge table: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateKnowledge(knowledge models.Knowledge) (err error) {
	if err = db.GetDBConn().Model(&models.Knowledge{}).Where("id = ?", knowledge.ID).Updates(&knowledge).Error; err != nil {
		logger.Error.Printf("[repository.UpdateKnowledge] Error while updating knowledge: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteKnowledge(knowledgeID uint) (err error) {
	knowledge, err := GetKnowledgeByID(knowledgeID)
	if err != nil {
		return TranslateGormError(err)
	}

	if err = db.GetDBConn().Model(&models.Knowledge{}).Delete(&knowledge).Error; err != nil {
		logger.Error.Printf("[repository.DeleteKnowledge] Error while deleting knowledge: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
