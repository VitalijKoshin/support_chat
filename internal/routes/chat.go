package routes

import (
	"supportchat/internal/config"
	"supportchat/internal/controllers"

	"github.com/labstack/echo/v4"
)

func ChatGroup(e *echo.Echo, chatController *controllers.ChatController, cfg *config.Config) {
	// Create a new group for chat routes with websocket support
	chatGroup := e.Group("/chat-support")
	chatGroup.Use(chatController.SetUpJWTConfig())
	chatGroup.GET("/ws", chatController.ChatSupport)
}
