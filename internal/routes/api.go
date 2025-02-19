package routes

import (
	"net/http"

	"supportchat/internal/config"
	"supportchat/internal/controllers"

	"github.com/labstack/echo/v4"
)

func ApiGroup(e *echo.Echo, apiController *controllers.ApiController, cfg *config.Config) {
	apiAuthGroup := e.Group("/auth/api")
	apiAuthGroup.POST("/login", apiController.Login)
	apiAuthGroup.POST("/signup", apiController.Signup)

	apiRestrictedGroup := e.Group("/api")
	apiRestrictedGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return next(c)
		}
	})
	apiRestrictedGroup.OPTIONS("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	apiRestrictedGroup.Use(apiController.SetUpJWTConfig())
	apiRestrictedGroup.GET("/user", apiController.User)
	apiRestrictedGroup.POST("/check-token", apiController.CheckToken)
}
