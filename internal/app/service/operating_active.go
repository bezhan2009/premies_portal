package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func AddOperatingActive(active models.OperatingActive) (id uint, err error) {
	id, err = repository.AddOperatingActive(active)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateOperatingActive(active models.OperatingActive) (id uint, err error) {
	id, err = repository.UpdateOperatingActive(active)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteOperatingActive(id uint) (err error) {
	err = repository.DeleteOperatingActive(id)
	if err != nil {
		return err
	}

	return nil
}
