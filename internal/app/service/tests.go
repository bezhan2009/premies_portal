package service

import (
	"math/rand"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/repository"
	"time"
)

func GetAllTests() (tests []models.Test, err error) {
	tests, err = repository.GetAllTests()
	if err != nil {
		return nil, err
	}

	return tests, nil
}

func GetTestsForWorker(showTests int) ([]models.Test, error) {
	tests, err := repository.GetAllTests()
	if err != nil {
		return nil, err
	}

	// Соберем все вопросы во flat-массив, сохраняя ссылку на тест
	type QuestionWithTest struct {
		Test     *models.Test
		Question models.Question
	}

	var allQuestions []QuestionWithTest
	for i := range tests {
		test := &tests[i]

		// Если у тебя тест содержит не один вопрос, а слайс Questions:
		for _, question := range test.Questions {
			allQuestions = append(allQuestions, QuestionWithTest{
				Test:     test,
				Question: question,
			})
		}

		// Если всё-таки в Test один вопрос (Question), то:
		/*
			allQuestions = append(allQuestions, QuestionWithTest{
				Test:     test,
				Question: test.Question,
			})
		*/
	}

	// Если вопросов меньше или столько, сколько надо — берем все
	if len(allQuestions) <= showTests {
		return tests, nil
	}

	// Перемешаем вопросы
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allQuestions), func(i, j int) {
		allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
	})

	// Выбираем нужное количество вопросов
	selected := allQuestions[:showTests]

	// Группируем выбранные вопросы обратно по тестам
	testMap := make(map[uint]*models.Test) // uint — ID теста

	for _, q := range selected {
		tid := q.Test.ID

		// Если тест еще не в мапе — создаем копию теста с пустыми Questions
		if _, exists := testMap[tid]; !exists {
			// Создаем shallow copy теста
			newTest := *q.Test
			newTest.Questions = nil
			testMap[tid] = &newTest
		}

		testMap[tid].Questions = append(testMap[tid].Questions, q.Question)
	}

	// Собираем результаты в слайс
	var result []models.Test
	for _, t := range testMap {
		result = append(result, *t)
	}

	return result, nil
}

func GetTestById(id int) (test models.Test, err error) {
	test, err = repository.GetTestById(id)
	if err != nil {
		return test, err
	}

	return test, nil
}

func CreateTest(test models.Test) (err error) {
	err = repository.CreateTest(test)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTest(test models.Test) (err error) {
	err = repository.UpdateTest(test)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTest(id int) (err error) {
	err = repository.DeleteTest(id)
	if err != nil {
		return err
	}

	return nil
}
