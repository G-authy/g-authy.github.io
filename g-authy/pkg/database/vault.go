package database

import "github.com/google/uuid"

// Vault represents the vault information.
type Vault struct {
	DefaultModel
	TenantID   uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Tenant     Tenant    `gorm:"foreignkey:TenantID" json:"-"`
	Secrets    []Secret  `gorm:"foreignkey:VaultID" json:"-"`
	VaultName  string    `gorm:"not null" json:"vault_name"`
	VaultOwner uuid.UUID `gorm:"foreignkey:UserID;not null;type:uuid"`
}
