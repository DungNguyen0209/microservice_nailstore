package usecases

import (
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/entity"
	"github.com/minhdung/nailstore/internal/domain/request"
)

type UserUsecase interface {
	CreateUser(user request.UserRequest) error
	FindUserById(id uuid.UUID) (*entity.User, error)
}
