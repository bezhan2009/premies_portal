package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
)

func AddCardSales(sales models.CardSales) (uint, error) {
	if err := db.GetDBConn().Create(&sales).Error; err != nil {

		return sales.ID, TranslateGormError(err)
	}

	return sales.ID, nil
}

func UpdateCardSales(sales models.CardSales) error {
	if err := db.GetDBConn().Model(&models.CardSales{}).Save(&sales).Error; err != nil {
		return TranslateGormError(err)
	}

	return nil
}

func DeleteCardSales(id uint) error {
	if err := db.GetDBConn().Delete(&models.CardSales{}, id).Error; err != nil {
		return TranslateGormError(err)
	}

	return nil
}
