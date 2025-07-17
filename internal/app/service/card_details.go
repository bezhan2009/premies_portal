package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"time"
)

func GetCardsStatistic(month, year uint) (cardsCharters models.CardsCharters, err error) {
	cards, err := repository.GetAllCardDetailsByPeriod(month, year)
	if err != nil {
		return models.CardsCharters{}, err
	}

	cardStats, err := repository.GetCardStats(month, year)
	if err != nil {
		return models.CardsCharters{}, err
	}

	cardsCharters.CardsInGeneral = uint(len(cards))
	cardsCharters.DebtOsd = cardStats.DebtOsd
	cardsCharters.DebtOsk = cardStats.DebtOsk
	cardsCharters.OutBalance = cardStats.OutBalance
	cardsCharters.InBalance = cardStats.InBalance

	for _, card := range cards {
		if card.IssueDate.Month() == time.Month(month) && card.IssueDate.Year() == int(year) {
			cardsCharters.CardsForMonth += 1
		}

		if !card.IssueDate.IsZero() && card.DebtOsd > 0 {
			cardsCharters.ActivatedCards += 1
		}
	}

	return cardsCharters, nil
}

func GetCardDetailsWorkers(afterID string, month, year int) (cardDetails []models.CardDetails, err error) {
	if afterID == "" {
		afterID = "0"
	}

	cardDetails, err = repository.GetCardDetailsWorkers(afterID, month, year)
	if err != nil {
		return nil, err
	}

	return cardDetails, nil
}

func GetCardDetailsWorker(workerID uint, afterID string, month, year int) (cardDetails []models.CardDetails, err error) {
	if afterID == "" {
		afterID = "0"
	}

	worker, err := repository.GetWorkerByUserID(workerID)
	if err != nil {
		return nil, err
	}

	cardDetails, err = repository.GetCardDetailsWorker(worker.ID, afterID, month, year)
	if err != nil {
		return nil, err
	}

	return cardDetails, nil
}
