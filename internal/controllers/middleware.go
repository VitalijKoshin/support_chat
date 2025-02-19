package controllers

import (
	"net/http"
	"strings"
	"time"

	"supportchat/internal/apperrors"
	"supportchat/internal/domain/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	UserAuthCtx      = "userAuth"
	updatePermission = "update"
	deletePermission = "delete"
)

func (u *UserController) SetUpJWTConfig() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(u.cfg.JWT.Secret),
	}

	return echojwt.WithConfig(config)
}

func (u *UserController) CanUpdateUser() echo.MiddlewareFunc {
	return u.hasPermission(updatePermission)
}

func (u *UserController) CanDeleteUser() echo.MiddlewareFunc {
	return u.hasPermission(deletePermission)
}

func (u *UserController) hasPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userUUID, err := uuid.Parse(ctx.Param("id"))
			if err != nil {
				appError := apperrors.HasPermissionUuidParse.AppendMessage(err)
				return ctx.JSON(appError.HTTPCode, appError.Error())
			}

			authUser, err := u.FetchJWTUser(ctx)
			if err != nil {
				appError := apperrors.HasPermissionFetchJwtUser.AppendMessage(err)
				return ctx.JSON(appError.HTTPCode, appError.Error())
			}
			if userUUID == authUser.UserID {
				return next(ctx)
			}

			err = authUser.Can(permission)
			if err != nil {
				appError := err.(*apperrors.AppError)
				return ctx.JSON(appError.HTTPCode, appError.Error())
			}

			return next(ctx)
		}
	}
}

func (u *UserController) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(*model.JwtCustomClaims)

		if !user.Valid {
			apperrors.MiddlewareJWTAuthValid.AppendMessage(echo.ErrUnauthorized)
		}

		_, err := u.VerifyJwtUser(ctx, claims.Nickname, claims.Role)
		if err != nil {
			apperrors.MiddlewareJWTAuthVerifyJwtUser.AppendMessage(echo.ErrUnauthorized)
		}

		return next(ctx)
	}
}

func (u *UserController) BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(u.VerifyAuthUser())
}

func (u *UserController) VerifyJwtUser(ctx echo.Context, nickname, role string) (bool, error) {
	user, err := u.userUsecase.GetUserByNickname(ctx.Request().Context(), nickname)
	if err != nil {
		return false, apperrors.MiddlewareVerifyJwtUserGetUserByNickname.AppendMessage(err)
	}
	if user == nil {
		return false, apperrors.MiddlewareVerifyJwtUserGetUserByNickname.AppendMessage(err)
	}
	if user.Role != role {
		return false, apperrors.MiddlewareVerifyAuthUserGetUserByNickname.AppendMessage(err)
	}
	ctx.Set(UserAuthCtx, user)

	return true, nil
}

func (u *UserController) VerifyAuthUser() func(username, password string, ctx echo.Context) (bool, error) {
	return func(username, password string, ctx echo.Context) (bool, error) {
		user, err := u.userUsecase.GetUserByNickname(ctx.Request().Context(), username)
		if err != nil {
			return false, apperrors.MiddlewareVerifyAuthUserGetUserByNickname.AppendMessage(err)
		}
		if user == nil {
			return false, apperrors.MiddlewareVerifyAuthUserGetUserByNickname.AppendMessage(err)
		}

		err = user.ComparePasswords(password)
		if err != nil {
			return false, apperrors.MiddlewareVerifyAuthUserComparePasswords.AppendMessage(err)
		}

		ctx.Set(UserAuthCtx, user)

		return true, nil
	}
}

func JWTMiddleware(next echo.HandlerFunc, secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			err := apperrors.MiddlewareJWTMiddlewareMissingToken.AppendMessage(echo.ErrUnauthorized)
			return c.JSON(http.StatusUnauthorized, err.JsonError())
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			appError := apperrors.MiddlewareJWTAuthValid.AppendMessage(err)
			return c.JSON(http.StatusUnauthorized, appError.JsonError())
		}

		// Store token in context
		c.Set("user", token)

		return next(c)
	}
}

func GenerateJWT(userID uuid.UUID, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// func SetUpJWTConfig(userID uuid.UUID, secret string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"user_id": userID,
// 		"exp":     time.Now().Add(time.Hour * 24).Unix(),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secret))
// }

// func SetUpJWTConfig() echo.MiddlewareFunc {
// 	config := echojwt.Config{
// 		NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 			return new(model.JwtCustomClaims)
// 		},
// 		SigningKey: []byte(uc.cfg.Jwt.Secret),
// 	}

// 	return echojwt.WithConfig(config)
// }
