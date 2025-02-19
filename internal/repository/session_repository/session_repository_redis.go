package sessionrepository

import (
	"context"
	"encoding/json"
	"time"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"
	"supportchat/internal/domain/model"
	"supportchat/internal/infrastructure/database"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	userPrefix = "user_session:"
	sessPrefix = "session_user:"
	userTtl    = 100 * time.Minute
)

type SessionRepositoryRedis struct {
	cfg   *config.Config
	redis *database.RedisDB
}

func NewSessionRepositoryRedis(cfg *config.Config, redis *database.RedisDB) ISessionRepository {
	sessionRepositoryRedis := &SessionRepositoryRedis{
		cfg:   cfg,
		redis: redis,
	}
	return sessionRepositoryRedis
}

func (sr *SessionRepositoryRedis) GetUserSessionBySessionID(ctx context.Context, sessionID string) (*model.UserSession, error) {
	key := sr.generateKey(sessionID)
	userSessionBytes, err := sr.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.SessionRepositoryRedisGetUserSessionBySessionNotFoundError.AppendMessage(err)
		}

		return nil, apperrors.SessionRepositoryRedisGetUserSessionBySessionGetError.AppendMessage(err)
	}

	userSession := model.UserSession{}
	err = json.Unmarshal(userSessionBytes, &userSession)
	if err != nil {
		return nil, apperrors.SessionRepositoryRedisGetUserSessionBySessionParseError.AppendMessage(err)
	}

	return &userSession, nil
}

func (sr *SessionRepositoryRedis) SetUserSessionBySessionID(ctx context.Context, userSession *model.UserSession) error {
	key := sr.generateKey(userSession.SessionID)
	userSessionBytes, err := json.Marshal(userSession)
	if err != nil {
		return apperrors.SessionRepositoryRedisSetUserSessionBySessionMarshalError.AppendMessage(err)
	}

	if err := sr.redis.RedisClient.Set(ctx, key, userSessionBytes, userTtl).Err(); err != nil {
		return apperrors.SessionRepositoryRedisSetUserSessionBySessionSetError.AppendMessage(err)
	}

	return nil
}

func (sr *SessionRepositoryRedis) GetUserIDBySessionID(ctx context.Context, sessionID string) (*uuid.UUID, error) {
	key := sr.generateKey(sessionID)
	userIDBytes, err := sr.redis.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, apperrors.SessionRepositoryRedisGetUserIDBySessionNotFoundError.AppendMessage(err)
		}
		return nil, apperrors.SessionRepositoryRedisGetUserIDBySessionGetError.AppendMessage(err)
	}

	userID := uuid.UUID{}
	err = json.Unmarshal(userIDBytes, &userID)
	if err != nil {
		return nil, apperrors.SessionRepositoryRedisGetUserIDBySessionParseError.AppendMessage(err)
	}

	return &userID, nil
}

func (sr *SessionRepositoryRedis) BindingSessionIDWithUserID(ctx context.Context, userSession *model.UserSession) error {
	key := sr.generateKey(userSession.SessionID)
	userIDBytes, err := json.Marshal(userSession.UserID)
	if err != nil {
		return apperrors.SessionRepositoryRedisBindingSessionIDWithUserIDMarshalError.AppendMessage(err)
	}

	if err := sr.redis.RedisClient.Set(ctx, key, userIDBytes, userTtl).Err(); err != nil {
		return apperrors.SessionRepositoryRedisBindingSessionIDWithUserIDSetError.AppendMessage(err)
	}

	return nil
}

func (sr *SessionRepositoryRedis) CreateUserIDBySessionID(ctx context.Context, sessionID string) (*uuid.UUID, error) {
	key := sr.generateKey(sessionID)
	userID := uuid.New()
	if err := sr.redis.RedisClient.Set(ctx, key, userID, userTtl).Err(); err != nil {
		return nil, apperrors.SessionRepositoryRedisCreateUserIDBySessionSetError.AppendMessage(err)
	}

	return &userID, nil
}

func (sr *SessionRepositoryRedis) DeleteSessionBySessionID(ctx context.Context, sessionID string) error {
	key := sr.generateKey(sessionID)
	if err := sr.redis.RedisClient.Del(ctx, key).Err(); err != nil {
		return apperrors.SessionRepositoryRedisDeleteSessionBySessionIDError.AppendMessage(err)
	}

	return nil
}

func (sr *SessionRepositoryRedis) generateKey(sessionID string) string {
	return userPrefix + sessionID
}
