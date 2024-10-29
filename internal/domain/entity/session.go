package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id           uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	UserId       uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	RefreshToken string    `gorm:"type:varchar(255);not null"`
	UserAgent    string    `gorm:"type:varchar(255);not null"`
	ClientIP     string    `gorm:"type:varchar(255);not null"`
	IsBlocked    bool      `gorm:"not null;default:false"`
	ExpiredAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
