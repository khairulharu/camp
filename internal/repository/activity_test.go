package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"testing"
)

func TestActivityRepository(t *testing.T) {
	//dbGorm Get from main_test.go there set a variable to containing database connection
	activityRepository := repository.NewActivityRepository(dbGorm)

	//create a mock data represent all data

	activityTestCase := domain.Activity{}

	t.Run("InsertNewActivity", func(t *testing.T) {
		if err := activityRepository.Insert(context.Background(), &activityTestCase); err != nil {
			t.Error(err)
		}
	})
}
