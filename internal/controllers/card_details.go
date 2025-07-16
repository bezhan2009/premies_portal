package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetStatisticsCards(c *gin.Context) {
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

	cardCharters, err := service.GetCardsStatistic(uint(month), uint(year))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, cardCharters)
}

func GetCardDetailsWorkers(c *gin.Context) {
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

	afterID := c.Query("after")

	var cardDetails []models.CardDetails

	cacheKey := GenerateRedisKeyFromQuery(c, fmt.Sprintf("card_details:%s:month:%d:year:%d", afterID, month, year))
	found, _ := db.GetCache(cacheKey, cardDetails)
	if found {
		c.JSON(http.StatusOK, cardDetails)
		return
	}

	cardDetails, err := service.GetCardDetailsWorkers(afterID, month, year)
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, cardDetails)

	c.JSON(http.StatusOK, cardDetails)
}

func GetCardDetailsWorkerByID(c *gin.Context) {
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

	workerIDStr := c.Param("id")
	workerID, err := strconv.Atoi(workerIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	afterID := c.Query("after")

	var cardDetails []models.CardDetails

	cacheKey := GenerateRedisKeyFromQuery(c, fmt.Sprintf("card_detail:%s:after:%s:month:%d:year:%d", workerIDStr, afterID, month, year))
	found, _ := db.GetCache(cacheKey, cardDetails)
	if found {
		c.JSON(http.StatusOK, cardDetails)
		return
	}

	cardDetails, err = service.GetCardDetailsWorker(uint(workerID), afterID, month, year)
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, cardDetails)

	c.JSON(http.StatusOK, cardDetails)
}

func GetMyCardDetailsWorker(c *gin.Context) {
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

	afterID := c.Query("after")

	var cardDetails []models.CardDetails

	cacheKey := GenerateRedisKeyFromQuery(c, fmt.Sprintf("card_detail:%d:after:%s:month:%d:year:%d", workerID, afterID, month, year))
	found, _ := db.GetCache(cacheKey, cardDetails)
	if found {
		c.JSON(http.StatusOK, cardDetails)
		return
	}

	cardDetails, err := service.GetCardDetailsWorker(workerID, afterID, month, year)
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, cardDetails)

	c.JSON(http.StatusOK, cardDetails)
}
