package repository

import (
	"errors"
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
)

func GetRoleByUserID(userID uint) (uint, error) {
	var roleID uint

	dbConn := db.GetDBConn()

	// Пытаемся найти роль напрямую по UserID
	err := dbConn.Model(&models.User{}).
		Select("role_id").
		Where("id = ?", userID).
		Scan(&roleID).Error

	if err == nil && roleID != 0 {
		return roleID, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, TranslateGormError(err)
	}

	// Не нашли пользователя, пробуем найти Worker
	var worker models.Worker
	err = dbConn.Model(&models.Worker{}).
		Select("user_id").
		Where("id = ?", userID).
		Take(&worker).Error
	if err != nil {
		return 0, TranslateGormError(err)
	}

	// Получаем RoleID по Worker.UserID
	err = dbConn.Model(&models.User{}).
		Select("role_id").
		Where("id = ?", worker.UserID).
		Scan(&roleID).Error
	if err != nil {
		return 0, TranslateGormError(err)
	}

	return roleID, nil
}
