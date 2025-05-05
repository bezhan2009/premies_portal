package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func AddOperatingActive(c *gin.Context) {
	var operatingActive models.OperatingActive
	if err := c.ShouldBindJSON(&operatingActive); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	operatingActiveID, err := service.AddOperatingActive(operatingActive)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"operating_active_id": operatingActiveID,
		"message":             "Added operating Active successfully",
	})
}

func UpdateOperatingActive(c *gin.Context) {
	operatingActiveStrID := c.Param("id")
	operatingActiveID, err := strconv.Atoi(operatingActiveStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var operatingActive models.OperatingActive
	if err := c.ShouldBindJSON(&operatingActive); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	operatingActive.ID = uint(operatingActiveID)
	if _, err := service.UpdateOperatingActive(operatingActive); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated operating Active successfully",
	})
}

func DeleteOperatingActive(c *gin.Context) {
	operatingActiveStrID := c.Param("id")
	operatingActiveID, err := strconv.Atoi(operatingActiveStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err := service.DeleteOperatingActive(uint(operatingActiveID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete operating Active successfully",
	})
}
