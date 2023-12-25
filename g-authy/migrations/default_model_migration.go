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
	Region  string    `gorm:"not null" json:"region"`
	Users   []User    `gorm:"foreignkey:TenantID" json:"-"`
	Groups  []Group   `gorm:"foreignkey:GroupID" json:"-"`
	OwnerID uuid.UUID `gorm:"foreignkey:UserID;not null;type:uuid"`
}

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

// Vault represents the vault information.
type Vault struct {
	DefaultModel
	TenantID   uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Tenant     Tenant    `gorm:"foreignkey:TenantID" json:"-"`
	Secrets    []Secret  `gorm:"foreignkey:VaultID" json:"-"`
	VaultName  string    `gorm:"not null" json:"vault_name"`
	VaultOwner uuid.UUID `gorm:"foreignkey:UserID;not null;type:uuid"`
}

// Secret represents the secret information.
type Secret struct {
	DefaultModel
	VaultID     uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Vault       Vault     `gorm:"foreignkey:VaultID" json:"-"`
	Type        string    `gorm:"not null" json:"type"`
	Value       string    `gorm:"not null" json:"value"`
	SecretName  string    `gorm:"not null" json:"secret_name"`
	SecretOwner uuid.UUID `gorm:"foreignkey:UserID;not null;type:uuid"`
}

// Group represents the Groups table
type Group struct {
	DefaultModel
}

// Resource represents the Resources table
type Resource struct {
	ResourceID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	ResourceName       string
	ResourceCreationID uuid.UUID `gorm:"type:uuid"`
	ModificationDate   time.Time
}

// Role represents the Roles table
type Role struct {
	DefaultModel
}

// Permission represents the Permissions table
type Permission struct {
	DefaultModel
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID" json:"-"`
	GroupID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Group      Group     `gorm:"foreignkey:GroupID" json:"-"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Role       Role      `gorm:"foreignkey:RoleID" json:"-"`
	VaultID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Vault      Vault     `gorm:"foreignkey:VaultID" json:"-"`
	Permission string    `gorm:"not null" json:"permission"`
}
