package repository

import (
	"context"

	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/google/uuid"
)

type UserRepositoryMemory struct {
	db *database.MemoryDB
}

func NewMemoryUserRepository(db *database.MemoryDB) IUserRepository {
	return &UserRepositoryMemory{
		db: db,
	}
}

func (r *UserRepositoryMemory) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryMemory) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryMemory) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryMemory) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryMemory) GetUsers(ctx context.Context, page int, perPage int) (*model.Users, error) {
	return nil, nil
}

func (r *UserRepositoryMemory) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (r *UserRepositoryMemory) CheckUserByNickname(ctx context.Context, nickname string) (bool, error) {
	return false, nil
}
