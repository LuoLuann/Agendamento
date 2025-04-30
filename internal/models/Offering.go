package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
)

type DayOfWeek int

const (
	Sunday    DayOfWeek = 0
	Monday    DayOfWeek = 1
	Tuesday   DayOfWeek = 2
	Wednesday DayOfWeek = 3
	Thursday  DayOfWeek = 4
	Friday    DayOfWeek = 5
	Saturday  DayOfWeek = 6
)

// Interface scanner para ler do bd
func (d *DayOfWeek) Scan(value interface{}) error {
	if v, ok := value.(int64); ok {
		*d = DayOfWeek(v)
		return nil
	}
	return errors.New("invalid value for DayOfWeek")
}

// Implementa a interface Valuer (para salvar no banco de dados)
func (d DayOfWeek) Value() (driver.Value, error) {
	return int(d), nil
}

// String() para exibição amigável
func (d DayOfWeek) String() string {
	names := [...]string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}
	if d < Sunday || d > Saturday {
		return "Desconhecido"
	}
	return names[d]
}

type Offering struct {
	Base
	ServiceID        uuid.UUID `gorm:"type:uuid;not null"`
	Title            string    `gorm:"type:TEXT;size:100;not null"`
	Description      string    `gorm:"type:TEXT;not null"`
	Enabled          bool      `gorm:"not null;default:true"`
	DayOfWeek        DayOfWeek `gorm:"type:INTEGER;not null"`
	StartAt          time.Time `gorm:"not null"`
	EndAt            time.Time `gorm:"not null"`
	RequiresApproval bool      `gorm:"not null;default:false"`
	NeedContact      bool      `gorm:"not null;default:false"`
	Vancancies       []Vacancy `gorm:"foreignKey:OfferingID;constraint:OnDelete:CASCADE"`
}
