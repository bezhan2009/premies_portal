package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/utils"
	"strings"
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

func GetUserByID(userID uint) (user models.User, err error) {
	user, err = repository.GetUserByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUserPassword(userID uint, oldPassword, newPassword string) error {
	newPassword = strings.TrimSpace(newPassword)
	newPassword = utils.GenerateHash(newPassword)

	oldPassword = strings.TrimSpace(oldPassword)
	oldPassword = utils.GenerateHash(oldPassword)

	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Password != oldPassword {
		return errs.ErrPermissionDenied
	}

	user.Password = newPassword
	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user models.User) (err error) {
	userDB, err := GetUserByID(user.ID)
	if err != nil {
		return err
	}

	user.Password = userDB.Password

	if user.Username == userDB.Username {
		user.Username = ""
	}

	if user.Email == userDB.Email {
		user.Email = ""
	}

	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
