package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetKnowledgeDocByID(knowledgeDocsID uint) (knowledgeDocs models.KnowledgeDocs, err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeDocs{}).Where("id = ?", knowledgeDocsID).First(&knowledgeDocs).Error; err != nil {
		logger.Error.Printf("[repository.GetKnowledgeDocByID] Error while getting knowledge docs by id: %v\n", err)

		return knowledgeDocs, err
	}

	return knowledgeDocs, err
}

func GetKnowledgeDocsByKnowledgeID(knowledgeID uint) (knowledgeDocs models.KnowledgeDocs, err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeDocs{}).Where("knowledge_id = ?", knowledgeID).First(&knowledgeDocs).Error; err != nil {
		logger.Error.Printf("[repository.GetKnowledgeDocsByKnowledgeID] Error while getting knowledge docs: %v\n", err)

		return knowledgeDocs, TranslateGormError(err)
	}

	return knowledgeDocs, nil
}

func CreateKnowledgeDocs(knowledgeDocs models.KnowledgeDocs) (err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeDocs{}).Create(&knowledgeDocs).Error; err != nil {
		logger.Error.Printf("[repository.CreateKnowledgeDocs] Error while creating knowledge doc: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateKnowledgeDocs(knowledgeDocs models.KnowledgeDocs) (err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeDocs{}).Updates(knowledgeDocs).Error; err != nil {
		logger.Error.Printf("[repository.UpdateKnowledgeDocs] Error while updating knowledge doc: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteKnowledgeDocs(knowledgeDocID uint) (err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeDocs{}).Where("id = ?", knowledgeDocID).Delete(&models.KnowledgeDocs{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteKnowledgeDocs] Error while deleting knowledge doc: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}
