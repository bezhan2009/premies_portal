package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetTestAnswers(c *gin.Context) {
	answers, err := service.GetTestAnswers()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, answers)
}

func GetTestAnswersByTestId(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	answers, err := service.GetTestAnswersByTestId(testID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, answers)
}

func GetTestAnswersByAnswerId(c *gin.Context) {
	answerIDStr := c.Param("id")
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	answer, err := service.GetTestAnswersByAnswerId(answerID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, answer)
}

func CreateTestAnswers(c *gin.Context) {
	var answer []models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateTestAnswers(answer); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, answer)
}
