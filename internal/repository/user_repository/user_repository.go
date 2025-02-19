package repository

import (
	"context"

	"supportchat/internal/apperrors"
	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/google/uuid"
)

var UserRepositoryTypes = struct {
	MySQL  string
	Redis  string
	Memory string
}{
	MySQL:  "MySQL",
	Redis:  "Redis",
	Memory: "Memory",
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	GetUsers(ctx context.Context, page int, perPage int) (*model.Users, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	CheckUserByNickname(ctx context.Context, nickname string) (bool, error)
}

func NewUserRepository(userRepositoryType string, db *database.Database) (IUserRepository, error) {
	switch userRepositoryType {
	case UserRepositoryTypes.MySQL:
		return NewMySQLUserRepository(db.Mysql), nil
	case UserRepositoryTypes.Redis:
		return NewRedisUserRepository(db.Redis), nil
	case UserRepositoryTypes.Memory:
		return NewMemoryUserRepository(db.Memory), nil
	default:
		return nil, apperrors.NewUserRepositoryTypeNotSupported.AppendMessage("user repository type: %s", userRepositoryType)
	}
}
