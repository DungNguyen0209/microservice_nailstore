package interfaceObject

import (
	"time"

	"github.com/minhdung/nailstore/internal/domain/entity"
)

// Maker is an interface for managing tokens
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *entity.PayLoad, error)

	//Verify check if token is valid or not
	VerifyToken(token string) (*entity.PayLoad, error)
}
