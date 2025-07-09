package controllers

import (
	"errors"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTestAnswers(c *gin.Context) {
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

	answers, err := service.GetTestAnswers(uint(month), uint(year))
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
	userID := c.GetUint(middlewares.UserIDCtx)

	var answers []models.Answer
	if err := c.ShouldBindJSON(&answers); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	for i, _ := range answers {
		answers[i].UserID = userID
	}

	if err := service.CreateTestAnswers(userID, answers); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, answers)
}

func AllowedAnswer(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	month, year := uint(time.Now().Month()), uint(time.Now().Year())

	_, err := repository.GetTestAnswersByUserID(userID, month, year)
	if err != nil {
		if !errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, err)
			return
		}
	}
	if err == nil {
		HandleError(c, errs.ErrAlreadyAnswered)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
