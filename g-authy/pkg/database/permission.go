package database

import "github.com/google/uuid"

// Permission represents the permission information.
type Permission struct {
	DefaultModel
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID" json:"-"`
	GroupID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Group      Group     `gorm:"foreignkey:GroupID" json:"-"`
	RoleID     int       `gorm:"type:int;autoIncrement;primaryKey"`
	Role       Role      `gorm:"foreignkey:RoleID" json:"role"`
	VaultID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Vault      Vault     `gorm:"foreignkey:VaultID" json:"-"`
	SecretID   uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Secret     Secret    `gorm:"foreignkey:SecretID" json:"-"`
	Permission string    `gorm:"not null" json:"permission"`
}
