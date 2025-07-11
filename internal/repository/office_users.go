package repository

import (
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func GetAllOfficeUsers(officeID uint) (officeUsers []models.OfficeUser, err error) {
	if err = db.GetDBConn().
		Model(&models.OfficeUser{}).
		Where("office_id = ?", officeID).
		Preload("Worker").
		Preload("Worker.User").
		Find(&officeUsers).Error; err != nil {
		logger.Error.Printf("[repository.GetAllOfficeUsers] Error while getting all office users: %v\n", err)
		return officeUsers, TranslateGormError(err)
	}

	return officeUsers, nil
}

func GetOfficeUserByID(officeUserID uint) (officeUser models.OfficeUser, err error) {
	if err = db.GetDBConn().
		Model(&models.OfficeUser{}).
		Where("id = ?", officeUserID).
		Preload("Worker").
		Preload("Worker.User").
		First(&officeUser).Error; err != nil {
		logger.Error.Printf("[repository.GetOfficeUserByID] Error while getting office user: %v\n", err)

		return officeUser, TranslateGormError(err)
	}

	return officeUser, nil
}

func GetOfficeWorkerByUserIDAndOfficeID(officeID, workerID uint) (officeUser models.OfficeUser, err error) {
	if err = db.GetDBConn().
		Model(&models.OfficeUser{}).
		Where("office_id = ? AND worker_id = ?", officeID, workerID).
		Preload("Worker").
		First(&officeUser).Error; err != nil {
		logger.Error.Printf("[repository.GetOfficeWorkerByUserIDAndOfficeID] Error while getting office user: %v\n", err)

		return officeUser, TranslateGormError(err)
	}

	return officeUser, nil
}

func AddUserToOffice(officeUser *models.OfficeUser) (err error) {
	if err = db.GetDBConn().Model(&models.OfficeUser{}).Create(officeUser).Error; err != nil {
		logger.Error.Printf("[repository.AddUserToOffice] Error while adding user to office: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func AddUserToOfficeByTX(tx *gorm.DB, officeUser *models.OfficeUser) (err error) {
	if err = tx.Model(&models.OfficeUser{}).Create(officeUser).Error; err != nil {
		logger.Error.Printf("[repository.AddUserToOffice] Error while adding user to office: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteUserFromOffice(officeUserID uint) (err error) {
	officeUser, err := GetOfficeUserByID(officeUserID)
	if err != nil {
		return err
	}

	if err = db.GetDBConn().Model(&models.OfficeUser{}).Delete(&officeUser).Error; err != nil {
		logger.Error.Printf("[repository.DeleteUserFromOffice] Error while deleting user from office: %v\n", err)

		return TranslateGormError(err)
	}

	return nil
}
