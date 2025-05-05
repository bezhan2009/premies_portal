package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func AddCardSales(c *gin.Context) {
	var cardSale models.CardSales
	if err := c.ShouldBindJSON(&cardSale); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardSaleID, err := service.AddCardSales(cardSale)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"card_sale_id": cardSaleID,
		"message":      "Added card_sale successfully",
	})
}

func UpdateCardSales(c *gin.Context) {
	cardSaleStrID := c.Param("card_sale_id")
	cardSaleID, err := strconv.Atoi(cardSaleStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var cardSale models.CardSales
	if err := c.ShouldBindJSON(&cardSale); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardSale.ID = uint(cardSaleID)
	if err := service.UpdateCardSales(cardSale); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated card_sale successfully",
	})
}

func DeleteCardSales(c *gin.Context) {
	cardSaleStrID := c.Param("card_sale_id")
	cardSaleID, err := strconv.Atoi(cardSaleStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err := service.DeleteCardSales(uint(cardSaleID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete card sale",
	})
}
