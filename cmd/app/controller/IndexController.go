package controller

import (
	"gin-api/cmd/helpers/responses"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type IndexController struct {
	BaseController
}

func NewIndexController(log *zerolog.Logger) *IndexController {
	return &IndexController{BaseController: BaseController{log: log}}
}

func (c *IndexController) Hi(ctx *gin.Context) {
	message := "Hi There"

	responses.ResponseOk(ctx, responses.ResponseData{Message: &message})
	return
}

func (c *IndexController) Ping(ctx *gin.Context) {
	message := "Pong"
	c.log.Info().Msg("Ulalaaa")
	responses.ResponseOk(ctx, responses.ResponseData{Message: &message})
}
