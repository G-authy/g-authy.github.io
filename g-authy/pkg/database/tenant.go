package database

// Tenant represents the tenant information.
type Tenant struct {
	DefaultModel
	Region string `gorm:"not null"`
	Users  []User `gorm:"foreignkey:TenantID"`
}
