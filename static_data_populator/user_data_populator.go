package static_data_populator

import (
	"context"
	"errors"
	"interview/models"
)

type UserData interface {
	GetUseData(ctx context.Context, userKey int64) (models.User, error)
}

type userData struct {
	userData []models.User
}

func (u userData) GetUseData(ctx context.Context, userKey int64) (models.User, error) {
	for _, v := range u.userData {
		if v.Key == userKey {
			return v, nil
		}
	}
	return models.User{}, errors.New("record not found")
}

func NewUserData() UserData {
	return &userData{
		userData: []models.User{
			{Key: 1, Name: "Rahul"},
			{Key: 2, Name: "Sam"},
		},
	}
}
