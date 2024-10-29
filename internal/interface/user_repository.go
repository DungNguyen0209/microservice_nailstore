package interfaceObject

import (
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/entity"
)

// UserRepository is the interface defining methods for user data access
type UserRepository interface {
	CreateUser(user *entity.User) error
	FindUserById(id uuid.UUID) (*entity.User, error)
	FindUserByName(name string) (*entity.User, error)
}
