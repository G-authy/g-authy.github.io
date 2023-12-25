package database

import (
	"time"

	"github.com/google/uuid"
)

// Resource represents the Resources table
type Resource struct {
	ResourceID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	ResourceName       string
	ResourceCreationID uuid.UUID `gorm:"type:uuid"`
	ModificationDate   time.Time
}
