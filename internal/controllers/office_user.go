package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetAllOfficeWorkers(c *gin.Context) {
	officeIDStr := c.Param("id")
	officeID, err := strconv.Atoi(officeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	officeUsers, err := service.GetAllOfficeUsers(uint(officeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, officeUsers)
}

func GetOfficeWorkerByID(c *gin.Context) {
	officeUserIDStr := c.Param("id")
	officeUserID, err := strconv.Atoi(officeUserIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	officeUser, err := service.GetOfficeUserById(uint(officeUserID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, officeUser)
}

func AddWorkerToOffice(c *gin.Context) {
	var officeUser models.OfficeUser
	if err := c.ShouldBindJSON(&officeUser); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.AddUserToOffice(&officeUser)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, officeUser)
}

func DeleteUserFromOffice(c *gin.Context) {
	officeUserIDStr := c.Param("id")
	officeUserID, err := strconv.Atoi(officeUserIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteUserFromOffice(uint(officeUserID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted user from office",
	})
}
