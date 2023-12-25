package database

// Group represents the group information.
type Group struct {
	DefaultModel
	Users []User `gorm:"foreignkey:TenantID"`
}
