package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"strconv"
	"time"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(int(time.Now().Month()), strconv.Itoa(int(id)))
	if err != nil {
		return user, err
	}

	return user, nil
}
