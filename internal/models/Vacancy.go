package models

import (
	"time"

	"github.com/google/uuid"
)

type Vacancy struct {
	Base
	OfferingID uuid.UUID `gorm:"type:uuid;not null"`
	Offering   Offering  `gorm:"foreignKey:OfferingID;references:ID"`
	StartAt    time.Time `gorm:"type:timestamp;not null"`
	EndAt      time.Time `gorm:"type:timestamp;not null"`
	Vacancies  int       `gorm:"not null"`
}
