package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"testing"
	"time"
)

func TestAvailabilityRepository(t *testing.T) {
	avalaibilityRepository := repository.NewAvailabilityRepository(dbGorm)

	testInsertCase := domain.Availability{
		ID:                23,
		CampsiteID:        23,
		Date:              time.Now(),
		QuantityAvailable: 3,
	}

	t.Run("TestInsert", func(t *testing.T) {
		if err := avalaibilityRepository.Insert(context.Background(), &testInsertCase); err != nil {
			t.Error("errorinsertedvalue")
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		testUpdateCase := domain.Availability{
			ID:                23,
			CampsiteID:        23,
			Date:              time.Now(),
			QuantityAvailable: 5,
		}

		if err := avalaibilityRepository.Update(context.Background(), &testUpdateCase); err != nil {
			t.Error("errorWhenUpdtaingAvailability")
		}
	})

	t.Run("TestFindByID", func(t *testing.T) {
		availabilityResult, err := avalaibilityRepository.FindByID(context.Background(), testInsertCase.ID)

		if err != nil {
			t.Error(err.Error())
		}

		if availabilityResult == (domain.Availability{}) {
			t.Error("not found response error ")
		}

		if availabilityResult.ID != testInsertCase.ID {
			t.Error("error")
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		availabities, err := avalaibilityRepository.GetAll(context.Background())

		if err != nil {
			t.Error(err.Error())
		}

		if availabities == nil {
			t.Error("error get all availabilites")
		}
	})
}
