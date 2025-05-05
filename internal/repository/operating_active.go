package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func AddOperatingActive(active models.OperatingActive) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.OperatingActive{}).Create(&active).Error; err != nil {
		logger.Error.Printf("[repository.AddOperatingActive] Error while adding operating active: %v", err)

		return 0, TranslateGormError(err)
	}

	return active.ID, nil
}

func UpdateOperatingActive(active models.OperatingActive) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.OperatingActive{}).Where("id = ?", active.ID).Updates(active).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOperatingActive] Error while updating operating active: %v", err)

		return 0, TranslateGormError(err)
	}

	return active.ID, nil
}

func DeleteOperatingActive(id uint) (err error) {
	if err = db.GetDBConn().Where("id = ?", id).Delete(&models.OperatingActive{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOperatingActive] Error while deleting operating active: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
