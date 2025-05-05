package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func AddCardSales(sales models.CardSales) (uint, error) {
	cardSaleID, err := repository.AddCardSales(sales)
	if err != nil {
		return 0, err
	}

	return cardSaleID, nil
}

func UpdateCardSales(sales models.CardSales) error {
	err := repository.UpdateCardSales(sales)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCardSales(id uint) error {
	err := repository.DeleteCardSales(id)
	if err != nil {
		return err
	}

	return nil
}
