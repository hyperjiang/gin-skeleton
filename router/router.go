package router

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/controller"
	"github.com/hyperjiang/gin-skeleton/middleware"
)

// Route makes the routing
func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	app.GET(
		"/", indexController.GetIndex,
	)

	userController := new(controller.UserController)
	app.GET(
		"/user/:id", userController.GetUser,
	).GET(
		"/signup", userController.SignupForm,
	).POST(
		"/signup", userController.Signup,
	).GET(
		"/login", userController.LoginForm,
	).POST(
		"/login", userController.Login,
	)

	api := app.Group("/api")
	{
		api.GET("/version", indexController.GetVersion)
	}

	auth := app.Group("/auth")
	authMiddleware := middleware.Auth()
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			c.JSON(200, gin.H{
				"user": claims["id"],
				"text": "Hello World.",
			})
		})
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
}
