package migrations

import (
	"book-rest-api/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {

	db.AutoMigrate(models.Book{})
}