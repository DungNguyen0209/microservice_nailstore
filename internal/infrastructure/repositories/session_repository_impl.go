package repositories

import (
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/entity"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

// CreateSession implements interfaceObject.SessionRepository.
func (s *SessionRepository) CreateSession(session *entity.Session) error {
	return s.db.Create(session).Error
}

// FindSessionByUserId implements interfaceObject.SessionRepository.
func (s *SessionRepository) FindSessionByUserId(id uuid.UUID) (*entity.Session, error) {
	var session entity.Session
	if err := s.db.First(&session, "UserId = ?", id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func NewSessionRepository(db *gorm.DB) interfaceObject.SessionRepository {
	return &SessionRepository{db}
}
