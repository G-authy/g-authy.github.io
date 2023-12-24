package migrations

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func MigrateDefaultModel(db *gorm.DB) {
	db.AutoMigrate(&DefaultModel{})
}

// DefaultModel includes common fields for all models.
type DefaultModel struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tenant represents the tenant information.
type Tenant struct {
	DefaultModel
	Region string `gorm:"not null" json:"region"`
	Users  []User `gorm:"foreignkey:TenantID" json:"-"`
}

// User represents the user information.
type User struct {
	DefaultModel
	Username string    `gorm:"unique_index;not null" json:"username"`
	Email    string    `gorm:"unique_index;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"`
	TenantID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Tenant   Tenant    `gorm:"foreignkey:TenantID" json:"-"`
}
