package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/repository"
)

func GetAllOfficeUsers(officeID uint) (officeUsers []models.OfficeUser, err error) {
	officeUsers, err = repository.GetAllOfficeUsers(officeID)
	if err != nil {
		return nil, err
	}

	return officeUsers, nil
}

func GetOfficeUserById(officeUserId uint) (officeUser models.OfficeUser, err error) {
	officeUser, err = repository.GetOfficeUserByID(officeUserId)
	if err != nil {
		return officeUser, err
	}

	return officeUser, nil
}

func AddUserToOffice(officeUser models.OfficeUser) (err error) {
	if err = validators.ValidateOfficeUser(officeUser); err != nil {
		return err
	}

	err = repository.AddUserToOffice(officeUser)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserFromOffice(officeUserID uint) (err error) {
	err = repository.DeleteUserFromOffice(officeUserID)
	if err != nil {
		return err
	}

	return nil
}
