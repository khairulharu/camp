package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"testing"
	"time"
)

func TestAvailabilityRepository(t *testing.T) {
	avalaibilityRepository := repository.NewAvailabilityRepository(*&dbGorm)

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
}
