package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetCardDetailsWorkers(afterID string, month, year int) (cardDetails []models.CardDetails, err error) {
	query := db.GetDBConn().Model(&models.CardDetails{}).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("Worker.User").
		Order("id ASC").
		Limit(security.AppSettings.AppLogicParams.PaginationParams.Limit)

	if afterID != "0" {
		query = query.Where("id > ?", afterID)
	}

	err = query.Find(&cardDetails).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCardDetailsWorkers] Error getting card details: %s", err.Error())
		return cardDetails, err
	}

	return cardDetails, nil
}

func GetCardDetailsWorker(workerID uint, afterID string, month, year int) (cardDetails []models.CardDetails, err error) {
	query := db.GetDBConn().Model(&models.CardDetails{}).
		Where("worker_id = ?", workerID).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("Worker.User").
		Order("id ASC").
		Limit(security.AppSettings.AppLogicParams.PaginationParams.Limit)

	if afterID != "0" {
		query = query.Where("id > ?", afterID)
	}

	err = query.Find(&cardDetails).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCardDetailsWorker] Error getting card details: %s", err.Error())
		return cardDetails, err
	}

	return cardDetails, nil
}
