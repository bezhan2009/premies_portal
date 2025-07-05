package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
)

func GetAllOffices(month, year uint) (offices []models.Office, err error) {
	offices, err = repository.GetAllOffices(month, year)
	if err != nil {
		return nil, err
	}

	return offices, nil
}

func GetOfficeById(month, year uint, officeID int) (office models.Office, err error) {
	office, err = repository.GetOfficeById(month, year, officeID)
	if err != nil {
		return office, err
	}

	return office, nil
}

func CreateOffice(office models.Office) (err error) {
	if err = validators.ValidateOffice(office); err != nil {
		return err
	}

	director, err := repository.GetUserByID(uint(*office.DirectorID))
	if err != nil {
		return err
	}

	if director.RoleID != 5 {
		return errs.ErrUserIsNotDirector
	}

	_, err = repository.GetOfficeByTitle(office.Title)
	if err == nil {
		return errs.ErrOfficeNameUniquenessFailed
	}

	err = repository.CreateOffice(office)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOffice(office models.Office) (err error) {
	director, err := repository.GetUserByID(uint(*office.DirectorID))
	if err != nil {
		return err
	}

	if director.RoleID != 5 {
		return errs.ErrUserIsNotDirector
	}

	_, err = repository.GetOfficeByIdOnlyOffice(int(office.ID))
	if err != nil {
		return errs.ErrOfficeNotFound
	}

	_, err = repository.GetOfficeByTitle(office.Title)
	if err == nil {
		return errs.ErrOfficeNameUniquenessFailed
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
