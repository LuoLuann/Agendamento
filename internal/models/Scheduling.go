package models

import (
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
)

type SchedulingStatus string

const (
	SchedulingStatusPending   SchedulingStatus = "pending"
	SchedulingStatusCancelled SchedulingStatus = "cancelled"
	SchedulingStatusAccepted  SchedulingStatus = "accepted"
)

// Implemente Valuer/Scanner para o Status
func (s *SchedulingStatus) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*s = SchedulingStatus(v)
		return nil
	}
	return errors.New("invalid scheduling status")
}

func (s SchedulingStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type Scheduling struct {
	Base
	VacancyID uuid.UUID        `gorm:"type:uuid;not null"`
	Vacancy   Vacancy          `gorm:"foreignKey:VacancyID"`
	UserID    uuid.UUID        `gorm:"type:uuid;not null"`
	Status    SchedulingStatus `gorm:"type:varchar(20);not null;default:'pending'"`
}
