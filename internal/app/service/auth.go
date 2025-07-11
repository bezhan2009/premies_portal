package service

import (
	"errors"
	"fmt"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
	"premiesPortal/pkg/utils"
)

func SignIn(userDataCheck, password string) (user models.User, accessToken string, refreshToken string, err error) {
	if userDataCheck == "" {
		return user, "", "", errs.ErrInvalidData
	}

	user, err = repository.GetUserByEmailAndPassword(userDataCheck, password)
	if err != nil {
		if !errors.Is(err, errs.ErrRecordNotFound) {
			return user, "", "", err
		}

		user, err = repository.GetUserByUsernameAndPassword(userDataCheck, password)
		if err != nil {
			if !errors.Is(err, errs.ErrRecordNotFound) {
				return user, "", "", err
			}

			return user, "", "", errs.ErrInvalidCredentials
		}
	}

	accessToken, refreshToken, err = utils.GenerateToken(user.ID, uint(user.RoleID), user.Username)
	if err != nil {
		return user, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func SignUp(user models.User, worker models.Worker, office models.Office) (uint, error) {
	if err := validators.SignUpValidator(user); err != nil {
		return 0, err
	}

	usernameExists, emailExists, phoneExists, err := repository.UserExists(user.Username, user.Email, user.Phone)
	if err != nil {
		return 0, fmt.Errorf("failed to check existing user: %w", err)
	}

	if user.Password == "" || user.Email == "" || user.Username == "" {
		return 0, errs.ErrInvalidData
	}

	if usernameExists {
		logger.Error.Printf("user with username %s already exists", user.Username)
		return 0, errs.ErrUsernameUniquenessFailed
	}

	if emailExists {
		logger.Error.Printf("user with email %s already exists", user.Email)
		return 0, errs.ErrEmailUniquenessFailed
	}

	if phoneExists {
		logger.Error.Printf("user with phone %s already exists", user.Phone)
		return 0, errs.ErrPhoneUniquenessFailed
	}

	user.Password = utils.GenerateHash(user.Password)

	tx := db.GetDBConn().Begin()

	var userDB models.User

	if userDB, err = repository.CreateUser(tx, user); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	if user.RoleID == 6 || user.RoleID == 8 {
		office, err := repository.GetOfficeByTitle(worker.PlaceWork)
		if err != nil {
			tx.Rollback()
			return 0, errs.ErrOfficeNotFound
		}

		worker.UserID = userDB.ID

		if err = CreateWorker(tx, office.ID, worker); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if user.RoleID == 5 {
		if office.Title == "" {
			return 0, errs.ErrValidationFailed
		}

		office.DirectorID = int(userDB.ID)
		if err = repository.CreateOfficeTX(tx, office); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()

	return userDB.ID, nil
}

func SignInWithGoogle(userRequest models.User) (user uint, accessToken string, refreshToken string, err error) {
	u, err := repository.GetUserByEmail(userRequest.Email)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return user, accessToken, refreshToken, errs.ErrInvalidCredentials
		}

		return user, accessToken, refreshToken, err
	}

	accessToken, refreshToken, err = utils.GenerateToken(u.ID, uint(u.RoleID), u.Username)
	if err != nil {
		return user, "", "", err
	}

	return u.ID, accessToken, refreshToken, nil
}
