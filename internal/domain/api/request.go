package api

import (
	"time"

	"github.com/google/uuid"
)

type LogInUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRequest struct {
	Username    string    `json:"username" binding:"required,alphanum"`
	Password    string    `json:"password" binding:"required,min=6"`
	FullName    string    `json:"full_name" binding:"required"`
	Tenant      string    `json:"tenant" binding:"required"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Note        string    `json:"note" binding:"required"`
	CreatedTime time.Time `json:"createdTime"`
	CreatedBy   uuid.UUID `json:"createdBy" binding:"required"`
	UpdatedBy   uuid.UUID `json:"updatedBy"`
}

type GetUserRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}
