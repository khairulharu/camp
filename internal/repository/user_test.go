package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestUserRepositoryRare(t *testing.T) {
	userRepository := repository.UseUserRepository(false, nil, dbRareTest)

	testInsertCase := domain.User{
		Name:        "kahirul",
		Email:       "aswad@gmail.commmm",
		Password:    "sjshdjs",
		PhoneNumber: "0888888888888",
		Address:     "akksjad, planet mars",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	testCaseUserId := 16

	t.Run("InsertRareUser", func(t *testing.T) {
		if err := userRepository.Insert(context.Background(), &testInsertCase); err != nil {
			t.Error(err)
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		userRes, err := userRepository.FindByID(context.Background(), int64(testCaseUserId))
		if err != nil {
			t.Error(err)
		}

		if userRes != (domain.User{}) {
			t.Error(userRes)
		}
	})
}
