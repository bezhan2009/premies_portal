package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func CreateTestOptions(options []models.Option) (err error) {
	if err = db.GetDBConn().Model(&models.Option{}).Create(&options).Error; err != nil {
		logger.Error.Printf("Failed to create options: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateTestOption(option models.Option) (err error) {
	if err = db.GetDBConn().Model(&models.Option{}).Where("id = ?", option.ID).Updates(&option).Error; err != nil {
		logger.Error.Printf("Failed to update options: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteTestOption(optionID uint) (err error) {
	if err = db.GetDBConn().Delete(&models.Option{}, optionID).Error; err != nil {
		logger.Error.Printf("Failed to delete option: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
