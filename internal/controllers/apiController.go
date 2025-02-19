package controllers

import (
	"fmt"
	"net/http"
	"time"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/domain/model"
	userusecase "supportchat/internal/usecase/user_usecase"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type ApiController struct {
	authUsecase userusecase.IAuthUsecase
	userUsecase userusecase.IUserUsecase
	cfg         *config.Config
}

func NewApiController(userUsecase userusecase.IUserUsecase, authUsecase userusecase.IAuthUsecase, cfg *config.Config) *ApiController {
	return &ApiController{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
		cfg:         cfg,
	}
}

func (a *ApiController) Login(ctx echo.Context) error {
	loginRequest := new(model.LoginRequest)
	if err := ctx.Bind(loginRequest); err != nil {
		appError := apperrors.ApiControllerLoginBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	userSession, err := a.authUsecase.Login(ctx.Request().Context(), loginRequest)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	claims := &model.JwtCustomClaims{
		UserID: userSession.UserID,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(a.cfg.JWT.Ttl)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(a.cfg.JWT.Secret))
	if err != nil {
		appErr := err.(*apperrors.AppError)
		return ctx.JSON(appErr.HTTPCode, appErr.Error())
	}

	// valiate token
	tokenParsed, errParsed := jwt.ParseWithClaims(tokenSigned, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.JWT.Secret), nil
	})
	if errParsed != nil {
		fmt.Println("Error parsing token: ", errParsed)
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	claimsParsed := tokenParsed.Claims.(*model.JwtCustomClaims)
	if !tokenParsed.Valid {
		fmt.Println("Token is not valid")
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	fmt.Println("Token is valid")
	fmt.Println("Claims: ", claimsParsed)

	return ctx.JSON(http.StatusOK, model.LoginResponse{Token: tokenSigned})
}

func (a *ApiController) Signup(ctx echo.Context) error {
	signupRequest := &model.SignupRequest{}
	if err := ctx.Bind(signupRequest); err != nil {
		appError := apperrors.ApiControllerSignupBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	signupRequest.Ip = ctx.RealIP()
	signupRequest.Browser = ctx.Request().UserAgent()
	signupRequest.Country = ctx.Request().Header.Get("CF-IPCountry")

	user, err := a.authUsecase.Signup(ctx.Request().Context(), signupRequest)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	return ctx.JSON(http.StatusOK, user.SignUpResponse())
}

func (a *ApiController) User(ctx echo.Context) error {
	authUser, err := a.FetchJWTUser(ctx)
	if err != nil {
		return ctx.JSON(err.(*apperrors.AppError).HTTPCode, err)
	}
	return ctx.JSON(http.StatusOK, authUser.MapUserModelToGetUserResponse())
}

func (a *ApiController) CheckToken(ctx echo.Context) error {
	user, err := a.FetchJWTUser(ctx)
	if err != nil {
		return ctx.JSON(err.(*apperrors.AppError).HTTPCode,
			err.(*apperrors.AppError).JsonError())
	}
	return ctx.JSON(http.StatusOK, model.CheckTokenResponse{
		Message: "Token is valid",
		UserID:  user.UserID,
	})
}

func (a *ApiController) FetchJWTUser(ctx echo.Context) (*model.User, error) {
	userContext, ok := ctx.Get("user").(*jwt.Token)
	if !ok || userContext == nil {
		appError := apperrors.UserControllerFetchJWTUser.AppendMessage(echo.ErrUnauthorized)
		return nil, appError
	}
	claims := userContext.Claims.(*model.JwtCustomClaims)

	if !userContext.Valid {
		appError := apperrors.UserControllerFetchJWTUserUserContextNotValid.AppendMessage(echo.ErrUnauthorized)
		return nil, appError
	}

	return &model.User{
		UserID:   claims.UserID,
		Nickname: claims.Nickname,
		Role:     claims.Role,
	}, nil
}

func (a *ApiController) SetUpJWTConfig() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(a.cfg.JWT.Secret),
	}

	return echojwt.WithConfig(config)
}
