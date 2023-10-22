package main

import (
	"fmt"
	controller "gin-api/cmd/app/controller"
	config "gin-api/cmd/config"
	"gin-api/cmd/helpers"
	"runtime"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// Recovered from panic
			fmt.Print(err)
			fmt.Println("Recovered from panic:", err)

			// Log the stack trace
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			fmt.Printf("Stack Trace:\n%s\n", stack)

			// Handle the panic in your ErrorHandler or another function
			message := fmt.Sprint(err)
			code := config.ResPanicError
			data := helpers.ResponseData{
				Message: &message,
				Code:    &code,
			}
			helpers.ResponseTemplate(c, data)
		}
	}()
	c.Next()
	fmt.Println(c.Errors)

	for _, err := range c.Errors {
		message := err.Error()
		stack := make([]byte, 4096)
		stack = stack[:runtime.Stack(stack, false)]
		fmt.Printf("Error Found:\n%s\n", message)
		fmt.Printf("Stack Trace:\n%s\n", stack)
		code := config.ResServerError
		data := helpers.ResponseData{
			Message: &message,
			Code:    &code,
		}
		helpers.ResponseTemplate(c, data)
		return
	}

}

func NoRoute(c *gin.Context) {
	c.Next()

	message := "Route not found"
	code := config.ResNoRoute
	data := helpers.ResponseData{
		Message: &message,
		Code:    &code,
	}
	helpers.ResponseTemplate(c, data)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(ErrorHandler)

	r.GET("/", controller.Hi)

	// Ping test
	// loading route
	apiRoute := r.Group("/api")
	config.ApiRouteV1(apiRoute)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })
	r.NoRoute(NoRoute)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
