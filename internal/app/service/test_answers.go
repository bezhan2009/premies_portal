package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func GetTestAnswers() (answers []models.Answer, err error) {
	answers, err = repository.GetTestAnswers()
	if err != nil {
		return nil, err
	}

	return answers, nil
}

func GetTestAnswersByTestId(testID int) (answer []models.Answer, err error) {
	answer, err = repository.GetTestAnswersByTestId(testID)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func GetTestAnswersByAnswerId(answerID int) (answer []models.Answer, err error) {
	answer, err = repository.GetTestAnswersByAnswerId(answerID)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func CreateTestAnswers(answer []models.Answer) (err error) {
	err = repository.CreateTestAnswers(answer)
	if err != nil {
		return err
	}

	return nil
}
