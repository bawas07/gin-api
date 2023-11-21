package configuration

import (
	"gin-api/cmd/app/repository"
	"gin-api/cmd/app/service"

	"github.com/rs/zerolog"
)

func LoadService(log *zerolog.Logger, repo *repository.RepoList) service.ServiceList {
	userService := service.NewUserService(log, repo.UserRepo)
	return service.ServiceList{
		User: userService,
	}
}
