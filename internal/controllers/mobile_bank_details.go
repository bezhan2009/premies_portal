package controllers

import (
	"fmt"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllWorkersMobileBankDetails(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

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

	cacheKey := GenerateAllWorkersRedisKey(c, fmt.Sprintf("workers_mbdt_cache:%d", userID))

	var MBDetails []models.MobileBankDetails

	found, _ := db.GetCache(cacheKey, &MBDetails)
	if found {
		c.JSON(http.StatusOK, MBDetails)
		return
	}

	MBDetails, err = service.GetAllWorkersMobileBankDetails(uint(afterID), userID, uint(month), uint(year))
	if err != nil {
		HandleError(c, err)
		return
	}

	_ = db.SetCache(cacheKey, MBDetails)

	c.JSON(http.StatusOK, MBDetails)
}
