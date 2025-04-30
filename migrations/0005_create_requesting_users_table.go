package migrations

import (
	"github.com/LuoLuann/Agendamento/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateRequestingUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250430_create_requesting_users_table",
		Migrate: func(d *gorm.DB) error {
			if err := d.AutoMigrate(&models.RequestingUser{}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(d *gorm.DB) error {
			if err := d.Migrator().DropTable("requesting_users"); err != nil {
				return err
			}
			return nil
		},
	}
}
