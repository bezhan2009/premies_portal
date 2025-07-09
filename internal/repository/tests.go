package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetAllTests() (tests []models.Test, err error) {
	if err = db.GetDBConn().
		Preload("Questions").
		Preload("Questions.Options").
		Preload("Answers").
		Preload("Answers.User").
		Find(&tests).Error; err != nil {
		logger.Error.Printf("Error getting all tests: %v\n", err)
		return nil, TranslateGormError(err)
	}

	return tests, nil
}

func GetTestById(id int) (test models.Test, err error) {
	if err = db.GetDBConn().
		Where("id = ?", id).
		Preload("Questions").
		Preload("Questions.Options").
		Preload("Answers").
		Preload("Answers.User").
		First(&test).Error; err != nil {
		logger.Error.Printf("Error getting test %v: %v\n", id, err)
		return models.Test{}, TranslateGormError(err)
	}

	return test, nil
}

func CreateTest(test models.Test) (err error) {
	if err = db.GetDBConn().Model(&test).Create(&test).Error; err != nil {
		logger.Error.Printf("Error creating test %v: %v\n", test, err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateTest(test models.Test) (err error) {
	if err = db.GetDBConn().Model(&test).Where("id = ?", test.ID).Updates(&test).Error; err != nil {
		logger.Error.Printf("Error updating test %v: %v\n", test, err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteTest(id int) (err error) {
	if err = db.GetDBConn().Model(&models.Test{}).Where("id = ?", id).Delete(&models.Test{}).Error; err != nil {
		logger.Error.Printf("Error deleting test %v: %v\n", id, err)

		return TranslateGormError(err)
	}

	return nil
}
