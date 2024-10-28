package usecase

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/entity"
	et "github.com/minhdung/nailstore/internal/domain/entity"
	"github.com/minhdung/nailstore/internal/domain/request"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
)

type userUsecaseImpl struct {
	UserRepository interfaceObject.UserRepository
}

// NewUserUsecaseImpl creates a new instance of UserUsecase with the provided repository
func NewUserUsecaseImpl(userRepo interfaceObject.UserRepository) interfaceObject.UserUsecase {
	return &userUsecaseImpl{
		UserRepository: userRepo,
	}
}

// CreateUser creates a new user
func (u *userUsecaseImpl) CreateUser(user request.UserRequest) error {
	arg := et.User{
		Id:          uuid.New(),
		Username:    user.Username,
		Password:    user.Password,
		UpdatedBy:   user.UpdatedBy,
		CreatedTime: time.Now().UTC(),
		Note:        user.Note,
		Email:       user.Email,
		Tenant:      user.Tenant,
	}
	err := u.UserRepository.CreateUser(&arg)
	if err != nil {
		log.Printf("Error creating user: %v", err)
	}
	return err
}

// FindByID finds a user by ID
func (u *userUsecaseImpl) FindUserById(id uuid.UUID) (*entity.User, error) {
	user, err := u.UserRepository.FindUserById(id)
	if err != nil {
		log.Printf("Error finding user by ID: %v", err)
		return nil, err
	}
	log.Printf("Found user: %+v", user)
	return user, nil
}
