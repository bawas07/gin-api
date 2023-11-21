package controller

import (
	"github.com/rs/zerolog"
)

type BaseController struct {
	log *zerolog.Logger
}

type ControllerList struct {
	Index *IndexController
	Auth  *AuthController
}
