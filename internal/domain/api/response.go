package api

import (
	"time"

	"github.com/google/uuid"
)

type UserRespone struct {
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Tenant    string `json:"tenant"`
	Phone     string `json:"phone"`
	Note      string `json:"note"`
	CreatedBy string `json:"created_by"`
}

type LogInUserResponse struct {
	SessionID             uuid.UUID   `json:"session_id"`
	AccessToken           string      `json:"access_token"`
	AccessTokenExpiredAt  time.Time   `json:"access_token_expired_at"`
	RefreshToken          string      `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time   `json:"refresh_token_expired_at"`
	User                  UserRespone `json:"user"`
}
