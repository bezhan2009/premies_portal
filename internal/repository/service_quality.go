package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)


func GetServiceQualitiesByID(servID uint) (serviceQualities []models.ServiceQuality, err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).
		Where("id = ?", servID).
		Find(&serviceQualities).Error; err != nil {
		logger.Error.Printf("Error getting service qualitie: %v", err)

		return nil, TranslateGormError(err)
	}

	return serviceQualities, nil
}

func GetServiceQualitiesByDateAndUserID(workerID, month, year uint) (serviceQualities []models.ServiceQuality, err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Where("worker_id = ?", workerID).
		Find(&serviceQualities).Error; err != nil {
		logger.Error.Printf("Error getting service qualities: %v", err)

		return nil, TranslateGormError(err)
	}

	return serviceQualities, nil
}

func AddServiceQuality(quality *models.ServiceQuality) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).Create(quality).Error; err != nil {
		logger.Error.Printf("[repository.AddServiceQuality] Error adding service quality %v", err)

		return 0, TranslateGormError(err)
	}

	return quality.ID, nil
}

func UpdateServiceQuality(quality models.ServiceQuality) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).Where("id = ?", quality.ID).Updates(&quality).Error; err != nil {
		logger.Error.Printf("[repository.UpdateServiceQuality] Error updating service quality %v", err)

		return 0, TranslateGormError(err)
	}

	return quality.ID, nil
}

func DeleteServiceQuality(id uint) (err error) {
	if err = db.GetDBConn().Model(&models.ServiceQuality{}).Delete(&models.ServiceQuality{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeleteServiceQuality] Error deleting service quality %v", err)

		return TranslateGormError(err)
	}

	return nil
}
