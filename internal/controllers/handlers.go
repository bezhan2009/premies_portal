package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

// Обработка ошибок, которые приводят к статусу 400 (Bad Request)
func handleBadRequestErrors(err error) bool {
	return errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrUsernameIsRequired) ||
		errors.Is(err, errs.ErrPasswordIsRequired) ||
		errors.Is(err, errs.ErrRoleIsRequired) ||
		errors.Is(err, errs.ErrAlreadyAnswered) ||
		errors.Is(err, errs.ErrIncorrectAnswer) ||
		errors.Is(err, errs.ErrInvalidPhoneNumber) ||
		errors.Is(err, errs.ErrEmailUniquenessFailed) ||
		errors.Is(err, errs.ErrPhoneUniquenessFailed) ||
		errors.Is(err, errs.ErrAlreadyInOfficeWorkers) ||
		errors.Is(err, errs.ErrKnowledgeAlreadyExists) ||
		errors.Is(err, errs.ErrKnowledgeBaseUniquenessFailed) ||
		errors.Is(err, errs.ErrYouAreNotWorker) ||
		errors.Is(err, errs.ErrOfficeNameUniquenessFailed) ||
		errors.Is(err, errs.ErrYouAreWorker) ||
		errors.Is(err, errs.ErrWrongRoleID) ||
		errors.Is(err, errs.ErrInvalidBaseID) ||
		errors.Is(err, errs.ErrInvalidAfterID) ||
		errors.Is(err, errs.ErrUserIsNotDirector) ||
		errors.Is(err, errs.ErrInvalidMonth) ||
		errors.Is(err, errs.ErrInvalidYear) ||
		errors.Is(err, errs.ErrFirstNameIsRequired) ||
		errors.Is(err, errs.ErrLastNameIsRequired) ||
		errors.Is(err, errs.ErrEmailIsRequired) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrInvalidCredentials) ||
		errors.Is(err, errs.ErrPathParametrized) ||
		errors.Is(err, errs.ErrDirectorIDIsEmpty) ||
		errors.Is(err, errs.ErrUserIDIsEmpty) ||
		errors.Is(err, errs.ErrOfficeIDIsEmpty) ||
		errors.Is(err, errs.ErrInvalidPrice) ||
		errors.Is(err, errs.ErrInvalidID) ||
		errors.Is(err, errs.ErrInvalidField) ||
		errors.Is(err, errs.ErrEmailIsEmpty) ||
		errors.Is(err, errs.ErrPasswordIsEmpty) ||
		errors.Is(err, errs.ErrUsernameIsEmpty) ||
		errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrDeleteFailed) ||
		errors.Is(err, errs.ErrInvalidTitle) ||
		errors.Is(err, errs.ErrInvalidDescription) ||
		errors.Is(err, errs.ErrInvalidAmount)
}

// Обработка ошибок, которые приводят к статусу 404 (Not Found)
func handleNotFoundErrors(err error) bool {
	return errors.Is(err, errs.ErrRecordNotFound) ||
		errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrFileNotFound) ||
		errors.Is(err, errs.ErrOfficeNotFound) ||
		errors.Is(err, errs.ErrTempReportNotFound) ||
		errors.Is(err, errs.ErrKnowledgeBaseNotFound)
}

// Обработка ошибок, которые приводят к статусу 401 (Unauthorized)
func handleUnauthorizedErrors(err error) bool {
	return errors.Is(err, errs.ErrInvalidToken) ||
		errors.Is(err, errs.ErrUnauthorized) ||
		errors.Is(err, errs.ErrRefreshTokenExpired)
}

// Обработка ошибок, которые приводят к статусу 500 (Internal Server Error)
func handleInternalErrors(err error) bool {
	return errors.Is(err, errs.ErrSomethingWentWrong) ||
		errors.Is(err, errs.ErrTempReportNotFound)
}

// HandleError Основная функция обработки ошибок
func HandleError(c *gin.Context, err error) {
	if handleBadRequestErrors(err) {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	} else if handleNotFoundErrors(err) {
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	} else if handleUnauthorizedErrors(err) {
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	} else if handleInternalErrors(err) {
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	} else {
		logger.Error.Printf("[controllers.HandleError] Err: %s", err)
		c.JSON(http.StatusInternalServerError, newErrorResponse(errs.ErrSomethingWentWrong.Error()))
	}
}
