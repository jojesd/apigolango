package migrations

import (
	"github.com/hyperyuri/webap-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Fleets{})
}

func RunMigrations1(db *gorm.DB) {
	db.AutoMigrate(models.Fleets_Alert{})
}

func RunMigrations2(db *gorm.DB) {
	db.AutoMigrate(models.Vehicle{})
}

func RunMigrations3(db *gorm.DB) {
	db.AutoMigrate(models.Vehicle_Position{})
}
