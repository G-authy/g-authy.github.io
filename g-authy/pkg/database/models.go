package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// DefaultModel includes common fields for all models.
type DefaultModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tenant represents the tenant information.
type Tenant struct {
	DefaultModel
	Region string `gorm:"not null"`
	Users  []User `gorm:"foreignkey:TenantID"`
}

// User represents the user information.
type User struct {
	DefaultModel
	Username string    `gorm:"unique_index;not null"`
	Email    string    `gorm:"unique_index;not null"`
	Password string    `gorm:"not null"`
	TenantID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Tenant   Tenant    `gorm:"foreignkey:TenantID" json:"-"`
}

// AutoMigrate will auto migrate the database tables.
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Tenant{})
	db.AutoMigrate(&User{})
}
