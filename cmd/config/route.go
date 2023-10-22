package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func ApiRouteV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

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
