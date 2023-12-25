package database

import "github.com/google/uuid"

// Permission represents the permission information.
type PermissionSecret struct {
	DefaultModel
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID" json:"-"`
	GroupID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Group      Group     `gorm:"foreignkey:GroupID" json:"-"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Role       Role      `gorm:"foreignkey:RoleID" json:"role"`
	SecretID   uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Secret     Secret    `gorm:"foreignkey:SecretID" json:"-"`
	Permission string    `gorm:"not null" json:"permission"`
}
