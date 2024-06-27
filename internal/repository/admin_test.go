package repository_test

import (
	"campsite/internal/domain"
	"campsite/internal/repository"
	"context"
	"testing"
)

func TestAdminRepository(t *testing.T) {

	adminRepository := repository.NewAdminRepository(dbGorm)

	testCases := []struct {
		input      domain.Admin
		updateData domain.Admin
		expected   domain.Admin
	}{
		{input: domain.Admin{
			ID:       1005,
			Name:     "jkjkjkj",
			Email:    "jkjjj@gmail.com",
			Password: "hghhskjhfhjsih",
		}},
		{expected: domain.Admin{
			ID:       477,
			Name:     "data_updated",
			Email:    "updated@gmail.com",
			Password: "updated",
		}},
		{updateData: domain.Admin{
			ID:       1005,
			Name:     "djSIAL;FSNSLFKNA",
			Email:    "DIUBIFUbFIUB@gmail.com",
			Password: "QWFUbaIFUBAfipuQ",
		}},
	}

	for _, testCase := range testCases {
		t.Run("InsertData", func(t *testing.T) {
			if err := adminRepository.Insert(context.Background(), &testCase.input); err != nil {
				t.Errorf("Fail InsertData Get, %s", err.Error())
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("TestUpdateData", func(t *testing.T) {
			if err := adminRepository.Update(context.Background(), &testCase.updateData); err != nil {
				t.Errorf("TestUpdatedData Fail Get, %s", err.Error())
			}
		})
	}
}
