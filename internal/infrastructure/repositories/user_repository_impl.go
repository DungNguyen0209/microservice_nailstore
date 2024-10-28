// infrastructure/repositories/user_repo.go
package repositories

import (
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/entity"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) interfaceObject.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindUserById(id uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
