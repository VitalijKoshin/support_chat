package repository

import (
	"context"
	"database/sql"
	"errors"

	"supportchat/internal/apperrors"
	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/google/uuid"
)

type UserRepositoryMysql struct {
	db *database.MySQLDB
}

func NewMySQLUserRepository(db *database.MySQLDB) IUserRepository {
	return &UserRepositoryMysql{
		db: db,
	}
}

func (r *UserRepositoryMysql) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	sqlStr := `INSERT INTO user (user_id, nickname, first_name, last_name, email, password, is_public, user_role, ip, user_agent, country, created_at, updated_at, deleted_at, login_date, created_by)
    			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, sqlStr, user.UserID, user.Nickname, user.FirstName, user.LastName, user.Email, user.Password, user.IsPublic, user.Role, user.Ip, user.UserAgent, user.Country, user.Created.At, user.UpdatedAt, user.DeletedAt, user.LoginDate, user.Created.By)
	if err != nil {
		return nil, apperrors.UserRepositoryMysqlCreateUserError.AppendMessage(err)
	}

	return user, nil
}

func (r *UserRepositoryMysql) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryMysql) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	userRow := r.db.QueryRowContext(ctx, "SELECT `user_id`, `nickname`, `email`, `password`, `created_at`, `updated_at` FROM `user` WHERE `user_id` = ?", userID)
	if userRow.Err() != nil {
		if errors.Is(userRow.Err(), sql.ErrNoRows) {
			return nil, apperrors.UserRepositoryMysqlGetUserNotFoundError.AppendMessage(userRow.Err())
		}
		return nil, apperrors.UserRepositoryMysqlGetUserError.AppendMessage(userRow.Err())
	}

	user := model.User{}
	scanError := userRow.Scan(&user.UserID, &user.Nickname, &user.Email, &user.Password, &user.Created.At, &user.UpdatedAt)
	if scanError != nil {
		if errors.Is(scanError, sql.ErrNoRows) {
			return nil, apperrors.UserRepositoryMysqlGetUserScanEmpty.AppendMessage(scanError)
		}
		return nil, apperrors.UserRepositoryMysqlGetUserError.AppendMessage(scanError)
	}

	return &user, nil
}

func (r *UserRepositoryMysql) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryMysql) GetUsers(ctx context.Context, page int, perPage int) (*model.Users, error) {
	return nil, nil
}

func (r *UserRepositoryMysql) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM `user` WHERE `user_id` = ?", userID)
	if err != nil {
		return apperrors.UserRepositoryMysqlDeleteUserError.AppendMessage(err)
	}

	return nil
}

func (r *UserRepositoryMysql) CheckUserByNickname(ctx context.Context, nickname string) (bool, error) {
	return false, nil
}
