package models

type RequestingUser struct {
	Base
	Name        string `gorm:"type:TEXT;not null"`
	Email       string `gorm:"type:TEXT;not null"`
	PhoneNumber string `gorm:"type:TEXT;not null"`
}
