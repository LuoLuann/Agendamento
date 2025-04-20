package migrations

import (
	"github.com/LuoLuann/Agendamento/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240601000000",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
