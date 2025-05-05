package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func AddServiceQuality(quality models.ServiceQuality) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).Create(&quality).Error; err != nil {
		logger.Error.Printf("[repository.AddServiceQuality] Error adding service quality %v", err)

		return 0, err
	}

	return quality.ID, nil
}

func UpdateServiceQuality(quality models.ServiceQuality) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).Updates(&quality).Error; err != nil {
		logger.Error.Printf("[repository.UpdateServiceQuality] Error updating service quality %v", err)

		return 0, err
	}

	return quality.ID, nil
}

func DeleteServiceQuality(id uint) error {
	if err := db.GetDBConn().Model(&models.ServiceQuality{}).Delete(&models.ServiceQuality{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeleteServiceQuality] Error deleting service quality %v", err)

		return err
	}

	return nil
}
