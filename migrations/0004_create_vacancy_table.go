package migrations

import (
	"github.com/LuoLuann/Agendamento/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateVacancyTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250430_create_vacancies_table",
		Migrate: func(d *gorm.DB) error {
			if err := d.AutoMigrate(&models.Vacancy{}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(d *gorm.DB) error {
			if err := d.Migrator().DropTable("vacancies"); err != nil {
				return err
			}
			return nil
		},
	}
}
