package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"testing"
	"time"
)

func TestUserRepositoryRare(t *testing.T) {
	userRepositoryRare := repository.NewUserRepositoryRare(nil, dbGorm)

	testInsertCase := domain.User{
		Name:        "kahirul",
		Email:       "aswad@gmail.commmm",
		Password:    "sjshdjs",
		PhoneNumber: "0888888888888",
		Address:     "akksjad, planet mars",
		CreatedAt:   time.Now(),
	}

	t.Run("InsertRareUser", func(t *testing.T) {
		if err := userRepositoryRare.Insert(context.Background(), &testInsertCase); err != nil {
			t.Log(err)
		}
	})
}
