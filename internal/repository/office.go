package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func GetAllOffices(month, year uint) (offices []models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).
		Preload("OfficeUsers").
		Preload("OfficeUsers.Worker").
		Preload("OfficeUsers.Worker.CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("OfficeUsers.Worker.CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("OfficeUsers.Worker.User").
		Find(&offices).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetAllOffices] Error while getting all offices: %v", err)

		return nil, TranslateGormError(err)
	}

	return offices, nil
}

func GetOfficeByDirectorID(directorID, month, year uint) (office models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Where("director_id = ?", directorID).
		Preload("OfficeUsers").
		Preload("OfficeUsers.Worker").
		Preload("OfficeUsers.Worker.CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("OfficeUsers.Worker.CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("OfficeUsers.Worker.User").First(&office).Error; err != nil {
		logger.Error.Printf("Error while getting office by worker id: %v", err)

		return models.Office{}, TranslateGormError(err)
	}

	return office, nil
}

func GetOfficesAndUsersById(officeID int) (officeAndUsers models.OfficeAndUsers, err error) {
	var office models.Office
	var officeUsers []models.OfficeUser

	if err = db.GetDBConn().Model(&models.Office{}).Where("id = ?", officeID).First(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetOfficesAndUsersById] Error while getting office: %v", err)
		return officeAndUsers, TranslateGormError(err)
	}

	officeUsers, err = GetAllOfficeUsers(uint(officeID))
	if err != nil {
		return officeAndUsers, TranslateGormError(err)
	}

	officeAndUsers.Office = office
	officeAndUsers.OfficeUsers = officeUsers

	return officeAndUsers, nil
}

func GetOfficeById(month, year uint, id int) (office models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Where("id = ?", id).
		Preload("OfficeUsers").
		Preload("OfficeUsers.Worker").
		Preload("OfficeUsers.Worker.CardSales", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("OfficeUsers.Worker.CardTurnovers", "EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Preload("OfficeUsers.Worker.User").First(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetOfficeById] Error while getting office: %v", err)

		return office, TranslateGormError(err)
	}

	return office, nil
}

func GetOfficeByIdOnlyOffice(id int) (office models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Where("id = ?", id).
		First(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetOfficeById] Error while getting office: %v", err)

		return office, TranslateGormError(err)
	}

	return office, nil
}

func GetOfficeByTitle(title string) (office models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Where("title = ?", title).First(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetOfficeByTitle] Error while getting office: %v", err)

		return office, TranslateGormError(err)
	}

	return office, nil
}

func CreateOffice(office models.Office) (err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Create(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.CreateOffice] Error while creating office: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func CreateOfficeTX(tx *gorm.DB, office models.Office) (err error) {
	if err = tx.Model(&models.Office{}).Create(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.CreateOffice] Error while creating office: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateOffice(office models.Office) (err error) {
	fmt.Println(office)
	if err = db.GetDBConn().Model(&models.Office{}).Where("id = ?", office.ID).Updates(office).Error; err != nil {
		logger.Error.Printf("[repositoroty.UpdateOffice] Error while updating office: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteOffice(officeID uint) (err error) {
	office, err := GetOfficeByIdOnlyOffice(int(officeID))
	if err != nil {
		if errors.Is(errs.ErrRecordNotFound, err) {
			return errs.ErrOfficeNotFound
		}

		return err
	}

	if err = db.GetDBConn().Model(&models.Office{}).Delete(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.DeleteOffice] Error while deleting office: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
