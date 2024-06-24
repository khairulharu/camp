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
		ID:          78,
		Name:        "testing_name",
		Description: "lorem ipsum dolor sit amet get another backend skill that providing and ruining another creatuition mean that is not impossible anymore that can be ruin everything can another mean is not possibel to containing othe creation",
	}

	t.Run("InsertNewActivity", func(t *testing.T) {
		if err := activityRepository.Insert(context.Background(), &activityTestCase); err != nil {
			t.Error(err)
		}
	})

	t.Run("FindById", func(t *testing.T) {
		activityResult, err := activityRepository.FindByID(context.Background(), activityTestCase.ID)

		if err != nil {
			t.Error(err.Error())
		}

		if activityResult != activityTestCase {
			t.Error("eror findById test case result different")
		}
	})

	t.Run("UpdateActivity", func(t *testing.T) {

		testUpdateActivity := domain.Activity{
			ID:          78,
			Name:        "testing_upddate_name",
			Description: "updated",
		}

		if err := activityRepository.Update(context.Background(), &testUpdateActivity); err != nil {
			t.Error(err)
		}

		resultUpdateActifity, _ := activityRepository.FindByID(context.Background(), testUpdateActivity.ID)

		if testUpdateActivity != resultUpdateActifity {
			t.Error("error when update test case")
		}
	})

	t.Run("GetAllActivity", func(t *testing.T) {
		results, err := activityRepository.GetAll(context.Background())

		if err != nil && results == nil {
			t.Error()
		}
	})
}
