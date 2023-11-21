package configuration

import (
	"gin-api/cmd/app/controller"
	"gin-api/cmd/app/service"

	"github.com/rs/zerolog"
)

func LoadController(log *zerolog.Logger, service *service.ServiceList) controller.ControllerList {
	indexController := controller.NewIndexController(log)
	authController := controller.NewAuthController(log, service.User)
	return controller.ControllerList{
		Index: indexController,
		Auth:  authController,
	}
}
