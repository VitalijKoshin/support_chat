package routes

import (
	"supportchat/internal/config"
	"supportchat/internal/controllers"

	"github.com/labstack/echo/v4"
)

func UserGroup(e *echo.Echo, userController *controllers.UserController, cfg *config.Config) {
	userGroup := e.Group("/users")
	userGroup.GET("/:id", userController.GetUserByID)
	userGroup.POST("", userController.CreateUser)
}
