package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func AddServiceQuality(c *gin.Context) {
	var serviceQuality models.ServiceQuality
	if err := c.ShouldBindJSON(&serviceQuality); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	serviceqID, err := service.AddServiceQuality(serviceQuality)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"card_turnover_id": serviceqID,
		"message":          "Added Service Quality successfully",
	})
}

func UpdateServiceQuality(c *gin.Context) {
	serviceQualityStrID := c.Param("id")
	serviceQualityID, err := strconv.Atoi(serviceQualityStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var serviceQuality models.ServiceQuality
	if err := c.ShouldBindJSON(&serviceQuality); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	serviceQuality.ID = uint(serviceQualityID)
	if _, err := service.UpdateServiceQuality(serviceQuality); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated Service Quality successfully",
	})
}

func DeleteServiceQuality(c *gin.Context) {
	serviceQualityStrID := c.Param("id")
	serviceQualityID, err := strconv.Atoi(serviceQualityStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err := service.DeleteServiceQuality(uint(serviceQualityID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Service Quality successfully",
	})
}
