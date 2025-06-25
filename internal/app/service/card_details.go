package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func GetCardDetailsWorkers(afterID string, month, year int) (cardDetails []models.CardDetails, err error) {
	if afterID == "" {
		afterID = "0"
	}

	cardDetails, err = repository.GetCardDetailsWorkers(afterID, month, year)
	if err != nil {
		return nil, err
	}

	return cardDetails, nil
}

func GetCardDetailsWorker(workerID uint, afterID string, month, year int) (cardDetails []models.CardDetails, err error) {
	if afterID == "" {
		afterID = "0"
	}

	worker, err := repository.GetWorkerByUserID(workerID)
	if err != nil {
		return nil, err
	}

	cardDetails, err = repository.GetCardDetailsWorker(worker.ID, afterID, month, year)
	if err != nil {
		return nil, err
	}

	return cardDetails, nil
}
