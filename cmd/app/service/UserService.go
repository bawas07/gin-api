package service

import (
	"gin-api/cmd/app/model"
	"time"

	"github.com/rs/zerolog"
)

type UserService struct {
	BaseService
}

func NewUserService(log *zerolog.Logger) *UserService {
	return &UserService{BaseService: BaseService{log: log}}
}

func (s *UserService) GetListUser() []model.User {
	email := "user@email.com"
	users := []model.User{
		{
			ID:        "userid-1",
			Username:  "username",
			Password:  nil,
			Email:     &email,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}

	return users
}
