package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleProvider UserRole = "provider"
	RoleCustomer UserRole = "customer"
)

type User struct {
	gorm.Model
	Base
	Name           string    `gorm:"size:100;not null"`
	Email          string    `gorm:"uniqueIndex;not null"`
	Password       string    `gorm:"not null"`
	Phone          string    `gorm:"type:TEXT;not null"`
	ProfilePicture string    `gorm:"type:TEXT;not null"`
	Enabled        bool      `gorm:"not null;default:true"`
	AccessLevel    string    `gorm:"type:TEXT;not null"`
	Role           string    `gorm:"type:varchar(20);not null;default:'provider'"`
	Services       []Service `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// ProviderDetails *ProviderDetails `gorm:"embedded"`
	// CustomerDetails *CustomerDetails `gorm:"embedded"`
}

// type CustomerDetails struct {
// 	PreferencesJSON string `gorm:"type:TEXT"` // Ou use JSONB no PostgreSQL
// }
