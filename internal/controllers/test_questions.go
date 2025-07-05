package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func CreateTestQuestions(c *gin.Context) {
	var questions []models.Question
	if err := c.ShouldBindJSON(&questions); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateTestQuestions(questions); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Test Questions Created successfully",
	})
}

func UpdateTestQuestions(c *gin.Context) {
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

	question.ID = uint(testID)

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
