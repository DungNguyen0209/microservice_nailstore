package interfaceObject

import (
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/entity"
)

type SessionRepository interface {
	CreateSession(session *entity.Session) error
	FindSessionByUserId(id uuid.UUID) (*entity.Session, error)
}
