package controller

import (
	"gin-api/cmd/helpers/responses"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	message := "Hi There"

	responses.ResponseOk(c, responses.ResponseData{Message: &message})
	return
}
