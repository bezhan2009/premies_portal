package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func AddServiceQuality(quality models.ServiceQuality) (id uint, err error) {
	id, err = repository.AddServiceQuality(quality)
	if err != nil {
		return id, err
	}

	return id, nil
}

func UpdateServiceQuality(quality models.ServiceQuality) (id uint, err error) {
	id, err = repository.UpdateServiceQuality(quality)
	if err != nil {
		return id, err
	}

	return id, nil
}

func DeleteServiceQuality(id uint) (err error) {
	err = repository.DeleteServiceQuality(id)
	if err != nil {
		return err
	}

	return nil
}
