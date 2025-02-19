package repository

import (
	"context"

	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/google/uuid"
)

type UserRepositoryRedis struct {
	db *database.RedisDB
}

func NewRedisUserRepository(db *database.RedisDB) IUserRepository {
	return &UserRepositoryRedis{
		db: db,
	}
}

func (r *UserRepositoryRedis) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryRedis) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryRedis) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryRedis) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryRedis) GetUsers(ctx context.Context, page int, perPage int) (*model.Users, error) {
	return nil, nil
}

func (r *UserRepositoryRedis) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (r *UserRepositoryRedis) CheckUserByNickname(ctx context.Context, nickname string) (bool, error) {
	return false, nil
}
