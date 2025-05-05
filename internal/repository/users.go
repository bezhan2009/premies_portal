package repository

import (
	"errors"
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().
		Preload("CardTurnovers").
		Preload("CardSales").
		Preload("OperatingActive").
		Preload("ServiceQuality").
		Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, TranslateGormError(err)
	}

	return users, nil
}

func GetUserByID(month int, id string) (user models.User, err error) {
	err = db.GetDBConn().
		Preload("CardTurnovers", "EXTRACT(MONTH FROM created_at) = ?", month).
		Preload("CardSales", "EXTRACT(MONTH FROM created_at) = ?", month).
		Preload("OperatingActive", "EXTRACT(MONTH FROM created_at) = ?", month).
		Preload("ServiceQuality", "EXTRACT(MONTH FROM created_at) = ?", month).
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, TranslateGormError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func UserExists(username, email, phone string) (bool, bool, bool, error) {
	users, err := GetAllUsers()
	if err != nil {
		return false, false, false, err
	}

	var usernameExists, emailExists, phoneExists bool
	for _, user := range users {
		if user.Username == username {
			usernameExists = true
		}
		if user.Email == email {
			emailExists = true
		}
		if user.Phone == phone {
			phoneExists = true
		}
	}

	return usernameExists, emailExists, phoneExists, nil
}

func CreateUser(user models.User) (userDB models.User, err error) {
	//logger.Debug.Println(user.ID)
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return userDB, TranslateGormError(err)
	}

	//logger.Debug.Println(user.ID)
	userDB = user
	return userDB, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByEmailAndPassword(email string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByEmailAndPassword] error getting user by email and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByPhone(phone string) (user models.User, err error) {
	err = db.GetDBConn().Where("phone = ? AND password = ?", phone).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByPhone] error getting user by phone: %v\n", err)

		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByEmailPasswordAndUsername(username, email, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("email = ? AND password = ? AND username = ?", email, password, username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByEmailPasswordAndUsername] error getting user by username, email and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByEmail(email string) (user models.User, err error) {
	err = db.GetDBConn().Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByEmail] error getting user by email: %v\n", err)

		return user, TranslateGormError(err)
	}

	return user, nil
}
