package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		CreateUserTable(),
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("Migrations executed successfully!")
}
