package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetAllWorkersMobileBankDetails(afterID, workerID, month, year uint) (mobileBankDetails []models.MobileBankDetails, err error) {
	if err = db.GetDBConn().Model(&models.MobileBankDetails{}).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Where("worker_id = ?", workerID).
		Where("id > ?", afterID).
		Order("id ASC").
		Limit(security.AppSettings.AppLogicParams.PaginationParams.Limit).
		Find(&mobileBankDetails).Error; err != nil {
		logger.Error.Printf("[repository.GetAllWorkersMobileBankDetails] Error while getting mobile bank details: %v", err)

		return nil, err
	}

	return mobileBankDetails, nil
}
