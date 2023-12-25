package database

import "github.com/google/uuid"

// Tenant represents the tenant information.
type Tenant struct {
	DefaultModel
	Region  string    `gorm:"not null" json:"region"`
	Users   []User    `gorm:"foreignkey:TenantID" json:"-"`
	Groups  []Group   `gorm:"foreignkey:GroupID" json:"-"`
	OwnerID uuid.UUID `gorm:"foreignkey:UserID;not null;type:uuid"`
}
