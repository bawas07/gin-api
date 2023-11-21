package main

import (
	"fmt"
	"gin-api/cmd/app/controller"
	config "gin-api/cmd/config"
	"gin-api/cmd/helpers/loggers"
	"gin-api/cmd/helpers/responses"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func ErrorHandler(c *gin.Context) {
	eLog := loggers.GetError()
	defer func() {
		if err := recover(); err != nil {
			// Recovered from panic
			fmt.Print(err)
			fmt.Println("Recovered from panic:", err)
			eLog.Panic().Stack().Interface("Panic Attack!", err).Msg("Panic Detected!")

			// Log the stack trace
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			fmt.Printf("Stack Trace:\n%s\n", stack)

			// Handle the panic in your ErrorHandler or another function
			message := fmt.Sprint(err)
			code := responses.ResPanicError
			data := responses.ResponseData{
				Message: &message,
				Code:    &code,
			}
			responses.ResponseTemplate(c, data)
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
		eLog.Error().Stack().Err(err).Msg("Unhandled Error Detected!")

		code := responses.ResServerError
		data := responses.ResponseData{
			Message: &message,
			Code:    &code,
		}
		responses.ResponseTemplate(c, data)
		return
	}

}

func NoRoute(c *gin.Context) {
	c.Next()

	message := "Route not found"
	code := responses.ResNoRoute
	data := responses.ResponseData{
		Message: &message,
		Code:    &code,
	}
	responses.ResponseTemplate(c, data)
}

func setupRouter(log *zerolog.Logger, controller *controller.ControllerList) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(loggers.RequestLogger(log))
	r.Use(ErrorHandler)

	r.GET("/", controller.Index.Hi)

	// Ping test
	// loading route
	apiRoute := r.Group("/api")
	config.ApiRouteV1(apiRoute, log, controller)

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

func init() {
	l := loggers.Get()

	l.Info().Msg("App initializing")
}

func main() {
	startTime := time.Now()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Setup Logger
	startModuleTime := time.Now()
	l := loggers.Get()
	l.Info().Msgf("Logger initialized: %s", time.Since(startModuleTime))

	// Load Service
	startModuleTime = time.Now()
	s := config.LoadService(&l)
	l.Info().Msgf("Service initialized: %s", time.Since(startModuleTime))

	// Load Controller
	startModuleTime = time.Now()
	c := config.LoadController(&l, &s)
	l.Info().Msgf("Controller initialized: %s", time.Since(startModuleTime))

	// Load Router
	startModuleTime = time.Now()
	r := setupRouter(&l, &c)
	l.Info().Msgf("Router initialized: %s", time.Since(startModuleTime))
	l.Info().Msgf("Total Startup Time: %s", time.Since(startTime))
	// Listen and Server in 0.0.0.0:8080
	l.Info().
		Dur("total_startup_time", time.Since(startTime)).
		Str("port", port).
		Msgf("Starting App Server on port '%s'", port)
	r.Run(":" + port)
}
