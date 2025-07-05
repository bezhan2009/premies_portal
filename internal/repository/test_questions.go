package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func CreateTestQuestions(questions []models.Question) (err error) {
	if err = db.GetDBConn().Model(&models.Question{}).Create(&questions).Error; err != nil {
		logger.Error.Printf("Failed to create test question: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateTestQuestion(question models.Question) (err error) {
	if err = db.GetDBConn().Model(&models.Question{}).Where("id = ?", question.ID).Updates(&question).Error; err != nil {
		logger.Error.Printf("Failed to update test question: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteTestQuestion(questionID uint) (err error) {
	if err = db.GetDBConn().Delete(&models.Question{}, questionID).Error; err != nil {
		logger.Error.Printf("Failed to delete test question: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
