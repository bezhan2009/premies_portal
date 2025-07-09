package service

import (
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
)

func CreateWorker(tx *gorm.DB, officeID uint, worker models.Worker) (err error) {
	if officeID < 1 {
		return errs.ErrOfficeIDIsEmpty
	}

	workerID, err := repository.CreateWorker(tx, worker)
	if err != nil {
		return err
	}

	err = AddUserToOfficeByTX(
		tx,
		&models.OfficeUser{
			OfficeID: int(officeID),
			WorkerID: int(workerID),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
