package controllers

import (
	"fmt"
	"net/http"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/domain/model"
	userusecase "supportchat/internal/usecase/user_usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type ChatController struct {
	userUsecase userusecase.IUserUsecase
	cfg         *config.Config
	upgrader    *websocket.Upgrader
	connections map[string]*websocket.Conn
}

func NewChatController(userUsecase userusecase.IUserUsecase, cfg *config.Config) *ChatController {
	return &ChatController{
		userUsecase: userUsecase,
		cfg:         cfg,
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		connections: make(map[string]*websocket.Conn),
	}
}

func (c *ChatController) ChatSupport(ctx echo.Context) error {
	conn, err := c.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		appError := apperrors.ChatControllerChatSupportUpgrade.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			appError := apperrors.ChatControllerChatSupportReadMessage.AppendMessage(err)
			return ctx.JSON(appError.HTTPCode, appError.JsonError())
		}
		fmt.Println(string(p), messageType)
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			appError := apperrors.ChatControllerChatSupportWriteMessage.AppendMessage(err)
			return ctx.JSON(appError.HTTPCode, appError.JsonError())
		}
	}
}

func (c *ChatController) SetUpJWTConfig() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(c.cfg.JWT.Secret),
	}

	return echojwt.WithConfig(config)
}
