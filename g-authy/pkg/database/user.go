package database

import "github.com/google/uuid"

// User represents the user information.
type User struct {
	DefaultModel
	Username       string    `gorm:"unique_index;not null" json:"username"`
	Email          string    `gorm:"unique_index;not null" json:"email"`
	Password       string    `gorm:"not null" json:"-"`
	TenantID       uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Tenant         Tenant    `gorm:"foreignkey:TenantID" json:"-"`
	Groups         []Group   `gorm:"foreignkey:GroupID;not null" json:"-"`
	IsActive       bool
	ActivationLink string
	IsSSOActive    bool
}
