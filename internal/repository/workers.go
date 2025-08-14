package repository

import (
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func CreateWorker(tx *gorm.DB, worker models.Worker) (workerID uint, err error) {
	if err = tx.Create(&worker).Error; err != nil {
		logger.Error.Printf("[repository.CreateWorker] Failed to create worker: %v", err)

		return workerID, TranslateGormError(err)
	}

	return worker.ID, nil
}

func GetAllWorkersPag(afterID, month, year uint, opts models.WorkerPreloadOptions) (workers []models.Worker, err error) {
	query := db.GetDBConn().Model(&models.Worker{})

	if opts.LoadCardTurnovers {
		query = query.Preload("CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadCardSales {
		query = query.Preload("CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadServiceQuality {
		query = query.Preload("ServiceQuality", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadMobileBank {
		query = query.Preload("MobileBank", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadCardDetails {
		query = query.Preload("CardDetails", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadUser {
		query = query.Preload("User")
	}

	err = query.
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

func GetWorkerByID(month, year, workerID uint, opts models.WorkerPreloadOptions) (worker models.Worker, err error) {
	query := db.GetDBConn().Model(&models.Worker{})

	if opts.LoadCardTurnovers {
		query = query.Preload("CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadCardSales {
		query = query.Preload("CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadServiceQuality {
		query = query.Preload("ServiceQuality", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadMobileBank {
		query = query.Preload("MobileBank", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadCardDetails {
		query = query.Preload("CardDetails", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year)
	}
	if opts.LoadUser {
		query = query.Preload("User")
	}

	err = query.Where("id = ?", workerID).First(&worker).Error
	if err != nil {
		logger.Error.Printf("[repository.GetWorkerByID] error getting worker by id: %v\n", err)
		return worker, TranslateGormError(err)
	}

	return worker, nil
}

func GetWorkerByUserID(workerID uint) (worker models.Worker, err error) {
	err = db.GetDBConn().Model(&models.Worker{}).Where("user_id = ?", workerID).First(&worker).Error
	if err != nil {
		logger.Error.Printf("[repository.GetWorkerByID] error getting worker by id: %v\n", err)

		return worker, TranslateGormError(err)
	}

	return worker, nil
}

func UpdateWorkerByID(worker models.Worker) (err error) {
	err = db.GetDBConn().Model(&models.Worker{}).Where("id = ?", worker.ID).Updates(&worker).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateWorkerByID] error updating worker by id: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteWorkerByID(workerID uint) (err error) {
	err = db.GetDBConn().Model(&models.Worker{}).Where("id = ?", workerID).Delete(&models.Worker{}).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteWorkerByID] error deleting worker by id: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}
