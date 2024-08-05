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

	var testCaseUserId int64

	t.Run("InsertRareUser", func(t *testing.T) {
		if err := userRepository.Insert(context.Background(), &testInsertCase); err != nil {
			t.Error(err)
		}
	})

	t.Run("FindByName", func(t *testing.T) {

		resultUser, err := userRepository.FindByUsername(context.Background(), "kahirul")

		if err != nil {
			t.Error(err)
		}

		if resultUser == (domain.User{}) {
			t.Error("error qery using username")
		}

		testCaseUserId = resultUser.ID
	})

	t.Run("FindByID", func(t *testing.T) {
		userRes, err := userRepository.FindByID(context.Background(), int64(testCaseUserId))
		if err != nil {
			t.Error(err)
		}

		if userRes != (domain.User{}) {
			t.Error(userRes)
		}

		if userRes.ID != testInsertCase.ID {
			t.Error("errro query rowww")
		}
	})
}
