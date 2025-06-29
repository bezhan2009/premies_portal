package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func GetAllWorkersMobileBankDetails(afterID, userID, month, year uint) (mobileBankDetails []models.MobileBankDetails, err error) {
	worker, err := repository.GetWorkerByUserID(userID)
	if err != nil {
		return nil, err
	}

	mobileBankDetails, err = repository.GetAllWorkersMobileBankDetails(afterID, worker.ID, month, year)
	if err != nil {
		return nil, err
	}

	return mobileBankDetails, nil
}
