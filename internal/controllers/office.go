package controllers

import (
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllOffices(c *gin.Context) {
	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidMonth)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidYear)
		return
	}

	offices, err := service.GetAllOffices(uint(month), uint(year))
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

	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidMonth)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidYear)
		return
	}

	office, err := service.GetOfficeById(uint(month), uint(year), officeID)
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
	officeIDStr := c.Param("id")
	officeID, err := strconv.Atoi(officeIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var office models.Office
	if err := c.ShouldBindJSON(&office); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	office.ID = uint(officeID)

	err = service.UpdateOffice(office)
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
