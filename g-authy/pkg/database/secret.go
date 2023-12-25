package database

import "github.com/google/uuid"

// Secret represents the secret information.
type Secret struct {
	DefaultModel
	VaultID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Vault   Vault     `gorm:"foreignkey:VaultID" json:"-"`
	Type    string    `gorm:"not null" json:"type"`
	Value   string    `gorm:"not null" json:"value"`
}
