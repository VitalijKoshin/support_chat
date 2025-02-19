package userusecase

import (
	"context"
	"fmt"
	"net/http"

	"supportchat/internal/apperrors"
	"supportchat/internal/domain/model"
	sessionrepository "supportchat/internal/repository/session_repository"
	userrepository "supportchat/internal/repository/user_repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IAuthUsecase interface {
	Login(ctx context.Context, loginRequest *model.LoginRequest) (*model.UserSession, error)
	Signup(ctx context.Context, signupRequest *model.SignupRequest) (*model.User, error)
	getUserBySessionID(ctx context.Context, sessionID string) (*model.User, error)
	getUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
}

type AuthUsecase struct {
	UserRepo         userrepository.IUserRepository
	UserRepoCache    userrepository.IUserRepository
	SessionRepo      sessionrepository.ISessionRepository
	SessionRepoCache sessionrepository.ISessionRepository
}

func NewAuthUsecase(
	userRepo userrepository.IUserRepository,
	userRepoCache userrepository.IUserRepository,
	sessionRepo sessionrepository.ISessionRepository,
	sessionRepoCache sessionrepository.ISessionRepository,
) IAuthUsecase {
	return &AuthUsecase{
		UserRepo:         userRepo,
		UserRepoCache:    userRepoCache,
		SessionRepo:      sessionRepo,
		SessionRepoCache: sessionRepoCache,
	}
}

func (a *AuthUsecase) Login(ctx context.Context, loginRequest *model.LoginRequest) (*model.UserSession, error) {
	userSession, err := a.geUserSessionBySessionId(ctx, loginRequest.SessionID)
	if err != nil {
		return nil, apperrors.AuthUsecaseLoginGeUserSessionBySessionId.AppendMessage(err)
	}

	if userSession == nil {
		appError := apperrors.AuthUsecaseLoginUserSessionNotFound.AppendMessage(echo.ErrUnauthorized)
		return nil, appError
	}

	return userSession, nil
}

func (a *AuthUsecase) Signup(ctx context.Context, signupRequest *model.SignupRequest) (*model.User, error) {
	// get user session id by session id
	userSession, err := a.geUserSessionBySessionId(ctx, signupRequest.SessionID)
	if err != nil {
		return nil, apperrors.AuthUsecaseSignupGeUserSessionBySessionId.AppendMessage(err)
	}

	userModel := &model.User{}
	// if user session is null return create user session
	if userSession == nil {
		fmt.Println("User session is null")
		userSession = model.NewUserSession(uuid.New(), signupRequest.SessionID)
		err := a.setUserSession(ctx, userSession)
		if err != nil {
			return nil, apperrors.AuthUsecaseSignupSetUserSession.AppendMessage(err)
		}

		userModel = model.NewUserFromUserSessionSignupRequest(userSession, signupRequest)
		fmt.Println("User session is null userModel: ", userModel)
		// save user to db
		_, err = a.UserRepo.CreateUser(ctx, userModel)
		if err != nil {
			return nil, apperrors.AuthUsecaseSignupCreateUser.AppendMessage(err)
		}
	} else {
		fmt.Println("User session is not null")
		// if user session is not null return user by session id
		userModel, err = a.getUserByID(ctx, userSession.UserID)
		if err != nil {
			return nil, apperrors.AuthUsecaseSignupGetUserByID.AppendMessage(err)
		}
	}

	return userModel, nil
}

func (a *AuthUsecase) setUserSession(ctx context.Context, userSession *model.UserSession) error {
	err := a.SessionRepo.SetUserSessionBySessionID(ctx, userSession)
	if err != nil {
		return apperrors.AuthUsecaseSessionRepoSetUserSession.AppendMessage(err)
	}

	err = a.SessionRepoCache.SetUserSessionBySessionID(ctx, userSession)
	if err != nil {
		return apperrors.AuthUsecaseSessionRepoCacheSetUserSession.AppendMessage(err)
	}

	return nil
}

func (a *AuthUsecase) geUserSessionBySessionId(ctx context.Context, sessionID string) (*model.UserSession, error) {
	userSession, err := a.SessionRepoCache.GetUserSessionBySessionID(ctx, sessionID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.HTTPCode != http.StatusNotFound {
			return nil, apperrors.AuthUsecaseSessionRepoCacheGetUserSessionBySessionID.AppendMessage(err)
		}
	}

	if userSession != nil {
		return userSession, nil
	}

	userSession, err = a.SessionRepo.GetUserSessionBySessionID(ctx, sessionID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.HTTPCode != http.StatusNotFound {
			return nil, apperrors.AuthUsecaseSessionRepoGetUserSessionBySessionID.AppendMessage(err)
		}
	}

	return userSession, nil
}

func (a *AuthUsecase) getUserBySessionID(ctx context.Context, sessionID string) (*model.User, error) {
	userModel := &model.User{}
	// get user id from cache by session id
	userId, err := a.SessionRepoCache.GetUserIDBySessionID(ctx, sessionID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.SessionRepositoryRedisGetUserIDBySessionNotFoundError.Code {
			return nil, apperrors.AuthUsecaseGetUserBySessionID.AppendMessage(err)
		}
	}

	// if user id is null check a.SessionRepo and return user model
	if userId == nil {
		userId, err = a.SessionRepo.GetUserIDBySessionID(ctx, sessionID)
		if err != nil {
			appError := err.(*apperrors.AppError)
			if appError.HTTPCode != http.StatusNotFound {
				return nil, apperrors.AuthUsecaseGetUserBySessionID.AppendMessage(err)
			}
		}
	}

	if userId == nil {
		return nil, nil
	}

	userModel, err = a.getUserByID(ctx, *userId)
	if err != nil {
		return nil, apperrors.AuthUsecaseGetUserBySessionID.AppendMessage(err)
	}

	return userModel, nil
}

func (a *AuthUsecase) getUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	userModel, err := a.UserRepoCache.GetUser(ctx, userID)
	if err != nil {
		return nil, apperrors.UsecaseGetUserByIDError.AppendMessage(err)
	}

	if userModel != nil {
		return userModel, nil
	}

	userModel, err = a.UserRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, apperrors.UsecaseGetUserByIDError.AppendMessage(err)
	}

	return userModel, nil
}
