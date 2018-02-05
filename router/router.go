package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/controller"
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
	)

	api := app.Group("/api")
	{
		api.GET("/version", indexController.GetVersion)
	}
}
