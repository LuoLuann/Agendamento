package migrations

import (
	"github.com/LuoLuann/Agendamento/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250430_create_users_table",
		Migrate: func(d *gorm.DB) error {
			if err := d.AutoMigrate(&models.User{}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(d *gorm.DB) error {
			if err := d.Migrator().DropTable("users"); err != nil {
				return err
			}
			return nil
		},
	}
}
