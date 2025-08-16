package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/db"
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

	cacheKey := GenerateAllWorkersRedisKey(c, "workers_cache")

	var workers []models.Worker

	found, _ := db.GetCache(cacheKey, &workers)
	if found {
		c.JSON(http.StatusOK, gin.H{
			"workers": workers,
		})
		return
	}

	preloadOptions := parsePreloadQueryParams(c)

	workers, err = service.GetAllWorkers(uint(afterID), uint(month), uint(year), preloadOptions)
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, workers)

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

	cacheKey := GenerateRedisKeyFromQuery(c, "worker_cache")

	var worker models.Worker

	found, _ := db.GetCache(cacheKey, &worker)
	if found {
		c.JSON(http.StatusOK, gin.H{"worker": worker})
		return
	}

	preloadOptions := parsePreloadQueryParams(c)

	worker, err = repository.GetWorkerByID(uint(month), uint(year), uint(workerID), preloadOptions)
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, worker)

	c.JSON(http.StatusOK, gin.H{"worker": worker})
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

	preloadOptions := parsePreloadQueryParams(c)

	userID := c.GetUint(middlewares.UserIDCtx)
	roleID := c.GetUint(middlewares.UserRoleIDCtx)

	var worker models.Worker

	worker, err := repository.GetWorkerByUserID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	cacheKey := GenerateRedisKeyFromQuery(c, fmt.Sprintf("worker_cache:%d", worker.ID))

	found, _ := db.GetCache(cacheKey, &worker)
	if found {
		c.JSON(http.StatusOK, worker)
		return
	}

	worker, err = service.GetWorkerByID(worker.ID, roleID, uint(month), uint(year), preloadOptions)
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, worker)

	c.JSON(http.StatusOK, worker)
}

func UpdateWorker(c *gin.Context) {
	workerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetWorkerByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var worker models.Worker
	if err = c.ShouldBindJSON(&worker); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	worker.ID = uint(workerID)

	err = service.UpdateWorkerByID(worker)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "worker updated successfully"})
}

func DeleteWorker(c *gin.Context) {
	workerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetWorkerByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteWorkerByID(uint(workerID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "worker has been deleted successfully"})
}
