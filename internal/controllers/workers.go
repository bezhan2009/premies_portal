package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
	"strconv"
)

func GetAllWorkers(c *gin.Context) {
	afterIDStr := c.Query("after")
	if afterIDStr == "" {
		afterIDStr = "0"
	}

	afterID, err := strconv.Atoi(afterIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidAfterID)
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

	workers, err := service.GetAllWorkers(uint(afterID), uint(month), uint(year))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"workers": workers,
	})
}

func GetWorkerByID(c *gin.Context) {
	workerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetWorkerByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrInvalidID)
		return
	}

	monthStr := c.Query("month")
	if monthStr == "" {
		monthStr = "0"
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidMonth)
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

	roleID, err := service.GetRoleByUserID(uint(workerID))
	if err != nil {
		HandleError(c, err)
		return
	}

	worker, err := service.GetWorkerByID(uint(workerID), roleID, uint(month), uint(year))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, worker)
}

func GetMyDataWorker(c *gin.Context) {
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

	workerID := c.GetUint(middlewares.UserIDCtx)
	roleID := c.GetUint(middlewares.UserRoleIDCtx)

	worker, err := service.GetWorkerByID(workerID, roleID, uint(month), uint(year))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, worker)
}
