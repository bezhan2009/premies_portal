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

	if len(tests) <= showTests {
		return tests, nil
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(tests), func(i, j int) {
		tests[i], tests[j] = tests[j], tests[i]
	})

	selectedTests := tests[:showTests]

	return selectedTests, nil
}

func GetTestById(id int) (test models.Test, err error) {
	test, err = repository.GetTestById(id)
	if err != nil {
		return test, err
	}

	return test, nil
}

func CreateTest(test []models.Test) (err error) {
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
