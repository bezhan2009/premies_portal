package controllers

import (
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTestQuestions(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	question.TestID = uint(testID)

	if err := service.CreateTestQuestions(question); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Test Questions Created successfully",
	})
}

func UpdateTestQuestions(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	question.ID = uint(questionID)

	if err := service.UpdateTestQuestion(question); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Test Questions Updated successfully",
	})
}

func DeleteTestQuestions(c *gin.Context) {
	questionIDStr := c.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err = service.DeleteTestQuestion(uint(questionID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Test Questions Deleted successfully",
	})
}
