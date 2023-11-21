package configuration

import (
	"gin-api/cmd/app/service"

	"github.com/rs/zerolog"
)

func LoadService(log *zerolog.Logger) service.ServiceList {
	userService := service.NewUserService(log)
	return service.ServiceList{
		User: userService,
	}
}
