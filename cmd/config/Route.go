package configuration

import (
	"gin-api/cmd/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var db = make(map[string]string)

func ApiRouteV1(
	r *gin.RouterGroup,
	log *zerolog.Logger,
	controller *controller.ControllerList,
) {
	v1 := r.Group("/v1")

	// indexController := controller.NewIndexController(log)
	// controller := LoadController(log)
	v1.GET("/ping", controller.Index.Ping)

	v1.POST("/register", controller.Auth.Register)

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
