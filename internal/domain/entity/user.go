package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID `gorm:"type:char(36);primary_key"`
	Username       string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	HashedPassword string    `gorm:"type:varchar(255);not null"`
	FullName       string    `gorm:"type:varchar(100);not null"`
	Tenant         string    `gorm:"type:varchar(50)"`
	Email          string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone          string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Note           string    `gorm:"type:text"`
	CreatedTime    time.Time `gorm:"autoCreateTime"`
	UpdatedTime    time.Time `gorm:"autoCreateTime"`
	UpdatedBy      uuid.UUID `gorm:"type:uuid"`
	CreatedBy      uuid.UUID `gorm:"type:uuid"`
}
