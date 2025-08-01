package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func AddMobileBankSale(c *gin.Context) {
	var mobSales models.MobileBankSales
	if err := c.ShouldBind(&mobSales); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.AddMobileBankSale(mobSales)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully add mobile bank sale",
	})
}

func UpdateMobileBankSale(c *gin.Context) {
	MobStrID := c.Param("id")
	MobID, err := strconv.Atoi(MobStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var mobSales models.MobileBankSales
	if err = c.ShouldBind(&mobSales); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	mobSales.ID = uint(MobID)

	err = service.UpdateMobileBankSale(mobSales)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update mobile bank sale",
	})
}

func DeleteMobileBankSale(c *gin.Context) {
	MobStrID := c.Param("id")
	MobID, err := strconv.Atoi(MobStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteMobileBankSale(MobID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete mobile bank sale",
	})
}
