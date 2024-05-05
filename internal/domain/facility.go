package domain

import (
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	ID          int64 `gorm:"primaryKey"`
	CampsiteID  int64
	Name        string
	Description string         `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3)"`
	DeletedAt   gorm.DeletedAt `gorm:"type:datetime(3)"`
}
