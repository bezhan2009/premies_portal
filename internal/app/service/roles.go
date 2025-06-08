package service

import "premiesPortal/internal/repository"

func GetRoleByUserID(userID uint) (roleID uint, err error) {
	roleID, err = repository.GetRoleByUserID(userID)
	if err != nil {
		return 0, err
	}

	return roleID, nil
}
