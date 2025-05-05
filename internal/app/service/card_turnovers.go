package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func AddCardTurnover(turnover models.CardTurnovers) (id uint, err error) {
	id, err = repository.AddCardTurnover(turnover)
	if err != nil {
		return id, err
	}

	return id, nil
}

func UpdateCardTurnover(turnovers models.CardTurnovers) (err error) {
	err = repository.UpdateCardTurnover(turnovers)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCardTurnover(turnoverID uint) (err error) {
	err = repository.DeleteCardTurnover(turnoverID)
	if err != nil {
		return err
	}

	return nil
}
