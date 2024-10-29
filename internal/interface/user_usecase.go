package interfaceObject

import (
	"github.com/google/uuid"
	request "github.com/minhdung/nailstore/internal/domain/api"
	"github.com/minhdung/nailstore/internal/domain/entity"
)

type UserUsecase interface {
	CreateUser(user request.UserRequest) error
	FindUserById(id uuid.UUID) (*entity.User, error)
	GetUserByName(name string) (*entity.User, error)
	CreateSession(entity.Session) (*entity.Session, error)
}
