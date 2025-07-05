package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func CreateTestOptions(options []models.Option) (err error) {
	err = repository.CreateTestOptions(options)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTestOption(option models.Option) (err error) {
	err = repository.UpdateTestOption(option)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTestOption(optionID uint) (err error) {
	err = repository.DeleteTestOption(optionID)
	if err != nil {
		return err
	}

	return nil
}
