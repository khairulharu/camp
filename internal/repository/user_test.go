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

	t.Run("InsertRareUser", func(t *testing.T) {
		if err := userRepository.Insert(context.Background(), &testInsertCase); err != nil {
			t.Error(err)
		}
	})

	t.Run("FindByName", func(t *testing.T) {

		resultUser, err := userRepository.FindByUsername(context.Background(), testInsertCase.Name)

		if err != nil {
			t.Error(err)
		}

		if resultUser == (domain.User{}) {
			t.Error("error qery using username or not found")
		}

		userRes, err := userRepository.FindByID(context.Background(), resultUser.ID)

		if err != nil {
			t.Error(err)
		}

		if userRes == (domain.User{}) {
			t.Error("user not found")
		}

		if userRes.ID != resultUser.ID {
			t.Error("errro query rowww")
		}
	})

	t.Run("SelectAllUsers", func(t *testing.T) {
		users, err := userRepository.GetAll(context.Background())

		if err != nil {
			t.Error(err)
		}

		if users == nil {
			t.Error("users not found")
		}
	})
}
