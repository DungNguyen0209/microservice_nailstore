package request

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	FullName    string `json:"full_name" binding:"required"`
	Tenant      string `json:"tenant" binding:"required"`
	Email       string `json:"email"`
	Note        string `json:"note" binding:"required"`
	CreatedTime time.Time
	UpdatedBy   uuid.UUID `json:"updatedBy"`
}

type GetUserRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}
