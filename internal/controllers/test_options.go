package controllers

import (
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTestOptions(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var option models.Option
	if err = c.ShouldBindJSON(&option); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	option.QuestionID = uint(questionID)

	if err = service.CreateTestOptions(option); err != nil {
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
