package configuration

import (
	"gin-api/cmd/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var db = make(map[string]string)

func ApiRouteV1(r *gin.RouterGroup, log *zerolog.Logger) {
	v1 := r.Group("/v1")

	indexController := controller.IndexController(log)
	v1.GET("/ping", indexController.Ping)

	// Get user value
	v1.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})
}
