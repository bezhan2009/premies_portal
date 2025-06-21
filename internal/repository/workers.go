package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetAllWorkersPag(afterID, month, year uint) (workers []models.Worker, err error) {
	err = db.GetDBConn().
		Model(&models.Worker{}).
		Preload("CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("ServiceQuality", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("MobileBank", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("User").
		Where("id > ?", afterID).
		Order("id ASC").
		Limit(security.AppSettings.AppLogicParams.PaginationParams.Limit).
		Find(&workers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllWorkersPag] error getting all workers: %s\n", err.Error())
		return nil, TranslateGormError(err)
	}

	return workers, nil
}

func GetWorkerByID(month, year, workerID uint) (worker models.Worker, err error) {
	err = db.GetDBConn().
		Preload("CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("ServiceQuality", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("MobileBank", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("User").
		Where("id = ? OR user_id = ?", workerID, workerID).
		First(&worker).Error

	if err != nil {
		logger.Error.Printf("[repository.GetWorkerByID] error getting worker by id: %v\n", err)
		return worker, TranslateGormError(err)
	}

	return worker, nil
}
