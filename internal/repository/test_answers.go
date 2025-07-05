package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetTestAnswers() (answers []models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).
		Preload("User").
		Find(&answers).Error; err != nil {
		logger.Error.Printf("Failed to get answers: %v", err)

		return nil, TranslateGormError(err)
	}

	return answers, nil
}

func GetTestAnswersByTestId(testID int) (answer []models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).
		Preload("User").
		Preload("Question").
		Where("test_id = ?", testID).
		Find(&answer).Error; err != nil {
		logger.Error.Printf("Failed to get answer: %v", err)

		return answer, TranslateGormError(err)
	}

	return answer, nil
}

func GetTestAnswersByAnswerId(answerID int) (answer []models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).Where("answer_id = ?", answerID).Find(&answer).Error; err != nil {
		logger.Error.Printf("Failed to get answer: %v", err)

		return answer, TranslateGormError(err)
	}

	return answer, nil
}

func CreateTestAnswers(answer []models.Answer) (err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).Create(&answer).Error; err != nil {
		logger.Error.Printf("Failed to create answer: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
