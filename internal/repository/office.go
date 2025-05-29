package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetAllOffices() (offices []models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Find(&offices).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetAllOffices] Error while getting all offices: %v", err)

		return nil, TranslateGormError(err)
	}

	return offices, nil
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

func GetOfficeById(id int) (office models.Office, err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Where("id = ?", id).First(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.GetOfficeById] Error while getting office: %v", err)

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

func UpdateOffice(office models.Office) (err error) {
	if err = db.GetDBConn().Model(&models.Office{}).Where("id = ?", office.ID).Updates(office).Error; err != nil {
		logger.Error.Printf("[repositoroty.UpdateOffice] Error while updating office: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteOffice(officeID uint) (err error) {
	office, err := GetOfficeById(int(officeID))
	if err != nil {
		return err
	}

	if err = db.GetDBConn().Model(&models.Office{}).Delete(&office).Error; err != nil {
		logger.Error.Printf("[repositoroty.DeleteOffice] Error while deleting office: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
