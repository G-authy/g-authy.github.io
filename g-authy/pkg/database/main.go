package database

import (
	"github.com/jinzhu/gorm"
)

// AutoMigrate will auto migrate the database tables.
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Tenant{})
	db.AutoMigrate(&User{})
}
