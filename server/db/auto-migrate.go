package db

import (
	"github.com/maggickkw/ticket-master/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
