package database

import (
	"github.com/google/uuid"
)

// Tenant represents the tenant information.
type Tenant struct {
	DefaultModel
	Region  string    `gorm:"not null"`
	Users   []User    `gorm:"foreignkey:TenantID"`
	Groups  []Group   `gorm:"foreignkey:GroupID"`
	OwnerID uuid.UUID `gorm:"foreignkey:UserID;not null;type:uuid"`
}
