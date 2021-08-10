package router

import (
	"Api/pkg/handler"
	_ "Api/pkg/router/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger API
// @version 1.0
// @description This is an echo swagger API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New() *echo.Echo {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/comments", handler.ReturnAllComments)
	e.GET("/comments/:id", handler.ReturnComment)
	e.DELETE("/comments/:id", handler.DeleteComment)
	e.PUT("/comments/:id", handler.UpdateComment)
	e.GET("/posts", handler.ReturnAllPosts)
	e.GET("/posts/:id", handler.ReturnPost)
	e.DELETE("/posts/:id", handler.DeletePost)
	e.PUT("/posts/:id", handler.UpdatePost)
	e.GET("/signin/google", handler.GoogleSignIn)
	e.GET("/signin/facebook", handler.FacebookSignIn)
	e.GET("/token/:source", handler.GetAuthToken)
	e.POST("/signup", handler.SignUp)
	e.POST("/signin", handler.SignIn)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("/comments", handler.CreateComment)
	r.POST("/posts", handler.CreatePost)

	return e
}
