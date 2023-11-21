package controller

import (
	"gin-api/cmd/app/service"
	"gin-api/cmd/helpers/responses"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AuthController struct {
	BaseController
	UserService *service.UserService
}

func NewAuthController(
	log *zerolog.Logger,
	userService *service.UserService,
) *AuthController {
	return &AuthController{
		BaseController: BaseController{log: log},
		UserService:    userService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	users := c.UserService.GetListUser()
	message := "Success"

	responses.ResponseOk(ctx, responses.ResponseData{Message: &message, Data: users})
	return
}
