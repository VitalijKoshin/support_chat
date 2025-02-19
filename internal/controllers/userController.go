package controllers

import (
	"net/http"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/domain/model"
	userusecase "supportchat/internal/usecase/user_usecase"
	"supportchat/internal/utils"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase userusecase.IUserUsecase
	cfg         *config.Config
}

func NewUserController(userUsecase userusecase.IUserUsecase, cfg *config.Config) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		cfg:         cfg,
	}
}

func (u *UserController) GetUserByID(ctx echo.Context) error {
	userUUID := ctx.Param("id")

	uid, err := uuid.Parse(userUUID)
	if err != nil {
		appError := apperrors.UserControllerGetUserUuidParse.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	user, err := u.userUsecase.GetUserByID(ctx.Request().Context(), uid)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	if user == nil {
		appError := apperrors.UserControllerGetUserUserNotExist
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	return ctx.JSON(http.StatusOK, user.MapUserModelToGetUserResponse())
}

func (u *UserController) CreateUser(ctx echo.Context) error {
	createUserRequest := new(model.CreateUserRequest)
	if err := ctx.Bind(createUserRequest); err != nil {
		appError := apperrors.UserControllerCreateUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	user := &model.User{}
	authUser, err := u.FetchJWTUser(ctx)
	if err != nil {
		return ctx.JSON(err.(*apperrors.AppError).HTTPCode, err)
	}
	user.MapCreateUserRequestToUserModel(createUserRequest)
	user.Created.By = authUser.Nickname

	createdUser, err := u.userUsecase.Create(ctx.Request().Context(), user)
	if err != nil {
		appError := apperrors.UserControllerCreateUserError.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	return ctx.JSON(http.StatusCreated, createdUser.MapUserModelToCreateUserResponse())
}

func (u *UserController) UpdateUser(ctx echo.Context) error {
	updateUser := &model.UpdateUserRequest{}
	userUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		appError := apperrors.UserControllerUpdateUserUuidParse.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	user, err := u.userUsecase.GetUserByID(ctx.Request().Context(), userUUID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	if user == nil {
		appError := apperrors.UserControllerUpdateUserUserNotExist
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	if err = ctx.Bind(updateUser); err != nil {
		appError := apperrors.UserControllerUpdateUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}
	user.MapUpdateUserRequestToUserModel(updateUser)

	authUser, err := u.FetchJWTUser(ctx)
	if err != nil {
		return ctx.JSON(err.(*apperrors.AppError).HTTPCode, err)
	}
	if user.IsAdmin() && !authUser.IsAdmin() {
		appError := apperrors.UserControllerTryToSetAdmin.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	err = user.ComparePasswords(user.Password)
	if err != nil {
		err = user.HashPassword()
	}
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	err = ctx.Validate(user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	_, err = u.userUsecase.CheckUserByNickname(ctx.Request().Context(), user.Nickname)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	updatedUser, err := u.userUsecase.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	return ctx.JSON(http.StatusOK, updatedUser.MapUserModelToUpdateUserResponse())
}

func (u *UserController) DeleteUser(ctx echo.Context) error {
	userID := ctx.Param("id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		appError := apperrors.UserControllerDeleteUserUuidParse.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	err = u.userUsecase.DeleteUser(ctx.Request().Context(), userUUID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.JsonError())
	}

	return ctx.JSON(http.StatusOK, userUUID)
}

func (u *UserController) GetUsers(ctx echo.Context) error {
	paginationQuery, err := utils.GetPaginationFromCtx(ctx.QueryParam("page"), ctx.QueryParam("size"), ctx.QueryParam("orderBy"))
	if err != nil {
		appError := apperrors.UserControllerGetUsersGetPaginationFromCtx.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	users, err := u.userUsecase.GetUsersByPaginationQuery(ctx.Request().Context(), paginationQuery)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, users.MapUserModelToGetUserResponse())
}

func (u *UserController) FetchAuthUser(ctx echo.Context, UserAuthCtx string) *model.User {
	return ctx.Get(UserAuthCtx).(*model.User)
}

func (u *UserController) FetchJWTUser(ctx echo.Context) (*model.User, error) {
	userContext, ok := ctx.Get("user").(*jwt.Token)
	if !ok || userContext == nil {
		appError := apperrors.UserControllerFetchJWTUser.AppendMessage(echo.ErrUnauthorized)
		return nil, appError
	}
	claims := userContext.Claims.(*model.JwtCustomClaims)

	if !userContext.Valid {
		appError := apperrors.UserControllerFetchJWTUser.AppendMessage(echo.ErrUnauthorized)
		return nil, appError
	}

	return &model.User{
		UserID:   claims.UserID,
		Nickname: claims.Nickname,
		Role:     claims.Role,
	}, nil
}
