package repository

import (
	"errors"
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func GetAllUsersPag(afterID uint) (users []models.User, err error) {
	err = db.GetDBConn().
		Model(&models.User{}).
		Where("id > ?", afterID).
		Order("id ASC").
		Limit(security.AppSettings.AppLogicParams.PaginationParams.Limit).
		Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsersPag] error getting all users: %s\n", err.Error())
		return nil, TranslateGormError(err)
	}

	return users, nil
}

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllWorkersPag] error getting all users: %s\n", err.Error())
		return nil, TranslateGormError(err)
	}

	return users, nil
}

func GetUserByID(userID uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", userID).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errs.ErrUserNotFound
		}

		logger.Error.Printf("[repository.GetUserByID] error getting user by username: %v\n", err)
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
		return nil, TranslateGormError(err)
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

func CreateUser(tx *gorm.DB, user models.User) (userDB models.User, err error) {
	if err = tx.Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return userDB, TranslateGormError(err)
	}

	userDB = user
	return userDB, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errs.ErrRecordNotFound
		}
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByEmailAndPassword(email string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errs.ErrRecordNotFound
		}
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

func UpdateUser(user models.User) (err error) {
	if err = db.GetDBConn().Model(models.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		logger.Error.Printf("[repository.UpdateUser] Error while updating user: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
