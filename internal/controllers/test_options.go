package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func CreateTestOptions(c *gin.Context) {
	var options []models.Option
	if err := c.ShouldBindJSON(&options); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateTestOptions(options); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Options created successfully",
	})
}

func UpdateTestOptions(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.Atoi(optionIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var option models.Option
	if err = c.ShouldBindJSON(&option); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	option.ID = uint(optionID)

	if err = service.UpdateTestOption(option); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Option updated successfully",
	})
}

func DeleteTestOptions(c *gin.Context) {
	optionIDStr := c.Param("id")
	optionID, err := strconv.Atoi(optionIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err = service.DeleteTestOption(uint(optionID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Option deleted successfully",
	})
}
