package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetKnowledgeBasesWithDetails() (kb []models.KnowledgeBase, err error) {
	err = db.GetDBConn().
		Preload("Knowledge").               // загружаем все записи Knowledge
		Preload("Knowledge.KnowledgeDocs"). // загружаем связанные документы для каждой Knowledge
		Find(&kb).Error                     // находим по ID

	if err != nil {
		logger.Error.Printf("[repository.GetKnowledgeBaseWithDetails] error getting knowledge base with details: %v\n", err)
		return kb, TranslateGormError(err)
	}

	return kb, nil
}

func GetKnowledgeBaseByID(kbID uint) (kb *models.KnowledgeBase, err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeBase{}).Where("id = ?", kbID).First(&kb).Error; err != nil {
		logger.Error.Println("[repository.GetKnowledgeBaseByID] Error getting knowledge base with ID: %v\n]")

		return kb, TranslateGormError(err)
	}

	return kb, nil
}

func CreateKnowledgeBase(kb models.KnowledgeBase) (err error) {
	if err = db.GetDBConn().Model(&models.KnowledgeBase{}).Create(&kb).Error; err != nil {
		logger.Error.Printf("[repository.CreateKnowledgeBase] error creating knowledge base: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateKnowledgeBase(kb models.KnowledgeBase) (err error) {
	if err = db.GetDBConn().Model(&models.Knowledge{}).Save(&kb).Error; err != nil {
		logger.Error.Printf("[repository.UpdateKnowledgeBase] error updating knowledge base: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteKnowledgeBase(kbID uint) (err error) {
	kb, err := GetKnowledgeBaseByID(kbID)
	if err != nil {
		return TranslateGormError(err)
	}

	if err = db.GetDBConn().Model(&models.Knowledge{}).Delete(&kb).Error; err != nil {
		logger.Error.Printf("[repository.DeleteKnowledgeBase] error deleting knowledge base: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}
