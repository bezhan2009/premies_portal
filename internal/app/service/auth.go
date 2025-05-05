package service

import (
	"errors"
	"fmt"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
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

func SignUp(user models.User) (uint, error) {
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

	var userDB models.User

	if userDB, err = repository.CreateUser(user); err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

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
