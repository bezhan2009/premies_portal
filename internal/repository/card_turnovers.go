package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
)

func AddCardTurnover(turnover models.CardTurnovers) (id uint, err error) {
	if err = db.GetDBConn().Model(&models.CardTurnovers{}).Create(turnover).Error; err != nil {
		return 0, TranslateGormError(err)
	}

	return id, nil
}

func UpdateCardTurnover(turnovers models.CardTurnovers) (err error) {
	if err = db.GetDBConn().Model(&models.CardTurnovers{}).Updates(turnovers).Error; err != nil {
		return TranslateGormError(err)
	}

	return nil
}

func DeleteCardTurnover(turnoverID uint) (err error) {
	if err := db.GetDBConn().Delete(&models.CardSales{}, turnoverID).Error; err != nil {
		return TranslateGormError(err)
	}

	return nil
}
