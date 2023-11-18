package controller

import (
	"gin-api/cmd/helpers/responses"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func IndexController(log *zerolog.Logger) *BaseController {
	return &BaseController{log}
}

func (c *BaseController) Hi(ctx *gin.Context) {
	message := "Hi There"

	responses.ResponseOk(ctx, responses.ResponseData{Message: &message})
	return
}

func (c *BaseController) Ping(ctx *gin.Context) {
	message := "Pong"
	c.log.Info().Msg("Ulalaaa")
	responses.ResponseOk(ctx, responses.ResponseData{Message: &message})
}
