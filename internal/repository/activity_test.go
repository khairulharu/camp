package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"reflect"
	"testing"
)

func TestActivityRepository(t *testing.T) {
	//dbGorm Get from main_test.go there set a variable to containing database connection
	activityRepository := repository.NewActivityRepository(dbGorm)

	//create a mock data represent all data

	testCases := [4]struct {
		input    domain.Activity
		expected domain.Activity
		update   domain.Activity
		updated  domain.Activity
	}{
		{input: domain.Activity{
			ID:          2989,
			Name:        "khairul",
			Description: "mynname  is kahriul aswad",
		}},
		{expected: domain.Activity{
			ID:          2989,
			Name:        "khairul",
			Description: "mynname  is kahriul aswad",
		}},
		{update: domain.Activity{
			ID:          2989,
			Name:        "update",
			Description: "updated",
		}},
		{updated: domain.Activity{
			ID:          2989,
			Name:        "update",
			Description: "updated",
		}},
	}

	// activityTestCase := domain.Activity{
	// 	ID:          111,
	// 	Name:        "testing_name",
	// 	Description: "lorem ipsum dolor sit amet get another backend skill that providing and ruining another creatuition mean that is not impossible anymore that can be ruin everything can another mean is not possibel to containing othe creation",
	// }
	for _, testCase := range testCases {
		t.Run("InsertNewActivity", func(t *testing.T) {
			if err := activityRepository.Insert(context.Background(), &testCase.input); err != nil {
				t.Error(err)
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("FindById", func(t *testing.T) {
			activityResult, err := activityRepository.FindByID(context.Background(), testCase.input.ID)

			if err != nil {
				t.Error(err.Error())
			}

			if activityResult == (domain.Activity{}) {
				t.Error("not found orr errorr")
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("UpdateActivity", func(t *testing.T) {
			if err := activityRepository.Update(context.Background(), &testCase.update); err != nil {
				t.Error(err)
			}

			resultUpdateActifity, _ := activityRepository.FindByID(context.Background(), testCase.update.ID)

			if !reflect.DeepEqual(resultUpdateActifity, testCase.updated) {
				t.Error("FailTest Update Activity")
			}
		})
	}

	t.Run("GetAllActivity", func(t *testing.T) {
		results, err := activityRepository.GetAll(context.Background())

		if err != nil && results == nil {
			t.Error()
		}
	})
}
