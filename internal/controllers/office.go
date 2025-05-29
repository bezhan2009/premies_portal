package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetAllOffices(c *gin.Context) {
	offices, err := service.GetAllOffices()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, offices)
}

func GetOfficeByID(c *gin.Context) {
	officeIDStr := c.Param("id")
	officeID, err := strconv.Atoi(officeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	office, err := service.GetOfficeById(officeID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, office)
}

func CreateOffice(c *gin.Context) {
	var office models.Office
	if err := c.ShouldBindJSON(&office); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateOffice(office)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "office created successfully",
	})
}

func UpdateOffice(c *gin.Context) {
	var office models.Office
	if err := c.ShouldBindJSON(&office); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.UpdateOffice(office)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "office updated successfully",
	})
}

func DeleteOffice(c *gin.Context) {
	officeIDStr := c.Param("id")
	officeID, err := strconv.Atoi(officeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteOffice(uint(officeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "office deleted successfully",
	})
}
