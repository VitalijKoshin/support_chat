package routes

import (
	"supportchat/internal/config"
	"supportchat/internal/controllers"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	UserController *controllers.UserController
	ApiController  *controllers.ApiController
	ChatController *controllers.ChatController
	cfg            *config.Config
}

func NewRoutes(
	e *echo.Echo,
	u *controllers.UserController,
	api *controllers.ApiController,
	chat *controllers.ChatController,
	cfg *config.Config,
) *echo.Echo {
	routes := Routes{
		UserController: u,
		ApiController:  api,
		ChatController: chat,
		cfg:            cfg,
	}

	// allow CORS
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return next(c)
		}
	})

	UserGroup(e, routes.UserController, routes.cfg)
	ApiGroup(e, routes.ApiController, routes.cfg)
	ChatGroup(e, routes.ChatController, routes.cfg)

	e.GET("routes", func(c echo.Context) error {
		return c.JSON(200, e.Routes())
	})
	return e
}
