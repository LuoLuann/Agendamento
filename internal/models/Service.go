package models

import (
	"github.com/google/uuid"
)

type Service struct {
	Base
	UserID      uuid.UUID  `gorm:"type:uuid;not null"`
	Title       string     `gorm:"type:TEXT;not null"`
	Description string     `gorm:"type:TEXT;not null"`
	Price       int        `gorm:"type:INTEGER"`
	Slug        string     `gorm:"type:TEXT;uniqueIndex"`
	Enabled     bool       `gorm:"not null;default:true"`
	Offerings   []Offering `gorm:"foreignKey:ServiceID;constraint:OnDelete:CASCADE"`
}
