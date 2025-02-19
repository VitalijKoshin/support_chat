package model

import "github.com/google/uuid"

type UserSession struct {
	UserSessionID uuid.UUID `json:"user_session_id" db:"user_session_id" validate:"omitempty"`
	UserID        uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	SessionID     string    `json:"session_id" db:"session_id" validate:"omitempty"`
}

type UserSessionDB struct {
	UserSessionID uuid.UUID `db:"user_session_id"`
	UserID        uuid.UUID `db:"user_id"`
	SessionID     string    `db:"session_id"`
	CreatedAt     string    `db:"created_at"`
	UpdatedAt     string    `db:"updated_at"`
	DeletedAt     string    `db:"deleted_at"`
}

func NewUserSession(userID uuid.UUID, sessionID string) *UserSession {
	return &UserSession{
		UserSessionID: uuid.New(),
		UserID:        userID,
		SessionID:     sessionID,
	}
}

func CreateUserSessionModelFromDB(userSessionDB *UserSessionDB) *UserSession {
	return &UserSession{
		UserSessionID: userSessionDB.UserSessionID,
		UserID:        userSessionDB.UserID,
		SessionID:     userSessionDB.SessionID,
	}
}
