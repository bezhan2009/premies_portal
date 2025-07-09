package service

import (
	"errors"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/errs"
	"time"

	"gorm.io/gorm"
)

func GetTestAnswers(month, year uint) (answers []models.Answer, err error) {
	answers, err = repository.GetTestAnswers(month, year)
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

func GetTestAnswersByAnswerId(answerID int) (answer models.Answer, err error) {
	answer, err = repository.GetTestAnswersByAnswerId(answerID)
	if err != nil {
		return models.Answer{}, err
	}

	return answer, nil
}

func isOptionIDInSlice(optionID uint, optionSlice []models.Option) bool {
	for _, option := range optionSlice {
		if option.ID == optionID {
			return true
		}
	}

	return false
}

func CalculateScore(totalTests, correctAnswers int) float64 {
	if totalTests <= 0 {
		return 0
	}
	if correctAnswers < 0 {
		correctAnswers = 0
	}
	if correctAnswers > totalTests {
		correctAnswers = totalTests // защита от переполнения
	}

	score := float64(correctAnswers) / float64(totalTests) * 10
	return score
}

func CreateTestAnswers(userID uint, answers []models.Answer) (err error) {
	month, year := uint(time.Now().Month()), uint(time.Now().Year())

	_, err = repository.GetTestAnswersByUserID(userID, month, year)
	if err != nil {
		if !errors.Is(err, errs.ErrRecordNotFound) {
			return err
		}
	}
	if err == nil {
		return errs.ErrAlreadyAnswered
	}

	correctAnswers := 0
	totalTests := security.AppSettings.AppLogicParams.TestsLogicParams.ShowTests

	for i, answer := range answers {
		answer.IsCorrectAnswer = false
		question, err := repository.GetQuestionByID(answer.QuestionID)
		if err != nil {
			return err
		}

		switch question.Type {
		case models.TextQuestion:
			if len(question.Options) < 1 {
				return errs.ErrIncorrectAnswer
			}
			if answer.TextAnswer != "" && answer.TextAnswer == question.Options[0].CorrectText {
				correctAnswers++
				answer.IsCorrectAnswer = true
			}

		case models.SingleChoiceQuestion:
			if len(answer.SelectedOptions) == 0 {
				break
			}
			var correctID uint
			for _, opt := range question.Options {
				if opt.IsCorrect {
					correctID = opt.ID
					break
				}
			}
			if answer.SelectedOptions[0].OptionID == correctID {
				correctAnswers++
				answer.IsCorrectAnswer = true
			}

		case models.MultipleChoiceQuestion:
			correctIDs := make(map[uint]struct{}, len(question.Options))
			for _, opt := range question.Options {
				if opt.IsCorrect {
					correctIDs[opt.ID] = struct{}{}
				}
			}
			if len(answer.SelectedOptions) == len(correctIDs) {
				allMatch := true
				for _, sel := range answer.SelectedOptions {
					if _, ok := correctIDs[sel.OptionID]; !ok {
						allMatch = false
						break
					}
				}
				if allMatch {
					answer.IsCorrectAnswer = true
					correctAnswers++
				}
			}
		}

		answers[i] = answer
	}

	worker, err := repository.GetWorkerByUserID(userID)
	if err != nil {
		return errs.ErrYouAreNotWorker
	}

	score := CalculateScore(int(totalTests), correctAnswers)

	serviceQualities, err := repository.GetServiceQualitiesByDateAndUserID(
		worker.ID,
		month,
		year,
	)
	if err != nil {
		return err
	}

	if len(serviceQualities) < 1 {
		sq := models.ServiceQuality{
			CallCenter:  0,
			Coefficient: 0,
			Complaint:   0,
			Tests:       score,
			WorkerID:    worker.ID,
		}

		_, err = repository.AddServiceQuality(&sq)
		if err != nil {
			return err
		}

		serviceQualities = append(serviceQualities, sq)
	}

	err = db.GetDBConn().Transaction(func(tx *gorm.DB) error {
		return repository.SaveScoresAndAnswers(tx, serviceQualities, score, answers)
	})

	return err
}
