package repository

import (
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetTestAnswers(month, year uint) (answers []models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).
		Preload("User").
		Preload("Question").
		Preload("Question.Options").
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Find(&answers).Error; err != nil {
		logger.Error.Printf("Failed to get answers: %v", err)

		return nil, TranslateGormError(err)
	}

	return answers, nil
}

func GetTestAnswersByUserID(userID, month, year uint) (answer models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).
		Where("user_id = ?", userID).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		First(&answer).Error; err != nil {
		logger.Error.Printf("Failed to get answers: %v", err)

		return models.Answer{}, TranslateGormError(err)
	}

	return answer, nil
}

func GetTestAnswersByTestId(testID int) (answer []models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).
		Preload("User").
		Preload("Question").
		Preload("Question.Options").
		Where("test_id = ?", testID).
		Find(&answer).Error; err != nil {
		logger.Error.Printf("Failed to get answer: %v", err)

		return answer, TranslateGormError(err)
	}

	return answer, nil
}

func GetTestAnswersByAnswerId(answerID int) (answer models.Answer, err error) {
	if err = db.GetDBConn().Model(&models.Answer{}).Where("id = ?", answerID).
		Preload("User").
		Preload("Question").
		Preload("Question.Options").
		First(&answer).Error; err != nil {
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

func SaveScoresAndAnswers(
	tx *gorm.DB,
	serviceQualities []models.ServiceQuality,
	score float64,
	answers []models.Answer,
) error {
	for _, sq := range serviceQualities {
		sq.Tests = score
		if err := tx.Save(&sq).Error; err != nil {
			return err
		}
	}

	// Сохраняем ответы и их опции по отдельности
	for i := range answers {
		ans := &answers[i]
		// Обнуляем ID, чтобы GORM использовал автоинкремент
		ans.ID = 0
		if err := tx.Create(ans).Error; err != nil {
			return err
		}
		// Сохраняем связанные выбранные опции
		for j := range ans.SelectedOptions {
			opt := &ans.SelectedOptions[j]
			opt.ID = 0
			opt.AnswerID = ans.ID
			if err := tx.Create(opt).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
