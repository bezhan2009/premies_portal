package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/repository"
)

func GetAllOffices() (offices []models.Office, err error) {
	offices, err = repository.GetAllOffices()
	if err != nil {
		return nil, err
	}

	return offices, nil
}

func GetOfficeById(officeID int) (office models.OfficeAndUsers, err error) {
	office, err = repository.GetOfficesAndUsersById(officeID)
	if err != nil {
		return office, err
	}

	return office, nil
}

func CreateOffice(office models.Office) (err error) {
	if err = validators.ValidateOffice(office); err != nil {
		return err
	}

	err = repository.CreateOffice(office)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOffice(office models.Office) (err error) {
	if err = validators.ValidateOffice(office); err != nil {
		return err
	}

	err = repository.UpdateOffice(office)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOffice(officeID uint) (err error) {
	err = repository.DeleteOffice(officeID)
	if err != nil {
		return err
	}

	return nil
}
