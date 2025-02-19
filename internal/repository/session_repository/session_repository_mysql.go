package sessionrepository

import (
	"context"
	"database/sql"
	"errors"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/google/uuid"
)

const (
	sessionPrefix = ""
)

type SessionRepositoryMysql struct {
	cfg *config.Config
	sql *database.MySQLDB
}

func NewSessionRepositoryMysql(cfg *config.Config, sql *database.MySQLDB) ISessionRepository {
	return &SessionRepositoryMysql{
		cfg: cfg,
		sql: sql,
	}
}

func (srm *SessionRepositoryMysql) GetUserSessionBySessionID(ctx context.Context, sessionID string) (*model.UserSession, error) {
	userSessionRow := srm.sql.QueryRowContext(ctx, "SELECT `user_session_id`, `user_id`, `session_id` FROM `user_session` WHERE `session_id` = ?", sessionID)
	if userSessionRow.Err() != nil {
		if errors.Is(userSessionRow.Err(), sql.ErrNoRows) {
			return nil, apperrors.SessionRepositoryMysqlGetUserSessionBySessionNotFoundError.AppendMessage(userSessionRow.Err())
		}
		return nil, apperrors.SessionRepositoryMysqlGetUserSessionBySessionGetError.AppendMessage(userSessionRow.Err())
	}

	userSession := model.UserSession{}
	scanError := userSessionRow.Scan(&userSession.UserSessionID, &userSession.UserID, &userSession.SessionID)
	if scanError != nil {
		if errors.Is(scanError, sql.ErrNoRows) {
			return nil, apperrors.SessionRepositoryMysqlGetUserSessionBySessionScanEmpty.AppendMessage(scanError)
		}
		return nil, apperrors.SessionRepositoryMysqlGetUserSessionBySessionGetError.AppendMessage(scanError)
	}

	return &userSession, nil
}

func (srm *SessionRepositoryMysql) SetUserSessionBySessionID(ctx context.Context, userSession *model.UserSession) error {
	_, err := srm.sql.ExecContext(ctx, "INSERT INTO `user_session` (`user_session_id`, `user_id`, `session_id`) VALUES (?, ?, ?)", userSession.UserSessionID, userSession.UserID, userSession.SessionID)
	if err != nil {
		return apperrors.SessionRepositoryMysqlSetUserSessionBySessionSetError.AppendMessage(err)
	}

	return nil
}

func (srm *SessionRepositoryMysql) GetUserIDBySessionID(ctx context.Context, sessionID string) (*uuid.UUID, error) {
	sid := srm.generateKey(sessionID)
	userIdRow := srm.sql.QueryRowContext(ctx, "SELECT `user_id` FROM `user_session` WHERE `session_id` = ?", sid)
	if userIdRow.Err() != nil {
		if errors.Is(userIdRow.Err(), sql.ErrNoRows) {
			return nil, apperrors.SessionRepositoryMysqlGetUserIDBySessionGetEmpty.AppendMessage(userIdRow.Err())
		}
		return nil, apperrors.SessionRepositoryMysqlGetUserIDBySessionGetError.AppendMessage(userIdRow.Err())
	}

	var userID string
	scanError := userIdRow.Scan(&userID)
	if scanError != nil {
		if errors.Is(scanError, sql.ErrNoRows) {
			return nil, apperrors.SessionRepositoryMysqlGetUserIDBySessionScanEmpty.AppendMessage(scanError)
		}
		return nil, apperrors.SessionRepositoryMysqlGetUserIDBySessionGetError.AppendMessage(scanError)
	}

	// convert string to uuid
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, apperrors.SessionRepositoryMysqlGetUserIDBySessionParseError.AppendMessage(err)
	}

	return &uid, nil
}

func (srm *SessionRepositoryMysql) BindingSessionIDWithUserID(ctx context.Context, userSession *model.UserSession) error {
	_, err := srm.sql.ExecContext(ctx, "INSERT INTO `user_session` (`user_session_id`, `user_id`, `session_id`) VALUES (?, ?, ?)", userSession.UserSessionID, userSession.UserID, userSession.SessionID)
	if err != nil {
		return apperrors.SessionRepositoryMysqlBindingSessionIDWithUserIDSetError.AppendMessage(err)
	}

	return nil
}

func (srm *SessionRepositoryMysql) CreateUserIDBySessionID(ctx context.Context, sessionID string) (*uuid.UUID, error) {
	sid := srm.generateKey(sessionID)
	userID := uuid.New()
	// insert user id and session id into user_session table
	_, err := srm.sql.ExecContext(ctx, "INSERT INTO `user_session` (`user_id`, `session_id`) VALUES (?, ?)", userID, sid)
	if err != nil {
		return nil, apperrors.SessionRepositoryMysqlCreateUserIDBySessionSetError.AppendMessage(err)
	}

	return &userID, nil
}

func (srm *SessionRepositoryMysql) DeleteSessionBySessionID(ctx context.Context, sessionID string) error {
	sid := srm.generateKey(sessionID)
	_, err := srm.sql.ExecContext(ctx, "DELETE FROM `user_session` WHERE `session_id` = ?", sid)
	if err != nil {
		return apperrors.SessionRepositoryMysqlDeleteSessionBySessionIDError.AppendMessage(err)
	}

	return nil
}

func (srm *SessionRepositoryMysql) generateKey(sessionID string) string {
	return sessionPrefix + sessionID
}
