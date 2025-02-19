package model

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

func (j *JwtCustomClaims) Valid() error {
	if j.ExpiresAt.Unix() < time.Now().Unix() {
		return fmt.Errorf("%s", jwt.ErrTokenExpired)
	}
	return nil
}
