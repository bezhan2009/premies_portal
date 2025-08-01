package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func AddMobileBankSale(sales models.MobileBankSales) (err error) {
	err = repository.AddMobileBankSale(sales)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMobileBankSale(sales models.MobileBankSales) (err error) {
	err = repository.UpdateMobileBankSale(sales)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMobileBankSale(id int) (err error) {
	err = repository.DeleteMobileBankSale(id)
	if err != nil {
		return err
	}

	return nil
}
