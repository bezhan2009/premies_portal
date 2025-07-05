package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetAllTests(c *gin.Context) {
	tests, err := service.GetAllTests()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, tests)
}

func GetTestsForWorker(c *gin.Context) {
	tests, err := service.GetTestsForWorker(int(security.AppSettings.AppLogicParams.TestsLogicParams.ShowTests))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, tests)
}

func GetTestById(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	test, err := service.GetTestById(testID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, test)
}

func CreateTest(c *gin.Context) {
	var test []models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateTest(test); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created tests"})
}

func UpdateTest(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	test.ID = uint(testID)

	if err := service.UpdateTest(test); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated tests"})
}

func DeleteTest(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err := service.DeleteTest(testID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted tests"})
}
