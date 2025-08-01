package repository

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
)

func AddMobileBankSale(sales models.MobileBankSales) (err error) {
	if err = db.GetDBConn().Model(&models.MobileBankSales{}).Create(&sales).Error; err != nil {
		logger.Error.Printf("Error in AddMobileBankSale: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateMobileBankSale(sales models.MobileBankSales) (err error) {
	if err = db.GetDBConn().Model(&models.MobileBankSales{}).Where("id = ?", sales.ID).Updates(&sales).Error; err != nil {
		logger.Error.Printf("Error in UpdateMobileBankSale: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteMobileBankSale(id int) (err error) {
	if err = db.GetDBConn().Model(&models.MobileBankSales{}).Where("id = ?", id).Delete(&models.MobileBankSales{}).Error; err != nil {
		logger.Error.Printf("Error in DeleteMobileBankSale: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
