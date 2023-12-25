package database

// Role represents the role information.
type Role struct {
	RoleID   int    `gorm:"type:int;autoIncrement;primaryKey"`
	RoleName string `gorm:"not null"`
}
