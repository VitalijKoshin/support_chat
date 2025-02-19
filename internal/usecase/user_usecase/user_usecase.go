package userusecase

import (
	"context"

	"supportchat/internal/apperrors"
	"supportchat/internal/domain/model"
	repository "supportchat/internal/repository/user_repository"
	"supportchat/internal/utils"

	"github.com/google/uuid"
)

type IUserUsecase interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	CheckUserByNickname(ctx context.Context, nickname string) (bool, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUsersByPaginationQuery(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error)
}

type UserUsecase struct {
	UserRepoCache repository.IUserRepository
	UserRepo      repository.IUserRepository
}

func NewUserUsecase(userRepoCache repository.IUserRepository, userRepo repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		UserRepoCache: userRepoCache,
		UserRepo:      userRepo,
	}
}

func (u *UserUsecase) Create(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := u.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return &model.User{}, apperrors.UsecaseCreateUserError.AppendMessage(err.Error())
	}

	return createdUser, nil
}

func (u *UserUsecase) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := u.UserRepoCache.GetUser(ctx, userID)
	if err != nil {
		user, err = u.UserRepo.GetUser(ctx, userID)
		if err != nil {
			return &model.User{}, apperrors.UsecaseGetUserByIDError.AppendMessage(err.Error())
		}
	}

	return user, nil
}

func (u *UserUsecase) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	user, err := u.UserRepoCache.GetUserByNickname(ctx, nickname)
	if err != nil {
		user, err = u.UserRepo.GetUserByNickname(ctx, nickname)
		if err != nil {
			return &model.User{}, apperrors.UsecaseGetUserByNicknameError.AppendMessage(err.Error())
		}
	}

	return user, nil
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	updatedUser, err := u.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return &model.User{}, apperrors.UsecaseUpdateUserError.AppendMessage(err.Error())
	}

	return updatedUser, nil
}

func (u *UserUsecase) CheckUserByNickname(ctx context.Context, nickname string) (bool, error) {
	isExist, err := u.UserRepo.CheckUserByNickname(ctx, nickname)
	if err != nil {
		return false, apperrors.UsecaseCheckUserByNicknameError.AppendMessage(err.Error())
	}

	return isExist, nil
}

func (u *UserUsecase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := u.UserRepo.DeleteUser(ctx, userID)
	if err != nil {
		return apperrors.UsecaseDeleteUserError.AppendMessage(err.Error())
	}

	return nil
}

func (u *UserUsecase) GetUsersByPaginationQuery(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	users, err := u.GetUsers(ctx, paginationQuery)
	if err != nil {
		return &model.Users{}, apperrors.UsecaseGetUsersError.AppendMessage(err.Error())
	}

	return users, nil
}

func (u *UserUsecase) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	users, err := u.UserRepo.GetUsers(ctx, paginationQuery.Page, paginationQuery.Size)
	if err != nil {
		return &model.Users{}, err
	}

	return users, nil
}
