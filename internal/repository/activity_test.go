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

	activityTestCase := domain.Activity{
		Name:        "testing_name",
		Description: "lorem ipsum dolor sit amet get another backend skill that providing and ruining another creatuition mean that is not impossible anymore that can be ruin everything can another mean is not possibel to containing othe creation",
	}

	t.Run("InsertNewActivity", func(t *testing.T) {
		if err := activityRepository.Insert(context.Background(), &activityTestCase); err != nil {
			t.Error(err)
		}
	})
}
