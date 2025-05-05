package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func AddCardTurnovers(c *gin.Context) {
	var cardTurnovers models.CardTurnovers
	if err := c.ShouldBindJSON(&cardTurnovers); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardTurnoverID, err := service.AddCardTurnover(cardTurnovers)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"card_turnover_id": cardTurnoverID,
		"message":          "Added cardTurnover successfully",
	})
}

func UpdateCardTurnovers(c *gin.Context) {
	cardTurnoverStrID := c.Param("id")
	cardTurnoverID, err := strconv.Atoi(cardTurnoverStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var cardTurnover models.CardTurnovers
	if err := c.ShouldBindJSON(&cardTurnover); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardTurnover.ID = uint(cardTurnoverID)
	if err := service.UpdateCardTurnover(cardTurnover); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated card_turnover successfully",
	})
}

func DeleteCardTurnovers(c *gin.Context) {
	cardTurnoverStrID := c.Param("id")
	cardTurnoverID, err := strconv.Atoi(cardTurnoverStrID)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	if err := service.DeleteCardTurnover(uint(cardTurnoverID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete card turnover successfully",
	})
}
