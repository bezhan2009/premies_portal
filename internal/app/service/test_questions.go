package service

import (
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
)

func CreateTestQuestions(questions []models.Question) (err error) {
	err = repository.CreateTestQuestions(questions)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTestQuestion(question models.Question) (err error) {
	err = repository.UpdateTestQuestion(question)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTestQuestion(questionID uint) (err error) {
	err = repository.DeleteTestQuestion(questionID)
	if err != nil {
		return err
	}

	return nil
}
