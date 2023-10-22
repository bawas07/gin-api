package controller

import (
	"gin-api/cmd/helpers"

	"github.com/gin-gonic/gin"
)

func Hi(c *gin.Context) {
	message := "Hi There"

	helpers.ResponseOk(c, helpers.ResponseData{Message: &message})
	return
}
