package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	afterIDStr := c.Query("after")
	if afterIDStr == "" {
		afterIDStr = "0"
	}

	afterID, err := strconv.Atoi(afterIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidAfterID)
		return
	}

	users, err := service.GetAllUsers(uint(afterID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrInvalidID)
		return
	}

	user, err := service.GetUserByID(uint(userID))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] error: %v\n", err)
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetMyDataUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	user, err := service.GetUserByID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	user.ID = uint(userID)

	err = service.UpdateUser(user)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func UpdateUsersPassword(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	var newPassword models.NewUsersPassword
	if err := c.ShouldBindJSON(&newPassword); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if newPassword.NewPassword == "" {
		HandleError(c, errs.ErrPasswordIsEmpty)
		return
	}

	err := service.UpdateUserPassword(userID, newPassword.OldPassword, newPassword.NewPassword)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
