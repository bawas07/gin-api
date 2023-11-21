package service

import (
	"gin-api/cmd/app/model"
	"gin-api/cmd/app/repository"

	"github.com/rs/zerolog"
)

type UserService struct {
	BaseService
	UserRepo *repository.UserRepository
}

func NewUserService(
	log *zerolog.Logger,
	userRepo *repository.UserRepository,
) *UserService {
	return &UserService{
		BaseService: BaseService{log: log},
		UserRepo:    userRepo,
	}
}

func (s *UserService) GetListUser() []model.User {
	// email := "user@email.com"
	// users := []model.User{
	// 	{
	// 		ID:        "userid-1",
	// 		Username:  "username",
	// 		Password:  nil,
	// 		Email:     &email,
	// 		CreatedAt: time.Now().UTC(),
	// 		UpdatedAt: time.Now().UTC(),
	// 	},
	// }
	users := s.UserRepo.GetAllUser()

	return users
}
