package database

import "github.com/google/uuid"

// User represents the user information.
type User struct {
	DefaultModel
	Username string    `gorm:"unique_index;not null"`
	Email    string    `gorm:"unique_index;not null"`
	Password string    `gorm:"not null"`
	TenantID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Tenant   Tenant    `gorm:"foreignkey:TenantID" json:"-"`
}
