package sessionrepository

import (
	"context"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/google/uuid"
)

var SessionRepositoryTypes = struct {
	MySQL string
	Redis string
}{
	MySQL: "MySQL",
	Redis: "Redis",
}

type ISessionRepository interface {
	GetUserSessionBySessionID(ctx context.Context, sessionID string) (*model.UserSession, error)
	SetUserSessionBySessionID(ctx context.Context, userSession *model.UserSession) error
	GetUserIDBySessionID(ctx context.Context, sessionID string) (*uuid.UUID, error)
	CreateUserIDBySessionID(ctx context.Context, sessionID string) (*uuid.UUID, error)
	BindingSessionIDWithUserID(ctx context.Context, userSession *model.UserSession) error
	DeleteSessionBySessionID(ctx context.Context, sessionID string) error
	generateKey(sessionID string) string
}

func NewSessionRepository(sessionRepositoryType string, cfg *config.Config, db *database.Database) (ISessionRepository, error) {
	switch sessionRepositoryType {
	case SessionRepositoryTypes.MySQL:
		return NewSessionRepositoryMysql(cfg, db.Mysql), nil
	case SessionRepositoryTypes.Redis:
		return NewSessionRepositoryRedis(cfg, db.Redis), nil
	default:
		return nil, apperrors.NewSessionRepositoryTypeNotSupported.AppendMessage("session repository type: %s", sessionRepositoryType)
	}
}
