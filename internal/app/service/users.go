package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
)

func GetAllWorkers(afterID, month, year uint, options models.WorkerPreloadOptions) (users []models.Worker, err error) {
	users, err = repository.GetAllWorkersPag(afterID, month, year, options)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetAllUsers(afterID uint) (users []models.User, err error) {
	users, err = repository.GetAllUsersPag(afterID)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetWorkerByID(workerID, roleID, month, year uint, options models.WorkerPreloadOptions) (worker models.Worker, err error) {
	if roleID != 2 && roleID != 6 && roleID != 8 {
		return worker, errs.ErrYouAreNotWorker
	}

	worker, err = repository.GetWorkerByID(month, year, workerID, options)
	if err != nil {
		return worker, err
	}

	return worker, nil
}

func GetUserByID(userID, roleID uint) (user models.User, err error) {
	if roleID == 2 || roleID == 6 || roleID == 8 {
		return user, errs.ErrYouAreWorker
	}
	user, err = repository.GetUserByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
